package types

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

var (
	defaultExecutor = func(ctx context.Context, name string, debug bool, args ...string) (string, error) {
		if debug {
			fmt.Println(name, strings.Join(args, " "))
		}

		output, err := exec.CommandContext(ctx, name, args...).CombinedOutput()

		return string(output), err
	}
)

type Executor func(ctx context.Context, name string, debug bool, args ...string) (string, error)

type Option func(g *Cmd)

type Cmd struct {
	Debug    bool
	Command  string
	Options  []string
	Executor Executor
}

func (g *Cmd) AddOption(option string) {
	g.Options = append(g.Options, option)
}

func (g *Cmd) AddOptions(options ...Option) {
	for _, opt := range options {
		opt(g)
	}
}

func (g *Cmd) Exec(ctx context.Context, name string, debug bool, args ...string) (string, error) {
	return g.Executor(ctx, name, debug, args...)
}

func NewCmd(name string) *Cmd {
	g := &Cmd{
		Debug:   false,
		Command: "git",
		Options: []string{name},
	}
	g.Executor = defaultExecutor
	return g
}

func SetDefaultExecutor(exec Executor) {
	defaultExecutor = exec
}
