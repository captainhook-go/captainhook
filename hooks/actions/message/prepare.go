package message

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/git/types"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

// Prepare is used to prefill your commit messages with some defaults.
// It only works if you commit message is empty so far.
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::Message.Prepare",
//	  "options": {
//	    "subject": "My prepared subject",
//	    "body": "My prepared body"
//	  }
//	}
type Prepare struct {
	hookBundle *hooks.HookBundle
}

func (a *Prepare) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *Prepare) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("prepare commit message", true, io.VERBOSE)

	msgFile := a.hookBundle.AppIO.Argument(info.ArgCommitMsgFile, "")
	msg, err := a.hookBundle.Repo.CommitMessage(msgFile)
	if err != nil {
		return err
	}
	if !msg.IsEmpty() {
		a.hookBundle.AppIO.Write("commit message not empty", true, io.VERBOSE)
		return nil
	}
	return a.hookBundle.Repo.PrepareCommitMessage(msgFile, a.createCommitMessage(action, msg))
}

func (a *Prepare) createCommitMessage(action *configuration.Action, msg *types.CommitMessage) *types.CommitMessage {
	subject := action.Options().AsString("subject", "")
	body := action.Options().AsString("body", "")

	newMsg := subject
	if body != "" {
		newMsg = newMsg + "\n\n" + body
	}

	return types.NewCommitMessage(newMsg, msg.CommentChar())
}

func NewPrepare(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := PrepareFromFile{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PrepareCommitMsg}),
	}
	return &a
}
