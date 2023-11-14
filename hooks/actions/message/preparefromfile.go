package message

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

// PrepareFromFile loads a suggested commit message from a given file.
// It is the counterpart of CacheOnFail. If your commit failed because of an
// invalid message you can use the action to preload the failed message,
// so you can edit it without having to completely write it from scratch again.
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::Message.PrepareFromFile",
//	  "options": {
//	    "file": ".git/CH_COMMIT_MSG"
//	  }
//	}
type PrepareFromFile struct {
	hookBundle *hooks.HookBundle
}

func (a *PrepareFromFile) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *PrepareFromFile) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("prepare from file", true, io.VERBOSE)
	// TODO: implement this
	return nil
}

func NewPrepareFromFile(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := PrepareFromFile{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PrepareCommitMsg}),
	}
	return &a
}
