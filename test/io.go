package test

import "github.com/captainhook-go/captainhook/io"

type IOMock struct {
	stdIn []string
	opts  map[string]string
	args  map[string]string
	Out   []string
}

func (inOut *IOMock) SetStdIn(input []string) {
	inOut.stdIn = input
}

func (inOut *IOMock) SetArguments(args map[string]string) {
	inOut.args = args
}

func (inOut *IOMock) SetOptions(args map[string]string) {
	inOut.opts = args
}

func (inOut *IOMock) Verbosity() int {
	return 0
}

func (inOut *IOMock) Options() map[string]string {
	return map[string]string{}
}

func (inOut *IOMock) Option(name, defaultValue string) string {
	val, ok := inOut.opts[name]
	if !ok {
		return defaultValue
	}
	return val
}

func (inOut *IOMock) Arguments() map[string]string {
	return map[string]string{}
}

func (inOut *IOMock) Argument(name, defaultValue string) string {
	val, ok := inOut.args[name]
	if !ok {
		return defaultValue
	}
	return val
}

func (inOut *IOMock) StandardInput() []string {
	return inOut.stdIn
}

func (inOut *IOMock) Input() io.Input {
	return io.NewStdIn(inOut.opts, inOut.args)
}

func (inOut *IOMock) IsInteractive() bool {
	return true
}

func (inOut *IOMock) IsDebug() bool {
	return true
}

func (inOut *IOMock) IsQuiet() bool {
	return true
}

func (inOut *IOMock) IsVerbose() bool {
	return true
}

func (inOut *IOMock) Write(message string, newline bool, verbosity int) {
	inOut.Out = append(inOut.Out, message)
}

func (inOut *IOMock) Ask(message string, defaultValue string) string {
	return ""
}
