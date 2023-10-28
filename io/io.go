package io

const (
	QUIET        = 1
	NORMAL       = 2
	VERBOSE      = 4
	VERY_VERBOSE = 8
	DEBUG        = 16
)

type IO interface {
	Arguments() map[string]string
	Argument(arg string) string
	StandardInput() []string
	IsQuiet() bool
	IsDebug() bool
	IsVeryVerbose() bool
	IsVerbose() bool
	IsInteractive() bool
	Write(message string, newline bool, verbosity int)
}
