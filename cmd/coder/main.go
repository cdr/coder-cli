package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"

	"cdr.dev/coder-cli/internal/cmd"
	"cdr.dev/coder-cli/internal/version"
	"cdr.dev/coder-cli/internal/x/xterminal"
	"cdr.dev/coder-cli/pkg/clog"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// If requested, spin up the pprof webserver.
	if os.Getenv("PPROF") != "" {
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}

	restoreTerminal := func() {}

	// Janky, but SSH on windows sets the output to raw.
	// If we set it ourselves, SSH fails because the FD isn't found.
	if len(os.Args) >= 2 && os.Args[1] != "tunnel" {
		state, err := xterminal.MakeOutputRaw(os.Stdout.Fd())
		if err != nil {
			clog.Log(clog.Fatal(fmt.Sprintf("set output to raw: %s", err)))
			cancel()
			os.Exit(1)
		}
		restoreTerminal = func() {
			// Best effort. Would result in broken terminal on window but nothing we can do about it.
			_ = xterminal.Restore(os.Stdout.Fd(), state)
		}
	}

	app := cmd.Make()
	app.Version = fmt.Sprintf("%s %s %s/%s", version.Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)

	if err := app.ExecuteContext(ctx); err != nil {
		clog.Log(err)
		cancel()
		restoreTerminal()
		os.Exit(1)
	}
	cancel()
	restoreTerminal()
}
