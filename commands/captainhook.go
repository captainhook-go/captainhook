package commands

import (
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
	quietFlag   bool
	verboseFlag bool
	debugFlag   bool
	colorFlag   bool
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
	rootCmd.PersistentFlags().BoolVarP(&quietFlag, "quiet", "q", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&colorFlag, "no-color", "", false, "disable colored output")

	hookCommand := setupHookCommand()
	hookCommand.AddCommand(SetupHookCommitMsgCommand())
	hookCommand.AddCommand(SetupHookPrepareCommitMsgCommand())
	hookCommand.AddCommand(SetupHookPreCommitCommand())
	hookCommand.AddCommand(SetupHookPrePushCommand())
	hookCommand.AddCommand(SetupHookPostCommitCommand())
	hookCommand.AddCommand(SetupHookPostRewriteCommand())
	hookCommand.AddCommand(SetupHookPostCheckoutCommand())
	hookCommand.AddCommand(SetupHookPostMergeCommand())

	rootCmd.AddCommand(setupInstallCommand())
	rootCmd.AddCommand(hookCommand)
}
