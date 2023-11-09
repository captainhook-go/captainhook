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
	Argument(name, defaultValue string) string
	StandardInput() []string
	IsInteractive() bool
	IsDebug() bool
	IsVerbose() bool
	Write(message string, newline bool, verbosity int)
	Ask(message string, defaultValue string) string
}
