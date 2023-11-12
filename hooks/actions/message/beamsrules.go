package message

import (
	"errors"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/git/types"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"strings"
)

type BeamsRules struct {
	hookBundle *hooks.HookBundle
}

func (a *BeamsRules) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *BeamsRules) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("checking beams rules", true, io.VERBOSE)

	rulebook := a.setupRulebook(action)
	commitMessageFile := a.hookBundle.AppIO.Argument("file", "")
	if commitMessageFile == "" {
		return errors.New("commit message file argument is missing")
	}
	msg, err := a.hookBundle.Repo.CommitMessage(commitMessageFile)
	if err != nil {
		return err
	}

	ok, messages := rulebook.IsFollowedBy(msg)
	if !ok {
		for _, message := range messages {
			a.hookBundle.AppIO.Write(message, true, io.NORMAL)
		}
		a.outputMessage(msg)
		return errors.New("commit message did not follow all rules")
	}
	return nil
}

func (a *BeamsRules) setupRulebook(action *configuration.Action) *Rulebook {
	maxBodyLineLength := action.Options().AsInt("body-line-length", 72)
	maxSubjectLength := action.Options().AsInt("subject-length", 50)
	checkImperativeBeginningOnly := action.Options().AsBool("check-imperative-beginning-only", false)

	rulebook := NewRulebook()
	rulebook.AddRule(
		NewCapitalizeSubject(),
		NewLimitBodyLineLength(maxBodyLineLength),
		NewLimitSubjectLineLength(maxSubjectLength),
		NewMsgNotEmpty(),
		NewNoPeriodOnSubjectEnd(),
		NewSeparateSubjectFromBodyWithBlankLine(),
		NewUseImperativeMood(checkImperativeBeginningOnly),
	)
	return rulebook
}

func (a *BeamsRules) outputMessage(msg *types.CommitMessage) {
	a.hookBundle.AppIO.Write("===========================[ commit message ]===========================", true, io.NORMAL)
	a.hookBundle.AppIO.Write(msg.Message(), true, io.NORMAL)
	a.hookBundle.AppIO.Write(strings.Repeat("=", 72), true, io.NORMAL)
}

func NewBeamsRules(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := BeamsRules{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.CommitMsg}),
	}
	return &a
}
