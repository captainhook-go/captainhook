package io

const (
	QUIET        = 1
	NORMAL       = 2
	VERBOSE      = 4
	VERY_VERBOSE = 8
	DEBUG        = 16
)

type IO interface {
	GetArguments() map[string]string
	GetArgument(arg string) string
	GetStandardInput() []string
	IsDebug() bool
	IsVeryVerbose() bool
	IsVerbose() bool
	IsInteractive() bool
	Write(message string, verbosity int)
}
