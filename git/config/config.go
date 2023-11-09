package config

import "github.com/captainhook-go/captainhook/git/types"

func Get(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOption("--get")
		g.AddOption(name)
	}
}
