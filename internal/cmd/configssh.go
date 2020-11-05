package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"cdr.dev/coder-cli/coder-sdk"
	"cdr.dev/coder-cli/internal/config"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

func configSSHCmd() *cobra.Command {
	var (
		configpath string
		remove     = false
	)

	cmd := &cobra.Command{
		Use:   "config-ssh",
		Short: "Configure SSH to access Coder environments",
		Long:  "Inject the proper OpenSSH configuration into your local SSH config file.",
		RunE:  configSSH(&configpath, &remove),
	}
	cmd.Flags().StringVar(&configpath, "filepath", filepath.Join("~", ".ssh", "config"), "overide the default path of your ssh config file")
	cmd.Flags().BoolVar(&remove, "remove", false, "remove the auto-generated Coder Enterprise ssh config")

	return cmd
}

func configSSH(configpath *string, remove *bool) func(cmd *cobra.Command, _ []string) error {
	const startToken = "# ------------START-CODER-ENTERPRISE-----------"
	startMessage := `# The following has been auto-generated by "coder config-ssh"
# to make accessing your Coder Enterprise environments easier.
#
# To remove this blob, run:
#
#    coder config-ssh --remove
#
# You should not hand-edit this section, unless you are deleting it.`
	const endToken = "# ------------END-CODER-ENTERPRISE------------"

	return func(cmd *cobra.Command, _ []string) error {
		ctx := cmd.Context()
		usr, err := user.Current()
		if err != nil {
			return xerrors.Errorf("get user home directory: %w", err)
		}

		privateKeyFilepath := filepath.Join(usr.HomeDir, ".ssh", "coder_enterprise")

		if strings.HasPrefix(*configpath, "~") {
			*configpath = strings.Replace(*configpath, "~", usr.HomeDir, 1)
		}

		currentConfig, err := readStr(*configpath)
		if os.IsNotExist(err) {
			// SSH configs are not always already there.
			currentConfig = ""
		} else if err != nil {
			return xerrors.Errorf("read ssh config file %q: %w", *configpath, err)
		}

		startIndex := strings.Index(currentConfig, startToken)
		endIndex := strings.Index(currentConfig, endToken)

		if *remove {
			if startIndex == -1 || endIndex == -1 {
				return xerrors.Errorf("the Coder Enterprise ssh configuration section could not be safely deleted or does not exist")
			}
			currentConfig = currentConfig[:startIndex-1] + currentConfig[endIndex+len(endToken)+1:]

			err = writeStr(*configpath, currentConfig)
			if err != nil {
				return xerrors.Errorf("write to ssh config file %q: %s", *configpath, err)
			}

			return nil
		}

		client, err := newClient(ctx)
		if err != nil {
			return err
		}

		if !isSSHAvailable(ctx) {
			return xerrors.New("SSH is disabled or not available for your Coder Enterprise deployment.")
		}

		user, err := client.Me(ctx)
		if err != nil {
			return xerrors.Errorf("fetch username: %w", err)
		}

		envs, err := getEnvs(ctx, client, coder.Me)
		if err != nil {
			return err
		}
		if len(envs) < 1 {
			return xerrors.New("no environments found")
		}
		newConfig, err := makeNewConfigs(user.Username, envs, startToken, startMessage, endToken, privateKeyFilepath)
		if err != nil {
			return xerrors.Errorf("make new ssh configurations: %w", err)
		}

		// if we find the old config, remove those chars from the string
		if startIndex != -1 && endIndex != -1 {
			currentConfig = currentConfig[:startIndex-1] + currentConfig[endIndex+len(endToken)+1:]
		}

		err = os.MkdirAll(filepath.Dir(*configpath), os.ModePerm)
		if err != nil {
			return xerrors.Errorf("make configuration directory: %w", err)
		}
		err = writeStr(*configpath, currentConfig+newConfig)
		if err != nil {
			return xerrors.Errorf("write new configurations to ssh config file %q: %w", *configpath, err)
		}
		err = writeSSHKey(ctx, client, privateKeyFilepath)
		if err != nil {
			if !xerrors.Is(err, os.ErrPermission) {
				return xerrors.Errorf("write ssh key: %w", err)
			}
			fmt.Printf("Your private ssh key already exists at \"%s\"\nYou may need to remove the existing private key file and re-run this command\n\n", privateKeyFilepath)
		} else {
			fmt.Printf("Your private ssh key was written to \"%s\"\n", privateKeyFilepath)
		}

		fmt.Printf("An auto-generated ssh config was written to \"%s\"\n", *configpath)
		fmt.Println("You should now be able to ssh into your environment")
		fmt.Printf("For example, try running\n\n\t$ ssh coder.%s\n\n", envs[0].Name)
		return nil
	}
}

func writeSSHKey(ctx context.Context, client *coder.Client, privateKeyPath string) error {
	key, err := client.SSHKey(ctx)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(privateKeyPath, []byte(key.PrivateKey), 0400)
}

func makeNewConfigs(userName string, envs []coder.Environment, startToken, startMsg, endToken, privateKeyFilepath string) (string, error) {
	hostname, err := configuredHostname()
	if err != nil {
		return "", err
	}

	newConfig := fmt.Sprintf("\n%s\n%s\n\n", startToken, startMsg)
	for _, env := range envs {
		newConfig += makeSSHConfig(hostname, userName, env.Name, privateKeyFilepath)
	}
	newConfig += fmt.Sprintf("\n%s\n", endToken)

	return newConfig, nil
}

func makeSSHConfig(host, userName, envName, privateKeyFilepath string) string {
	return fmt.Sprintf(
		`Host coder.%s
   HostName %s
   User %s-%s
   StrictHostKeyChecking no
   ConnectTimeout=0
   IdentityFile="%s"
   ServerAliveInterval 60
   ServerAliveCountMax 3
`, envName, host, userName, envName, privateKeyFilepath)
}

func isSSHAvailable(ctx context.Context) bool {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	host, err := configuredHostname()
	if err != nil {
		return false
	}

	var dialer net.Dialer
	_, err = dialer.DialContext(ctx, "tcp", net.JoinHostPort(host, "22"))
	return err == nil
}

func configuredHostname() (string, error) {
	u, err := config.URL.Read()
	if err != nil {
		return "", err
	}
	url, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	return url.Hostname(), nil
}

func writeStr(filename, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0777)
}

func readStr(filename string) (string, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
