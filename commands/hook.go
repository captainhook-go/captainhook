package commands

import (
	"fmt"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/spf13/cobra"
	"strings"
)

func setupHookCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hook",
		Short: "Execute all actions registered for a git hook",
		Long:  "Execute all actions registered for a git hook",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Usage:")
			fmt.Println("captainhook hook [command] [options]\n")
			fmt.Println("Available Commands:")
			for _, hookName := range hooks.GetNativeHooks() {
				spaces := strings.Repeat(" ", 19-len(hookName)+2)
				fmt.Printf("  %s %sExecute %s actions\n", hookName, spaces, hookName)
			}
			fmt.Println("\nrun captainhook hook [command] --help for more details on each hook command")
		},
	}
	return cmd
}
