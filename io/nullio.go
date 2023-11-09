package io

type NullIO struct {
	verbosity int
	arguments map[string]string
}

func (n *NullIO) Verbosity() int {
	return n.verbosity
}

func (n *NullIO) Arguments() map[string]string {
	myMap := map[string]string{}
	return myMap
}

func (n *NullIO) Argument(name, defaultValue string) string {
	return ""
}

func (n *NullIO) StandardInput() []string {
	return []string{}
}

func (n *NullIO) IsInteractive() bool {
	return false
}

func (n *NullIO) IsDebug() bool {
	return false
}

func (n *NullIO) IsVerbose() bool {
	return false
}

func (n *NullIO) Write(message string, newLine bool, verbosity int) {
}
