package commands

import (
	"github.com/captainhook-go/captainhook/commands/hooks"
	"github.com/spf13/cobra"
	"os"
)

type Response struct {
	// Err is set when the command failed to execute.
	Err error
	// The command that was executed.
	Cmd *cobra.Command
}

func (r Response) IsUserError() bool {
	return r.Err != nil && isUserError(r.Err)
}

var (
	verboseFlag bool
	rootCmd     = &cobra.Command{
		Use:   "captainhook",
		Short: "Git hook manager",
		Long:  "CaptainHook is a git hook manager",
	}
)

func Execute([]string) Response {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	var resp Response
	resp.Cmd = rootCmd

	return resp
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&verboseFlag,
		"verbose",
		"v", false,
		"verbose output",
	)

	hookCommand := setupHookCommand()
	hookCommand.AddCommand(hooks.SetupHookCommitMsgCommand())
	hookCommand.AddCommand(hooks.SetupHookPrepareCommitMsgCommand())
	hookCommand.AddCommand(hooks.SetupHookPreCommitCommand())
	hookCommand.AddCommand(hooks.SetupHookPrePushCommand())
	hookCommand.AddCommand(hooks.SetupHookPostCommitCommand())
	hookCommand.AddCommand(hooks.SetupHookPostRewriteCommand())
	hookCommand.AddCommand(hooks.SetupHookPostCheckoutCommand())
	hookCommand.AddCommand(hooks.SetupHookPostMergeCommand())

	rootCmd.AddCommand(setupInstallCommand())
	rootCmd.AddCommand(hookCommand)
}
