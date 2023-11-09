package file

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

type IsNotEmpty struct {
	hookBundle *hooks.HookBundle
}

func (a *IsNotEmpty) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *IsNotEmpty) Run(action *configuration.Action) error {
	// read configured file
	// check if file is empty
	a.hookBundle.AppIO.Write("checking if file is empty", true, io.VERBOSE)
	return nil
}

func NewIsNotEmpty(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := IsNotEmpty{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
	return &a
}
