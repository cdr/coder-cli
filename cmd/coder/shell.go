package main

import (
	"context"
	"io"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/pflag"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
	"golang.org/x/time/rate"
	"golang.org/x/xerrors"

	"go.coder.com/cli"
	"go.coder.com/flog"

	"cdr.dev/wsep"
)

type shellCmd struct {
	container string
}

func (cmd *shellCmd) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:  "sh",
		Usage: "<env name> [flags] -- [<command [args...]>]",
		Desc:  "executes a remote command on the environment\nIf no command is specified, the default shell is opened.",
	}
}

func (cmd *shellCmd) RegisterFlags(fl *pflag.FlagSet) {
	fl.StringVar(&cmd.container, "container", "", "The container to execute against. Defaults to main container. Useful for execing into services.")
}

func enableTerminal(fd int) (restore func(), err error) {
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		return restore, xerrors.Errorf("make raw term: %w", err)
	}
	return func() {
		err := terminal.Restore(fd, state)
		if err != nil {
			flog.Error("restore term state: %v", err)
		}
	}, nil
}

func sendResizeEvents(ctx context.Context, termfd int, process wsep.Process) {
	sigs := make(chan os.Signal, 16)
	signal.Notify(sigs, unix.SIGWINCH)

	// Limit the frequency of resizes to prevent a stuttering effect.
	resizeLimiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 1)

	for ctx.Err() == nil {
		if ctx.Err() != nil {
			return
		}
		width, height, err := terminal.GetSize(termfd)
		if err != nil {
			flog.Error("get term size: %v", err)
			return
		}

		err = process.Resize(ctx, uint16(height), uint16(width))
		if err != nil {
			flog.Error("set term size: %v", err)
			return
		}

		// Do this last so the first resize is sent.
		<-sigs
		resizeLimiter.Wait(ctx)
	}
}

func (cmd *shellCmd) Run(fl *pflag.FlagSet) {
	if len(fl.Args()) < 1 {
		exitUsage(fl)
	}
	var (
		envName = fl.Arg(0)
		command = fl.Arg(1)
		ctx     = context.Background()
	)

	var args []string
	if command != "" {
		args = fl.Args()[2:]
	}

	// Bring user into shell if no command is specified.
	if command == "" {
		command = "sh"
		args = []string{"-c", "exec $(getent passwd $(whoami) | awk -F: '{ print $7 }')"}
	}

	err := runCommand(ctx, envName, cmd.container, command, args)
	if exitErr, ok := err.(wsep.ExitError); ok {
		os.Exit(exitErr.Code)
	}
	if err != nil {
		flog.Fatal("run command: %v. Is %q online?", err, envName)
	}
}

func runCommand(ctx context.Context, envName, container, command string, args []string) error {
	var (
		entClient = requireAuth()
		env       = findEnv(entClient, envName)
	)

	termfd := int(os.Stdin.Fd())

	tty := terminal.IsTerminal(termfd)
	if tty {
		restore, err := enableTerminal(termfd)
		if err != nil {
			return err
		}
		defer restore()
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn, err := entClient.DialWsep(ctx, env, container)
	if err != nil {
		return err
	}

	execer := wsep.RemoteExecer(conn)
	process, err := execer.Start(ctx, wsep.Command{
		Command: command,
		Args:    args,
		TTY:     tty,
	})
	if err != nil {
		return err
	}

	if tty {
		go sendResizeEvents(ctx, termfd, process)
	}

	go func() {
		stdin := process.Stdin()
		defer stdin.Close()
		_, err := io.Copy(stdin, os.Stdin)
		if err != nil {
			cancel()
		}
	}()
	go func() {
		_, err := io.Copy(os.Stdout, process.Stdout())
		if err != nil {
			cancel()
		}
	}()
	go func() {
		_, err := io.Copy(os.Stderr, process.Stderr())
		if err != nil {
			cancel()
		}
	}()
	err = process.Wait()
	if xerrors.Is(err, ctx.Err()) {
		return xerrors.Errorf("network error")
	}
	return err
}
