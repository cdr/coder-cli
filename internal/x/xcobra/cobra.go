// Package xcobra wraps the cobra package to provide richer functionality.
package xcobra

import (
	"fmt"

	"github.com/spf13/cobra"

	"cdr.dev/coder-cli/pkg/clog"
)

// ExactArgs returns an error if there are not exactly n args.
func ExactArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != n {
			return clog.Error(
				fmt.Sprintf("accepts %d arg(s), received %d", n, len(args)),
				clog.Bold("usage: ")+cmd.UseLine(),
				clog.BlankLine,
				clog.Tipf("use \"--help\" for more info"),
			)
		}
		return nil
	}
}
