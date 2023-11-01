package io

import (
	"fmt"
)

type DefaultIO struct {
	verbosity int
	arguments map[string]string
}

func NewDefaultIO(verbosity int, arguments map[string]string) *DefaultIO {
	io := DefaultIO{verbosity: verbosity, arguments: arguments}
	return &io
}

func (d *DefaultIO) Verbosity() int {
	return d.verbosity
}

func (d *DefaultIO) Arguments() map[string]string {
	return d.arguments
}
func (d *DefaultIO) Argument(name string) string {
	value, ok := d.arguments[name]
	if !ok {
		value = ""
	}
	return value
}

func (d *DefaultIO) StandardInput() []string {
	return make([]string, 0)
}

func (d *DefaultIO) IsInteractive() bool {
	return false
}

func (d *DefaultIO) Write(message string, newline bool, verbosity int) {
	if d.isQuiet() {
		return
	}
	var linebreak = ""
	if newline {
		linebreak = "\n"
	}

	if verbosity >= d.verbosity {
		fmt.Print(Colorize(message + linebreak))
	}
}

func (d *DefaultIO) isQuiet() bool {
	return !(d.verbosity > QUIET)
}
