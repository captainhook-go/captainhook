package io

const (
	QUIET   = 1
	NORMAL  = 2
	VERBOSE = 4
	DEBUG   = 8
)

type IO interface {
	Verbosity() int
	Options() map[string]string
	Option(name, defaultValue string) string
	Arguments() map[string]string
	Argument(name, defaultValue string) string
	StandardInput() []string
	Input() Input
	IsInteractive() bool
	IsDebug() bool
	IsQuiet() bool
	IsVerbose() bool
	Write(message string, newline bool, verbosity int)
	Ask(message string, defaultValue string) string
}
