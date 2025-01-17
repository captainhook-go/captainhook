package commands

import (
	"github.com/captainhook-go/captainhook/exec"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/io"
	"github.com/spf13/cobra"
	"os"
)

func setupInstallCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install hooks into your local .git/hooks directory",
		Long:  "Install hooks into your local .git/hooks directory",
		Run: func(cmd *cobra.Command, args []string) {
			force, _ := cmd.Flags().GetBool("force")
			skip, _ := cmd.Flags().GetBool("skip-existing")
			onlyEnabled, _ := cmd.Flags().GetBool("only-enabled")

			conf, err := setUpConfig(cmd, true)
			if err != nil {
				DisplayCommandError(err)
			}

			repo, errRepo := git.NewRepository(conf.GitDirectory())
			if errRepo != nil {
				DisplayCommandError(errRepo)
			}

			appIO := io.NewDefaultIO(conf.Verbosity(), map[string]string{}, map[string]string{})

			installer := exec.NewInstaller(appIO, conf, repo)
			installer.SkipExisting(skip)
			installer.OnlyEnabled(onlyEnabled)
			installer.Force(force)
			instError := installer.Run()

			if instError != nil {
				os.Exit(1)
			}
		},
	}

	setUpInstallFlags(cmd)
	configurationAware(cmd)
	repositoryAware(cmd)

	return cmd
}

func setUpInstallFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("skip-existing", "s", false, "skip existing hooks")
	cmd.Flags().BoolP("force", "f", false, "force installation, overwrite existing hooks")
	cmd.Flags().BoolP("only-enabled", "e", false, "install only enabled hooks")
}
