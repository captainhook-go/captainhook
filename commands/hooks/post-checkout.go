package hooks

import (
	"github.com/captainhook-go/captainhook/cli"
	"github.com/captainhook-go/captainhook/exec"
	"github.com/captainhook-go/captainhook/io"
	"github.com/spf13/cobra"
)

func SetupHookPostCheckoutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post-checkout",
		Short: "Execute post-checkout actions",
		Long:  "Execute all actions configured for post-checkout",
		Run: func(cmd *cobra.Command, args []string) {
			appIO := io.NewDefaultIO(io.NORMAL, cli.MapArgs([]string{}, args))

			conf, repo, err := cli.SetUpConfigAndRepo(cmd)
			if err != nil {
				cli.DisplayCommandError(err)
			}

			runner := exec.NewPostCheckoutRunner(appIO, conf, repo)
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
