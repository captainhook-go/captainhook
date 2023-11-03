package commands

import (
	"fmt"
	"github.com/captainhook-go/captainhook/exec"
	"github.com/captainhook-go/captainhook/git"
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
			out := "Usage:\n" +
				"captainhook hook [command] [options]\n" +
				"\n" +
				"Available Commands:\n"
			fmt.Print(out)
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
	return setupHookSubCommand(info.COMMIT_MSG, []string{})
}

func SetupHookPrepareCommitMsgCommand() *cobra.Command {
	return setupHookSubCommand(info.PREPARE_COMMIT_MSG, []string{})
}

func SetupHookPostCheckoutCommand() *cobra.Command {
	return setupHookSubCommand(info.POST_CHECKOUT, []string{})
}

func SetupHookPostCommitCommand() *cobra.Command {
	return setupHookSubCommand(info.POST_COMMIT, []string{})
}

func SetupHookPostMergeCommand() *cobra.Command {
	return setupHookSubCommand(info.POST_MERGE, []string{})
}

func SetupHookPostRewriteCommand() *cobra.Command {
	return setupHookSubCommand(info.POST_REWRITE, []string{})
}

func SetupHookPreCommitCommand() *cobra.Command {
	return setupHookSubCommand(info.PRE_COMMIT, []string{})
}

func SetupHookPrePushCommand() *cobra.Command {
	return setupHookSubCommand(info.PRE_PUSH, []string{})
}

func setupHookSubCommand(hook string, argMap []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   hook,
		Short: "Execute " + hook + " actions",
		Long:  "Execute all actions configured for " + hook,
		Run: func(cmd *cobra.Command, args []string) {
			conf, err := setUpConfig(cmd)
			if err != nil {
				DisplayCommandError(err)
			}

			repo, errRepo := git.NewRepository(conf.GitDirectory())
			if errRepo != nil {
				DisplayCommandError(errRepo)
			}

			io.ColorStatus(conf.AnsiColors())
			appIO := io.NewDefaultIO(conf.Verbosity(), mapArgs(argMap, args))
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
