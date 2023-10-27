package hooks

import (
	"fmt"
	"github.com/captainhook-go/captainhook/cli"
	"github.com/spf13/cobra"
)

func SetupHookPreCommitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pre-commit",
		Short: "Execute pre-commit actions",
		Long:  "Execute all actions configured for pre-commit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("PRE COMMIT HOOK")
		},
	}

	cli.ConfigurationAware(cmd)
	cli.RepositoryAware(cmd)

	return cmd
}
