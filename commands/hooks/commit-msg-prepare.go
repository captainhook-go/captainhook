package hooks

import (
	"github.com/captainhook-go/captainhook/cli"
	"github.com/captainhook-go/captainhook/exec"
	"github.com/captainhook-go/captainhook/io"
	"github.com/spf13/cobra"
)

func SetupHookPrepareCommitMsgCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prepare-commit-msg",
		Short: "Execute pre-push actions",
		Long:  "Execute all actions configured for pre-push",
		Run: func(cmd *cobra.Command, args []string) {
			appIO := io.NewDefaultIO(io.NORMAL, cli.MapArgs([]string{"file", "mode", "hash"}, args))

			confPath, _ := cmd.Flags().GetString("config")
			repoPath, _ := cmd.Flags().GetString("repository")

			conf, repo, err := cli.SetUpConfigAndRepo(confPath, repoPath)
			if err != nil {
				cli.DisplayCommandError(err)
			}

			runner := exec.NewPrepareCommitMsgRunner(appIO, conf, repo)
			errRun := runner.Run()
			if errRun != nil {
				cli.DisplayCommandError(errRun)
			}
		},
	}

	cli.ConfigurationAware(cmd)
	cli.RepositoryAware(cmd)

	return cmd
}
