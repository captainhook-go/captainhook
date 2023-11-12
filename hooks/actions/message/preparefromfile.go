package message

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

type PrepareFromFile struct {
	hookBundle *hooks.HookBundle
}

func (a *PrepareFromFile) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *PrepareFromFile) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("prepare from file", true, io.VERBOSE)
	return nil
}

func NewPrepareFromFile(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := PrepareFromFile{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PrepareCommitMsg}),
	}
	return &a
}
