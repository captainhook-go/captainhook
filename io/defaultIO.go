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

func (d *DefaultIO) isQuiet() bool {
	return !(d.verbosity > QUIET)
}

func (d *DefaultIO) IsDebug() bool {
	return d.verbosity == DEBUG
}

func (d *DefaultIO) IsVerbose() bool {
	return d.verbosity == VERBOSE
}

func (d *DefaultIO) Write(message string, newline bool, verbosity int) {
	if d.isQuiet() {
		return
	}
	var linebreak = ""
	if newline {
		linebreak = "\n"
	}

	if d.verbosity >= verbosity {
		fmt.Print(Colorize(message + linebreak))
	}
}

func (d *DefaultIO) Ask(message string, defaultValue string) string {
	value, err := getUserInput(message)
	if err != nil {
		value = defaultValue
	}
	if len(value) == 0 {
		value = defaultValue
	}
	return value
}
