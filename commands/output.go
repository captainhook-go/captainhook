package commands

import "os"

func DisplayCommandError(err error) {
	println("\n====== ERROR ======")
	println(err.Error())
	os.Exit(1)
}
