package message

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

type BeamsRules struct {
	hookBundle *hooks.HookBundle
}

func (a *BeamsRules) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *BeamsRules) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("checking beams rules", true, io.VERBOSE)
	return nil
}

func NewBeamsRules(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := BeamsRules{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.CommitMsg}),
	}
	return &a
}
