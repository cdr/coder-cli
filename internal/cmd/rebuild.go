package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	"cdr.dev/coder-cli/coder-sdk"
	"cdr.dev/coder-cli/internal/x/xcobra"
	"cdr.dev/coder-cli/pkg/clog"
)

func rebuildWorkspaceCommand() *cobra.Command {
	var follow bool
	var force bool
	var user string
	cmd := &cobra.Command{
		Use:   "rebuild [workspace_name]",
		Short: "rebuild a Coder workspace",
		Args:  xcobra.ExactArgs(1),
		Example: `coder workspaces rebuild front-end-workspace --follow
coder workspaces rebuild backend-workspace --force`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client, err := newClient(ctx, true)
			if err != nil {
				return err
			}
			workspace, err := findWorkspace(ctx, client, args[0], user)
			if err != nil {
				return err
			}

			if !force && workspace.LatestStat.ContainerStatus == coder.WorkspaceOn {
				_, err = (&promptui.Prompt{
					Label:     fmt.Sprintf("Rebuild workspace %q? (will destroy any work outside of your home directory)", workspace.Name),
					IsConfirm: true,
				}).Run()
				if err != nil {
					return clog.Fatal(
						"failed to confirm prompt", clog.BlankLine,
						clog.Tipf(`use "--force" to rebuild without a confirmation prompt`),
					)
				}
			}

			if err = client.RebuildWorkspace(ctx, workspace.ID); err != nil {
				return err
			}
			if follow {
				if err = trailBuildLogs(ctx, client, workspace.ID); err != nil {
					return err
				}
			} else {
				clog.LogSuccess(
					"successfully started rebuild",
					clog.Tipf("run \"coder workspaces watch-build %s\" to follow the build logs", workspace.Name),
				)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&user, "user", coder.Me, "Specify the user whose resources to target")
	cmd.Flags().BoolVar(&follow, "follow", false, "follow build log after initiating rebuild")
	cmd.Flags().BoolVar(&force, "force", false, "force rebuild without showing a confirmation prompt")
	return cmd
}

// trailBuildLogs follows the build log for a given workspace and prints the staged
// output with loaders and success/failure indicators for each stage.
func trailBuildLogs(ctx context.Context, client coder.Client, workspaceID string) error {
	const check = "✅"
	const failure = "❌"

	newSpinner := func() *spinner.Spinner { return spinner.New(spinner.CharSets[11], 100*time.Millisecond) }

	// this tells us whether to show dynamic loaders when printing output
	isTerminal := showInteractiveOutput

	logs, err := client.FollowWorkspaceBuildLog(ctx, workspaceID)
	if err != nil {
		return err
	}

	var s *spinner.Spinner
	for l := range logs {
		if l.Err != nil {
			return l.Err
		}

		logTime := l.BuildLog.Time.Local()
		msg := fmt.Sprintf("%s %s", logTime.Format(time.RFC3339), l.BuildLog.Msg)

		switch l.BuildLog.Type {
		case coder.BuildLogTypeStart:
			// the FE uses this to reset the UI
			// the CLI doesn't need to do anything here given that we only append to the trail

		case coder.BuildLogTypeStage:
			if !isTerminal {
				fmt.Println(msg)
				continue
			}

			if s != nil {
				s.Stop()
				fmt.Print("\n")
			}

			s = newSpinner()
			s.Suffix = fmt.Sprintf("  -- %s", msg)
			s.FinalMSG = fmt.Sprintf("%s -- %s", check, msg)
			s.Start()

		case coder.BuildLogTypeSubstage:
			// TODO(@f0ssel) add verbose substage printing
			if !verbose {
				continue
			}

		case coder.BuildLogTypeError:
			if !isTerminal {
				fmt.Println(msg)
				continue
			}

			if s != nil {
				s.FinalMSG = fmt.Sprintf("%s %s", failure, strings.TrimPrefix(s.Suffix, "  "))
				s.Stop()
				fmt.Print("\n")
			}

			s = newSpinner()
			s.Suffix = color.RedString("  -- %s", msg)
			s.FinalMSG = color.RedString("%s -- %s", failure, msg)
			s.Start()

		case coder.BuildLogTypeDone:
			if s != nil {
				s.Stop()
				fmt.Print("\n")
			}

			return nil
		default:
			return xerrors.Errorf("unknown buildlog type: %s", l.BuildLog.Type)
		}
	}
	return nil
}

func watchBuildLogCommand() *cobra.Command {
	var user string
	cmd := &cobra.Command{
		Use:     "watch-build [workspace_name]",
		Example: "coder workspaces watch-build front-end-workspace",
		Short:   "trail the build log of a Coder workspace",
		Args:    xcobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client, err := newClient(ctx, true)
			if err != nil {
				return err
			}
			workspace, err := findWorkspace(ctx, client, args[0], user)
			if err != nil {
				return err
			}

			if err = trailBuildLogs(ctx, client, workspace.ID); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&user, "user", coder.Me, "Specify the user whose resources to target")
	return cmd
}
