package log

import (
	"github.com/captainhook-go/captainhook/git/types"
	"testing"
)

func TestAbbrevRev(t *testing.T) {
	g := types.NewCmd("log")
	g.AddOptions(AbbrevCommit)

	if len(g.Options) < 2 {
		t.Errorf("Option not added correctly")
	}
	if g.Options[1] != "--abbrev-commit" {
		t.Errorf("Wrong option")
	}
}
