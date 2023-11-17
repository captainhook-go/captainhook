package commands

import (
	"fmt"
	"github.com/captainhook-go/captainhook/exec"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"github.com/spf13/cobra"
	"os"
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
	return setupHookSubCommand(info.CommitMsg)
}

func SetupHookPrepareCommitMsgCommand() *cobra.Command {
	return setupHookSubCommand(info.PrepareCommitMsg)
}

func SetupHookPostCheckoutCommand() *cobra.Command {
	return setupHookSubCommand(info.PostCheckout)
}

func SetupHookPostCommitCommand() *cobra.Command {
	return setupHookSubCommand(info.PostCommit)
}

func SetupHookPostMergeCommand() *cobra.Command {
	return setupHookSubCommand(info.PostMerge)
}

func SetupHookPostRewriteCommand() *cobra.Command {
	return setupHookSubCommand(info.PostRewrite)
}

func SetupHookPreCommitCommand() *cobra.Command {
	return setupHookSubCommand(info.PreCommit)
}

func SetupHookPrePushCommand() *cobra.Command {
	return setupHookSubCommand(info.PrePush)
}

func setupHookSubCommand(hook string) *cobra.Command {
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
			appIO := io.NewDefaultIO(conf.Verbosity(), mapArgs(info.HookArguments(hook), args, hook))
			runner := exec.NewHookRunner(hook, appIO, conf, repo)
			errRun := runner.Run()
			if errRun != nil {
				os.Exit(1)
			}
		},
	}

	configurationAware(cmd)
	repositoryAware(cmd)

	return cmd
}
