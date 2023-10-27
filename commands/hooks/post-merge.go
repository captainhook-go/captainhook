package hooks

import (
	"fmt"
	"github.com/captainhook-go/captainhook/cli"
	"github.com/spf13/cobra"
)

func SetupHookPostMergeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post-merge",
		Short: "Execute post-merge actions",
		Long:  "Execute all actions configured for post-merge",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("POST MERGE HOOK")
		},
	}

	cli.ConfigurationAware(cmd)
	cli.RepositoryAware(cmd)

	return cmd
}
