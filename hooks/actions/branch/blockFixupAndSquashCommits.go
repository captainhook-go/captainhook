package branch

import (
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

type PreventPushOfFixupAndSquashCommits struct {
	hookBundle *hooks.HookBundle
}

func (a *PreventPushOfFixupAndSquashCommits) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *PreventPushOfFixupAndSquashCommits) Run(action *config.Action) error {
	// check branch config
	// check log
	a.hookBundle.AppIO.Write("blocking fixup and squash commits", true, io.VERBOSE)
	return nil
}

func NewPreventPushOfFixupAndSquashCommits(appIO io.IO, conf *config.Configuration, repo *git.Repository) hooks.Action {
	a := PreventPushOfFixupAndSquashCommits{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PRE_PUSH}),
	}
	return &a
}
