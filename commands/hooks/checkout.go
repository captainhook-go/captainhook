package hooks

import (
	"fmt"
	"github.com/captainhook-go/captainhook/cli"
	"github.com/spf13/cobra"
)

func SetupHookPostCheckoutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post-checkout",
		Short: "Execute post-checkout actions",
		Long:  "Execute all actions configured for post-checkout",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("POST CHECKOUT HOOK")
		},
	}

	cli.ConfigurationAware(cmd)
	cli.RepositoryAware(cmd)

	return cmd
}
