package file

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

type MaxSize struct {
	hookBundle *hooks.HookBundle
}

func (a *MaxSize) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *MaxSize) Run(action *configuration.Action) error {
	// read configured file
	// check if file is empty
	a.hookBundle.AppIO.Write("checking max file size", true, io.VERBOSE)
	return nil
}

func NewMaxSize(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := MaxSize{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
	return &a
}
