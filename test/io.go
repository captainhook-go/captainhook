package test

import "github.com/captainhook-go/captainhook/io"

type IOMock struct {
}

func (InOut *IOMock) Verbosity() int {
	return 0
}

func (InOut *IOMock) Arguments() map[string]string {
	return map[string]string{}
}

func (InOut *IOMock) Argument(name, defaultValue string) string {
	return ""
}

func (InOut *IOMock) StandardInput() []string {
	return []string{}
}

func (InOut *IOMock) Input() io.Input {
	return io.NewStdIn(map[string]string{})
}

func (InOut *IOMock) IsInteractive() bool {
	return true
}

func (InOut *IOMock) IsDebug() bool {
	return true
}

func (InOut *IOMock) IsQuiet() bool {
	return true
}

func (InOut *IOMock) IsVerbose() bool {
	return true
}

func (InOut *IOMock) Write(message string, newline bool, verbosity int) {

}

func (InOut *IOMock) Ask(message string, defaultValue string) string {
	return ""
}
