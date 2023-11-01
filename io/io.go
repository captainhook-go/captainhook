package io

const (
	QUIET   = 1
	NORMAL  = 2
	VERBOSE = 4
	DEBUG   = 8
)

type IO interface {
	Verbosity() int
	Arguments() map[string]string
	Argument(arg string) string
	StandardInput() []string
	IsInteractive() bool
	Write(message string, newline bool, verbosity int)
}
