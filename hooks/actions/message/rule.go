package message

import (
	"github.com/captainhook-go/captainhook/git/types"
)

type Rule interface {
	AppliesTo(msg types.CommitMessage) (bool, string)
}
