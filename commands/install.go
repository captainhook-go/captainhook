package commands

import (
	"github.com/captainhook-go/captainhook/cli"
	"github.com/captainhook-go/captainhook/exec"
	"github.com/captainhook-go/captainhook/io"
	"github.com/spf13/cobra"
)

func setupInstallCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install CaptainHook",
		Long:  "Install CaptainHook into your local .git/hooks directory",
		Run: func(cmd *cobra.Command, args []string) {
			appIO := io.NewDefaultIO(io.NORMAL, make(map[string]string))

			confPath, _ := cmd.Flags().GetString("config")
			repoPath, _ := cmd.Flags().GetString("repository")
			force, _ := cmd.Flags().GetBool("force")
			skip, _ := cmd.Flags().GetBool("skip-existing")

			conf, repo, err := cli.SetUpConfigAndRepo(confPath, repoPath)
			if err != nil {
				cli.DisplayCommandError(err)
			}

			installer := exec.NewInstaller(appIO, conf, repo)
			installer.SkipExisting(skip)
			installer.Force(force)
			instError := installer.Run()
			if instError != nil {
				cli.DisplayCommandError(instError)
			}
		},
	}

	setUpFlags(cmd)
	cli.ConfigurationAware(cmd)
	cli.RepositoryAware(cmd)

	return cmd
}

func setUpFlags(cmd *cobra.Command) {
	var skip = false
	cmd.Flags().BoolP("skip-existing", "s", skip, "skip existing hooks")
	var force = false
	cmd.Flags().BoolP("force", "f", force, "force installation, overwrite existing hooks")
}
