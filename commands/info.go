package commands

import (
	"github.com/captainhook-go/captainhook/exec"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/io"
	"github.com/spf13/cobra"
	"os"
)

func setupInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Displays configuration details",
		Long:  "Displays configuration details",
		Run: func(cmd *cobra.Command, args []string) {
			listActions, _ := cmd.Flags().GetBool("list-actions")
			listConditions, _ := cmd.Flags().GetBool("list-conditions")
			listOptions, _ := cmd.Flags().GetBool("list-options")
			extended, _ := cmd.Flags().GetBool("extended")

			conf, err := setUpConfig(cmd)
			if err != nil {
				DisplayCommandError(err)
			}

			repo, errRepo := git.NewRepository(conf.GitDirectory())
			if errRepo != nil {
				DisplayCommandError(errRepo)
			}

			appIO := io.NewDefaultIO(conf.Verbosity(), map[string]string{}, map[string]string{})

			info := exec.NewConfigInfo(appIO, conf, repo)
			info.Display("actions", listActions)
			info.Display("conditions", listConditions)
			info.Display("options", listOptions)
			info.Extended(extended)
			instError := info.Run()

			if instError != nil {
				os.Exit(1)
			}
		},
	}

	setUpInfoFlags(cmd)
	configurationAware(cmd)
	repositoryAware(cmd)

	return cmd
}

func setUpInfoFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("list-actions", "a", false, "list actions")
	cmd.Flags().BoolP("list-conditions", "p", false, "list all conditions")
	cmd.Flags().BoolP("list-options", "o", false, "list all options and args")
	cmd.Flags().BoolP("extended", "e", false, "show detailed information")
}
