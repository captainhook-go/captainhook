package message

import (
	"errors"
	"fmt"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"regexp"
)

type ContainsRegex struct {
	hookBundle *hooks.HookBundle
}

func (a *ContainsRegex) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *ContainsRegex) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("checking regex", true, io.VERBOSE)

	commitMessageFile := a.hookBundle.AppIO.Argument("file", "")
	msg, err := a.hookBundle.Repo.CommitMessage(commitMessageFile)
	if err != nil {
		return err
	}
	regex := action.Options().AsString("regex", "")
	if regex == "" {
		return errors.New("option 'regex' is missing")
	}
	match, _ := regexp.MatchString(regex, msg.Message())
	if !match {
		return fmt.Errorf("unable to find '%s' in commit message", regex)
	}
	return nil
}

func NewContainsRegex(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := ContainsRegex{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.CommitMsg}),
	}
	return &a
}
