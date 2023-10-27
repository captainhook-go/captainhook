package cli

func DisplayCommandError(err error) {
	println("====== ERROR ======")
	println(err.Error())
}
