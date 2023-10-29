package hooks

import (
	"github.com/captainhook-go/captainhook/cli"
	"github.com/captainhook-go/captainhook/exec"
	"github.com/captainhook-go/captainhook/io"
	"github.com/spf13/cobra"
)

func SetupHookPreCommitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pre-commit",
		Short: "Execute pre-commit actions",
		Long:  "Execute all actions configured for pre-commit",
		Run: func(cmd *cobra.Command, args []string) {
			appIO := io.NewDefaultIO(io.NORMAL, cli.MapArgs([]string{}, args))

			conf, repo, err := cli.SetUpConfigAndRepo(cmd)
			if err != nil {
				cli.DisplayCommandError(err)
			}

			runner := exec.NewPreCommitRunner(appIO, conf, repo)
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
