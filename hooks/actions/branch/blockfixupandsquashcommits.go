package branch

import (
	"errors"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/git/types"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/input"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"slices"
	"strings"
)

// PreventPushOfFixupAndSquashCommits prevents you from pushing fixup! or squash! commits. Either for every
// branch in general or for a given list of branches.
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::Branch.BlockFixupAndSquashCommits",
//	  "options": {
//	    "block-squash-commits": true,
//	    "block-fixup-commits": true,
//	    "branches-to-protect": ["main", "master", "integration"]
//	  }
//	}
type PreventPushOfFixupAndSquashCommits struct {
	hookBundle         *hooks.HookBundle
	blockFixupCommits  bool
	blockSquashCommits bool
	protectedBranches  []string
}

func (a *PreventPushOfFixupAndSquashCommits) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *PreventPushOfFixupAndSquashCommits) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("blocking fixup and squash commits", true, io.VERBOSE)

	refsToPush := input.DetectRanges(a.hookBundle.AppIO)

	if len(refsToPush) == 0 {
		return nil
	}

	a.handleOptions(action.Options())

	for _, aRange := range refsToPush {
		if len(a.protectedBranches) > 0 && !slices.Contains(a.protectedBranches, aRange.From().Branch()) {
			continue
		}
		commits := a.blockedCommits(aRange.From().Id(), aRange.To().Id())

		if len(commits) > 0 {
			return errors.New(a.createFailureMessage(commits, aRange.From().Branch()))
		}
	}
	return nil
}

func (a *PreventPushOfFixupAndSquashCommits) handleOptions(options *configuration.Options) {
	a.blockSquashCommits = options.AsBool("block-squash-commits", true)
	a.blockFixupCommits = options.AsBool("block-fixup-commits", true)
	a.protectedBranches = options.AsSliceOfStrings("branches-to-protect")
}

func (a *PreventPushOfFixupAndSquashCommits) blockedCommits(remoteHash, localHash string) []*types.Commit {
	typesToCheck := a.typesToBlock()
	var blocked []*types.Commit
	for _, commit := range a.hookBundle.Repo.CommitsBetween(remoteHash, localHash) {
		if a.hasToBeBlocked(commit.Subject, typesToCheck) {
			blocked = append(blocked, commit)
		}
	}
	return blocked
}

func (a *PreventPushOfFixupAndSquashCommits) typesToBlock() []string {
	var typesToBlock []string
	if a.blockFixupCommits {
		typesToBlock = append(typesToBlock, "fixup!")
	}
	if a.blockSquashCommits {
		typesToBlock = append(typesToBlock, "squash!")
	}
	return typesToBlock
}

func (a *PreventPushOfFixupAndSquashCommits) hasToBeBlocked(subject string, typesToBlock []string) bool {
	for _, typeToBlock := range typesToBlock {
		if strings.HasPrefix(subject, typeToBlock) {
			return true
		}
	}
	return false
}

func (a *PreventPushOfFixupAndSquashCommits) createFailureMessage(commits []*types.Commit, branch string) string {
	var out []string
	for _, commit := range commits {
		out = append(out, " - "+commit.Hash+" "+commit.Subject)
	}
	return "You are prohibited to push the following commits:\n" +
		" --[ " + branch + " ]-- \n" +
		strings.Join(out, "\n")
}

func NewPreventPushOfFixupAndSquashCommits(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := PreventPushOfFixupAndSquashCommits{
		hookBundle:         hooks.NewHookBundle(appIO, conf, repo, []string{info.PrePush}),
		blockFixupCommits:  true,
		blockSquashCommits: true,
	}
	return &a
}
