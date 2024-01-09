package config

import (
	"github.com/captainhook-go/captainhook/git/types"
	"testing"
)

func TestGet(t *testing.T) {
	g := types.NewCmd("config")
	g.AddOptions(Get("user.name"))

	if len(g.Options) < 3 {
		t.Errorf("Option not added correctly")
	}
	if g.Options[1] != "--get" {
		t.Errorf("Wrong option")
	}
	if g.Options[2] != "user.name" {
		t.Errorf("Wrong option")
	}
}
