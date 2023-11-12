package message

import (
	"github.com/captainhook-go/captainhook/git/types"
)

type Rule interface {
	IsFollowedBy(msg *types.CommitMessage) (bool, string)
}
