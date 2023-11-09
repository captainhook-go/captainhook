package git

import (
	"context"
	"github.com/captainhook-go/captainhook/git/types"
)

func command(ctx context.Context, name string, options ...types.Option) (string, error) {
	g := types.NewCmd(name)
	g.AddOptions(options...)

	return g.Exec(ctx, g.Command, g.Debug, g.Options...)
}

func SetExecutor(executor types.Executor) types.Option {
	return func(g *types.Cmd) {
		g.Executor = executor
	}
}

func SetDebug(debug bool) types.Option {
	return func(g *types.Cmd) {
		g.Debug = debug
	}
}

func Config(options ...types.Option) (string, error) {
	return command(context.Background(), "config", options...)
}
func DiffIndex(options ...types.Option) (string, error) {
	return command(context.Background(), "diff-index", options...)
}
func DiffTree(options ...types.Option) (string, error) {
	return command(context.Background(), "diff-tree", options...)
}
func Log(options ...types.Option) (string, error) {
	return command(context.Background(), "log", options...)
}
func RevParse(options ...types.Option) (string, error) {
	return command(context.Background(), "rev-parse", options...)
}
