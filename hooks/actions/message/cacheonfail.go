package message

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

type CacheOnFail struct {
	hookBundle *hooks.HookBundle
}

func (a *CacheOnFail) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *CacheOnFail) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("checking regex", true, io.VERBOSE)
	return nil
}

func NewCacheOnFail(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := CacheOnFail{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.CommitMsg}),
	}
	return &a
}