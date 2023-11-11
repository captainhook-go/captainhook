package commands

import (
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
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
	quietFlag       bool
	verboseFlag     bool
	debugFlag       bool
	colorFlag       bool
	interactionFlag bool
	rootCmd         = &cobra.Command{
		Use:   "captainhook",
		Short: description(),
		Long:  description(),
	}
)

func description() string {
	return io.Colorize(
		"<ok>CaptainHook</ok> " +
			"version <comment>" + info.VERSION + "</comment> " +
			info.RELEASE_DATE +
			" <strong>#StandWith</strong><comment>Ukraine</comment>")
}

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
	rootCmd.PersistentFlags().BoolVarP(&quietFlag, "quiet", "q", false, "Do not output any message")
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "Increase the verbosity of messages")
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "Increase the verbosity even more")
	rootCmd.PersistentFlags().BoolVarP(&colorFlag, "no-color", "", false, "Disable colored output")
	rootCmd.PersistentFlags().BoolVarP(&interactionFlag, "no-interaction", "n", false, "Do not ask interactive questions")

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
