package branch

import (
	"errors"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"regexp"
)

// EnsureNaming prevents you from pushing to branches not following a given naming scheme.
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::Branch.EnsureNaming",
//	  "options": {
//	    "regex": "feature/[a-z]\-+[0-9]+"
//	  }
//	}
type EnsureNaming struct {
	hookBundle *hooks.HookBundle
}

func (a *EnsureNaming) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *EnsureNaming) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("ensure branch naming", true, io.VERBOSE)

	regex := action.Options().AsString("regex", "")
	if regex == "" {
		return errors.New("option 'regex' is missing")
	}

	errorMsg := action.Options().AsString("error-msg", "branch did not match '"+regex+"'")
	successMsg := action.Options().AsString("success-msg", "branch matches regex")

	branch := a.hookBundle.Repo.BranchName()
	match, _ := regexp.MatchString(regex, branch)
	if !match {
		return errors.New(errorMsg)
	}

	a.hookBundle.AppIO.Write(successMsg, true, io.VERBOSE)
	return nil
}

func NewEnsureNaming(appIO io.IO, conf *configuration.Configuration, repo git.Repo) hooks.Action {
	a := PreventPushOfFixupAndSquashCommits{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PrePush}),
	}
	return &a
}
