package revparse

import (
	"github.com/captainhook-go/captainhook/git/types"
	"testing"
)

func TestAbbrevRev(t *testing.T) {
	g := types.NewCmd("revparse")
	g.AddOptions(AbbrevRef)

	if len(g.Options) < 2 {
		t.Errorf("Option not added correctly")
	}
	if g.Options[1] != "--abbrev-ref" {
		t.Errorf("Wrong option")
	}
}
