package io

import "fmt"

type DefaultIO struct {
	verbosity int
	arguments map[string]string
}

func NewDefaultIO(verbosity int, arguments map[string]string) *DefaultIO {
	io := DefaultIO{verbosity: verbosity, arguments: arguments}
	return &io
}

func (d *DefaultIO) Arguments() map[string]string {
	myMap := map[string]string{}
	return myMap
}
func (d *DefaultIO) Argument(name string) string {
	return d.arguments[name]
}

func (d *DefaultIO) IsQuiet() bool {
	return !(d.verbosity > QUIET)
}

func (d *DefaultIO) IsDebug() bool {
	return d.verbosity >= DEBUG
}

func (d *DefaultIO) IsVeryVerbose() bool {
	return d.verbosity >= VERY_VERBOSE
}

func (d *DefaultIO) IsVerbose() bool {
	return d.verbosity >= VERBOSE
}

func (d *DefaultIO) Write(message string, verbosity int) {
	if d.IsQuiet() {
		return
	}
	if verbosity <= d.verbosity {
		fmt.Print(message)
	}
}
