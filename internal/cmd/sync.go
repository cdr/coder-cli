package cmd

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	"cdr.dev/coder-cli/coder-sdk"
	"cdr.dev/coder-cli/internal/sync"
	"cdr.dev/coder-cli/internal/x/xcobra"
	"cdr.dev/coder-cli/pkg/clog"
)

func syncCmd() *cobra.Command {
	var init bool
	cmd := &cobra.Command{
		Use:   "sync [local directory] [<workspace name>:<remote directory>]",
		Short: "Establish a one way directory sync to a Coder workspace",
		Args:  xcobra.ExactArgs(2),
		RunE:  makeRunSync(&init),
	}
	cmd.Flags().BoolVar(&init, "init", false, "do initial transfer and exit")
	return cmd
}

// rsyncVersion returns local rsync protocol version as a string.
func rsyncVersion() string {
	cmd := exec.Command("rsync", "--version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	firstLine, err := bytes.NewBuffer(out).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	versionString := strings.Split(firstLine, "protocol version ")

	return versionString[1]
}

func makeRunSync(init *bool) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var (
			ctx    = cmd.Context()
			local  = args[0]
			remote = args[1]
		)

		client, err := newClient(ctx, true)
		if err != nil {
			return err
		}

		remoteTokens := strings.SplitN(remote, ":", 2)
		if len(remoteTokens) != 2 {
			return xerrors.New("remote malformatted")
		}
		var (
			workspaceName = remoteTokens[0]
			remoteDir     = remoteTokens[1]
		)

		workspace, err := findWorkspace(ctx, client, workspaceName, coder.Me)
		if err != nil {
			return err
		}

		info, err := os.Stat(local)
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() {
			return sync.SingleFile(ctx, local, remoteDir, workspace, client)
		}
		if !info.IsDir() {
			return xerrors.Errorf("local path must lead to a regular file or directory: %w", err)
		}

		absLocal, err := filepath.Abs(local)
		if err != nil {
			return xerrors.Errorf("make abs path out of %s, %s: %w", local, absLocal, err)
		}

		s := sync.Sync{
			Init:                *init,
			Workspace:           *workspace,
			RemoteDir:           remoteDir,
			LocalDir:            absLocal,
			Client:              client,
			OutW:                cmd.OutOrStdout(),
			ErrW:                cmd.ErrOrStderr(),
			InputReader:         cmd.InOrStdin(),
			IsInteractiveOutput: showInteractiveOutput,
		}

		localVersion := rsyncVersion()
		remoteVersion, rsyncErr := s.Version()

		if rsyncErr != nil {
			clog.LogInfo("unable to determine remote rsync version: proceeding cautiously")
		} else if localVersion != remoteVersion {
			return xerrors.Errorf("rsync protocol mismatch: local = %s, remote = %s", localVersion, remoteVersion)
		}

		for err == nil || err == sync.ErrRestartSync {
			err = s.Run()
		}
		if err != nil {
			return err
		}
		return nil
	}
}
