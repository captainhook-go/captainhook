package commands

import (
	"github.com/captainhook-go/captainhook/cli"
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/exec"
	"github.com/captainhook-go/captainhook/git"
	"github.com/spf13/cobra"
)

func setupInstallCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install CaptainHook",
		Long:  "Install CaptainHook into your local .git/hooks directory",
		Run: func(cmd *cobra.Command, args []string) {

			confPath, _ := cmd.Flags().GetString("config")
			repoPath, _ := cmd.Flags().GetString("repository")

			settings := config.Settings{}
			conf, err := config.NewConfiguration(confPath, true, settings)
			if err != nil {
				cli.DisplayCommandError(err)
				return
			}
			repo, err := git.NewRepository(repoPath)
			if err != nil {
				cli.DisplayCommandError(err)
				return
			}
			exec.Install(conf, repo)
		},
	}

	cli.ConfigurationAware(cmd)
	cli.RepositoryAware(cmd)

	return cmd
}
