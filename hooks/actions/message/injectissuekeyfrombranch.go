package message

import (
	"errors"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/git/types"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"regexp"
	"strings"
)

// InjectIssueKeyFromBranch loads a suggested commit message from a given file.
// It is the counterpart of CacheOnFail. If your commit failed because of an
// invalid message you can use the action to preload the failed message,
// so you can edit it without having to completely write it from scratch again.
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::Message.InjectIssueKeyFromBranch",
//	  "options": {
//	    "regex": "([A-Z]+\\-[0-9]+)",
//	    "into": "body",
//	    "mode": "append",
//	    "prefix": "issue: ",
//	    "force": true
//	  }
//	}
type InjectIssueKeyFromBranch struct {
	hookBundle *hooks.HookBundle
}

func (a *InjectIssueKeyFromBranch) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *InjectIssueKeyFromBranch) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("inject issue key from branch", true, io.VERBOSE)

	branch := a.hookBundle.Repo.BranchName()
	pattern := action.Options().AsString("regex", "([A-Za-z]+\\-[0-9]+)")
	msgFile := a.hookBundle.AppIO.Argument(info.ArgCommitMsgFile, "")

	msg, err := a.hookBundle.Repo.CommitMessage(msgFile)
	if err != nil {
		return err
	}

	r, regErr := regexp.Compile(pattern)
	if regErr != nil {
		return errors.New("invalid regex pattern")
	}

	// can we actually find an issue id?
	match := r.FindStringSubmatch(branch)
	if match == nil {
		return a.noMatchFound(action.Options())
	}

	if len(match) < 2 {
		a.hookBundle.AppIO.Write("issue key must be in regex group (ISSUE-KEY)", true, io.NORMAL)
		return a.noMatchFound(action.Options())
	}

	issueID := match[1]
	// make sure the issue key is not already in our commit message
	if strings.Contains(msg.Subject()+msg.Body(), issueID) {
		return nil
	}
	return a.hookBundle.Repo.PrepareCommitMessage(msgFile, a.createNewCommitMessage(action.Options(), msg, issueID))
}

func (a *InjectIssueKeyFromBranch) createNewCommitMessage(options *configuration.Options, msg *types.CommitMessage, issueID string) *types.CommitMessage {
	// let's figure out where to put the issueID
	target := options.AsString("into", "body")
	mode := options.AsString("mode", "append")
	prefix := options.AsString("prefix", "issue: ")

	// overwrite either subject or body
	newMsgData := map[string]string{"subject": msg.Subject(), "body": msg.Body()}
	newMsgData[target] = a.injectIssueId(issueID, newMsgData[target], mode, prefix)

	var comments []string
	for _, line := range msg.Lines() {
		if strings.HasPrefix(strings.TrimSpace(line), msg.CommentChar()) {
			comments = append(comments, line)
		}
	}

	return types.NewCommitMessage(
		newMsgData["subject"]+"\n\n"+newMsgData["body"]+"\n"+strings.Join(comments, "\n"),
		msg.CommentChar(),
	)
}

func (a *InjectIssueKeyFromBranch) injectIssueId(issueID, target, mode, prefix string) string {
	if mode == "prepend" {
		return prefix + issueID + " " + target
	}
	return target + prefix + issueID

}

func (a *InjectIssueKeyFromBranch) noMatchFound(options *configuration.Options) error {
	if options.AsBool("force", false) {
		return errors.New("no issue key found in branch name")
	}
	return nil
}

func NewInjectIssueKeyFromBranch(appIO io.IO, conf *configuration.Configuration, repo git.Repo) hooks.Action {
	a := InjectIssueKeyFromBranch{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PrepareCommitMsg}),
	}
	return &a
}
