/*
CaptainHook is a git hook manager and can be used to easily share git hooks with teammates.

The Cap'n works with a json configuration file to easily configure actions that should
get executed during git hooks. It offers a lot of features like commit message validation
or only executing actions if some conditions apply.

Once you committed the config file to the repository all your teammates have
to do is run `captainhook install` and they all have the same git hook setup.

To use the Cap'n you have to run the `init` and the `install` command.

Usage:

	captainhook [commands]

The commands are:

	init
		Creates an empty captainhook.json configuration. You should
		call this command in your project repository root.
	install
		This will activate captainhook by installing hooks into your local
		.git/hooks/ directory. After executing this the captain will be
		triggered by git hooks and execute everything you configured.

The Cap'n assumes you store your `captainhook.json` in the repository root.
You can change that by moving the configuration to a subdirectory and then
use the `--configuration=my/sub/dir/captainhook.json` option to link the
correct configuration.
*/
package main

import (
	"github.com/captainhook-go/captainhook/commands"
	"os"
)

func main() {
	resp := commands.Execute(os.Args[1:])

	if resp.Err != nil {
		if resp.IsUserError() {
			resp.Cmd.Println("")
			resp.Cmd.Println(resp.Cmd.UsageString())
		}
		os.Exit(-1)
	}
}
