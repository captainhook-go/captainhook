package commands

import (
	"github.com/captainhook-go/captainhook/io"
	"os"
)

func DisplayCommandError(err error) {
	println(io.Colorize("<warning>" + err.Error() + "</warning>"))
	os.Exit(1)
}
