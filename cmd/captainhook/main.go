/*
CaptainHook is a git hook manager to easily configure and share your git hooks with teammates.

The Cap'n works with a json configuration file where you configure actions that should
get executed during git hooks. The Cap'n has a lot of build-in features like commit message validation
or only executing actions if some conditions apply.

To set up the Cap'n you have to run the `init` command to create a configuration
file followed by the `install` command to activate your local git hooks.

Once you committed the config file to the repository your teammates just have to install the Cap'n
and to run `captainhook install` and they all have the same git hook setup.

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
