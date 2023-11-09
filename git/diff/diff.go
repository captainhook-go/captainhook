package diff

import "github.com/captainhook-go/captainhook/git/types"

func Algorithm(algo string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOption("--diff-algorithm=" + algo)
	}
}

func Cached(g *types.Cmd) {
	g.AddOption("--cached")
}

func Filter(filter string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOption("--diff-filter=" + filter)
	}
}

func FromTo(from, to string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOption(from)
		g.AddOption(to)
	}
}

func NoExtDiff(g *types.Cmd) {
	g.AddOption("--no-ext-diff")
}

func Recursive(g *types.Cmd) {
	g.AddOption("-r")
}

func To(to string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOption(to)
	}
}
