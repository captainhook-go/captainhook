package branch

import (
	"github.com/captainhook-go/captainhook/configuration"
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

func (a *PreventPushOfFixupAndSquashCommits) Run(action *configuration.Action) error {
	// check branch config
	// check log
	a.hookBundle.AppIO.Write("blocking fixup and squash commits", true, io.VERBOSE)
	return nil
}

func NewPreventPushOfFixupAndSquashCommits(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := PreventPushOfFixupAndSquashCommits{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PrePush}),
	}
	return &a
}
