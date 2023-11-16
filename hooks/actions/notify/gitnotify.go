package notify

import (
	"errors"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/git/types"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/input"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"math"
	"regexp"
	"strings"
)

// GitNotify checks the git history for git-notify: messages and displays the notifications.
// Applicable for: 'post-checkout', 'post-merge', 'post-rewrite'
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::Notify.GitNotify",
//	  "options": {
//	    "prefix": "git-notify:"
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
	prefix := action.Options().AsString("prefix", "git-notify:")
	ranges := input.DetectRanges(a.hookBundle.AppIO)

	for _, r := range ranges {
		commits := a.hookBundle.Repo.CommitsBetween(r.From().Id(), r.To().Id())
		for _, c := range commits {
			if a.containsNotification(c, prefix) {
				msg, _ := a.extractNotification(c, prefix)
				a.showNotification(msg)
			}
		}
	}
	return nil
}

func (a *GitNotify) containsNotification(c *types.Commit, prefix string) bool {
	return strings.Contains(c.Subject+c.Body, prefix)
}

func (a *GitNotify) extractNotification(c *types.Commit, prefix string) (string, error) {
	r, regErr := regexp.Compile("(?m)" + prefix + "(.*)")
	if regErr != nil {
		return "", errors.New("invalid regex pattern")
	}

	match := r.FindStringSubmatch(c.Subject + c.Body)
	if len(match) < 1 {
		return "", nil
	}
	return strings.TrimSpace(match[1]), nil
}

func (a *GitNotify) showNotification(msg string) {
	headline := "+" + strings.Repeat("-", 28) + "[ git-notify ]" + strings.Repeat("-", 28) + "+"
	blank := "|" + strings.Repeat(" ", 70) + "|"
	footer := "+" + strings.Repeat("-", 70) + "+"

	a.hookBundle.AppIO.Write(headline, true, io.NORMAL)
	a.hookBundle.AppIO.Write(blank, true, io.NORMAL)

	for _, m := range io.SplitLines(msg) {
		trimmed := strings.TrimSpace(m)
		lPad := int(math.Floor(float64(68-len(trimmed)) / 2))
		rPad := int(math.Ceil(float64(68-len(trimmed)) / 2))
		padL := strings.Repeat(" ", lPad)
		padR := strings.Repeat(" ", rPad)
		line := "| " + padL + "<comment>" + trimmed + "</comment>" + padR + " |"
		a.hookBundle.AppIO.Write(line, true, io.NORMAL)
	}

	a.hookBundle.AppIO.Write(blank, true, io.NORMAL)
	a.hookBundle.AppIO.Write(footer, true, io.NORMAL)
}

func NewGitNotify(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := GitNotify{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PostRewrite, info.PostCheckout, info.PostMerge}),
	}
	return &a
}
