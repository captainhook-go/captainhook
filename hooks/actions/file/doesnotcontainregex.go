package file

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

type DoesNotContainRegex struct {
	hookBundle *hooks.HookBundle
}

func (a *DoesNotContainRegex) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *DoesNotContainRegex) Run(action *configuration.Action) error {
	// read configured file
	// check regex
	a.hookBundle.AppIO.Write("checking if file does not contain regex", true, io.VERBOSE)
	return nil
}

func NewDoesNotContainRegex(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := DoesNotContainRegex{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
	return &a
}
