package commands

import "os"

func DisplayCommandError(err error) {
	println("====== ERROR ======")
	println(err.Error())
	os.Exit(1)
}
