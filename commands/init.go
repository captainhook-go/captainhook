package commands

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/exec"
	"github.com/captainhook-go/captainhook/io"
	"github.com/spf13/cobra"
	"os"
)

func setupInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Creates a CaptainHook configuration file",
		Long:  "Creates a CaptainHook configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			force, _ := cmd.Flags().GetBool("force")
			config, _ := cmd.Flags().GetString("configuration")

			appIO := io.NewDefaultIO(configuration.MapVerbosity(getVerbosity(cmd)), map[string]string{}, make(map[string]string))

			initializer := exec.NewInitializer(appIO)
			initializer.UseConfig(config)
			initializer.Force(force)
			initError := initializer.Run()

			if initError != nil {
				os.Exit(1)
			}
		},
	}

	setUpInitFlags(cmd)
	configurationAware(cmd)

	return cmd
}

func setUpInitFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("force", "f", false, "force initialization, overwrite existing configuration")
}
