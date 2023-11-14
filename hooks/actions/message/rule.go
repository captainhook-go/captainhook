package message

import (
	"github.com/captainhook-go/captainhook/git/types"
)

// Rule is used to define rules that commit messages should follow.
// If a commit message does not follow a rule the commit gets aborted.
type Rule interface {
	IsFollowedBy(msg *types.CommitMessage) (bool, string)
}
