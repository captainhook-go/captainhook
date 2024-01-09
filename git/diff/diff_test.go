package diff

import (
	"github.com/captainhook-go/captainhook/git/types"
	"testing"
)

func TestCached(t *testing.T) {
	g := types.NewCmd("diff")
	g.AddOptions(Cached)

	if len(g.Options) < 2 {
		t.Errorf("Option not added correctly")
	}
	if g.Options[1] != "--cached" {
		t.Errorf("Wrong option")
	}
}
