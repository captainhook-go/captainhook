package io

type NullIO struct {
	verbosity int
	arguments map[string]string
}

func (n *NullIO) Arguments() map[string]string {
	myMap := map[string]string{}
	return myMap
}
func (n *NullIO) Argument(name string) string {
	return ""
}

func (n *NullIO) IsQuiet() bool {
	return true
}

func (n *NullIO) IsDebug() bool {
	return false
}

func (n *NullIO) IsVeryVerbose() bool {
	return false
}

func (n *NullIO) IsVerbose() bool {
	return false
}

func (n *NullIO) Write(message string, verbosity int) {
}
