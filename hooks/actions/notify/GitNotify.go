package notify

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

// GitNotify checks the git history for git-notify: messages if your working directory gets changed.
// Applicable for: 'post-checkout', 'post-merge', 'post-rewrite'
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::Message.PrepareFromFile",
//	  "options": {
//	    "file": ".git/CH_COMMIT_MSG"
//	  }
//	}
type GitNotify struct {
	hookBundle *hooks.HookBundle
}

func (a *GitNotify) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *GitNotify) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("check history for notifications", true, io.VERBOSE)
	// TODO: implement this
	//oldHash := ""
	//newHash := ""
	// get commit between oldHash newHash
	// check all commit messages for "git-notify: .*"
	// if found output notification
	return nil
}

func NewGitNotify(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := GitNotify{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PostRewrite, info.PostCheckout, info.PostMerge}),
	}
	return &a
}
