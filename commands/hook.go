package commands

import (
	"fmt"
	"github.com/captainhook-go/captainhook/exec"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"github.com/spf13/cobra"
	"strings"
)

func setupHookCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hook",
		Short: "Execute all actions registered for a git hook",
		Long:  "Execute all actions registered for a git hook",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Usage:")
			fmt.Println("captainhook hook [command] [options]\n")
			fmt.Println("Available Commands:")
			for _, hookName := range info.GetNativeHooks() {
				spaces := strings.Repeat(" ", 19-len(hookName)+2)
				fmt.Printf("  %s %sExecute %s actions\n", hookName, spaces, hookName)
			}
			fmt.Println("\nrun captainhook hook [command] --help for more details on each hook command")
		},
	}
	return cmd
}

func SetupHookCommitMsgCommand() *cobra.Command {
	return setupHookSubCommand(info.COMMIT_MSG)
}

func SetupHookPrepareCommitMsgCommand() *cobra.Command {
	return setupHookSubCommand(info.PREPARE_COMMIT_MSG)
}

func SetupHookPostCheckoutCommand() *cobra.Command {
	return setupHookSubCommand(info.POST_CHECKOUT)
}

func SetupHookPostCommitCommand() *cobra.Command {
	return setupHookSubCommand(info.POST_COMMIT)
}

func SetupHookPostMergeCommand() *cobra.Command {
	return setupHookSubCommand(info.POST_MERGE)
}

func SetupHookPostRewriteCommand() *cobra.Command {
	return setupHookSubCommand(info.POST_REWRITE)
}

func SetupHookPreCommitCommand() *cobra.Command {
	return setupHookSubCommand(info.PRE_COMMIT)
}

func SetupHookPrePushCommand() *cobra.Command {
	return setupHookSubCommand(info.PRE_PUSH)
}

func setupHookSubCommand(hook string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   hook,
		Short: "Execute " + hook + " actions",
		Long:  "Execute all actions configured for " + hook,
		Run: func(cmd *cobra.Command, args []string) {
			verbosity := setupGlobalFlags(cmd)

			appIO := io.NewDefaultIO(verbosity, mapArgs([]string{}, args))

			conf, repo, err := setUpConfigAndRepo(cmd)
			if err != nil {
				DisplayCommandError(err)
			}

			runner := exec.NewHookRunner(hook, appIO, conf, repo)
			errRun := runner.Run()
			if errRun != nil {
				DisplayCommandError(errRun)
			}
		},
	}

	configurationAware(cmd)
	repositoryAware(cmd)

	return cmd
}
