package commands

import (
	"github.com/captainhook-go/captainhook/exec"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/io"
	"github.com/spf13/cobra"
	"os"
)

func setupUninstallCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Deletes hooks from your local .git/hooks directory",
		Long:  "Deletes hooks from your local .git/hooks directory",
		Run: func(cmd *cobra.Command, args []string) {
			force, _ := cmd.Flags().GetBool("force")
			backup, _ := cmd.Flags().GetBool("backup")

			conf, err := setUpConfig(cmd)
			if err != nil {
				DisplayCommandError(err)
			}

			repo, errRepo := git.NewRepository(conf.GitDirectory())
			if errRepo != nil {
				DisplayCommandError(errRepo)
			}

			appIO := io.NewDefaultIO(conf.Verbosity(), make(map[string]string))

			unInstaller := exec.NewUninstaller(appIO, conf, repo)
			unInstaller.Force(force)
			unInstaller.EnableBackup(backup)
			unInstError := unInstaller.Run()

			if unInstError != nil {
				os.Exit(1)
			}
		},
	}

	setUpUninstallFlags(cmd)
	configurationAware(cmd)
	repositoryAware(cmd)

	return cmd
}

func setUpUninstallFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("force", "f", false, "do not ask for confirmation")
	cmd.Flags().BoolP("backup", "b", false, "backup existing hooks to .git/hooks/{HOOKFILE}.bak")
}
