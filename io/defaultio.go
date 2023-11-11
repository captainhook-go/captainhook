package io

import (
	"bufio"
	"fmt"
	"os"
)

type DefaultIO struct {
	verbosity int
	arguments map[string]string
	stdIn     []string
	stdInRead bool
}

func NewDefaultIO(verbosity int, arguments map[string]string) *DefaultIO {
	io := DefaultIO{verbosity: verbosity, arguments: arguments, stdInRead: false}
	return &io
}

func (d *DefaultIO) Verbosity() int {
	return d.verbosity
}

func (d *DefaultIO) Arguments() map[string]string {
	return d.arguments
}
func (d *DefaultIO) Argument(name, defaultValue string) string {
	value, ok := d.arguments[name]
	if !ok {
		value = defaultValue
	}
	return value
}

func (d *DefaultIO) StandardInput() []string {
	if !d.stdInRead {
		var lines []string
		if isPiped() {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				return []string{}
			}
		}
		d.stdInRead = true
		d.stdIn = lines
	}

	return d.stdIn
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
	value, err := askForUserInput(message)
	if err != nil {
		value = defaultValue
	}
	if len(value) == 0 {
		value = defaultValue
	}
	return value
}
