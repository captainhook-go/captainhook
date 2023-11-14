package revparse

import "github.com/captainhook-go/captainhook/git/types"

func AbbrevRef(g *types.Cmd) {
	g.AddOption("--abbrev-ref")
}

func ShowTopLevel(g *types.Cmd) {
	g.AddOption("--show-toplevel")
}
