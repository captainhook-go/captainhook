package hooks

import (
	"fmt"
	"github.com/captainhook-go/captainhook/cli"
	"github.com/spf13/cobra"
)

func SetupHookPostCommitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post-commit",
		Short: "Execute post-commit actions",
		Long:  "Execute all actions configured for post-commit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("POST COMMIT HOOK")
		},
	}

	cli.ConfigurationAware(cmd)
	cli.RepositoryAware(cmd)

	return cmd
}
