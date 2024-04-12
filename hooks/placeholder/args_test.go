package placeholder

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/io"
	"testing"
)

func TestArgs(t *testing.T) {
	argMap := map[string]string{"message-file": "bar.txt"}
	optMap := map[string]string{"input": ""}
	expected := "bar.txt"

	config := configuration.NewConfiguration("foo", false)
	repo, _ := git.NewRepository(".git")
	ctx := app.NewContext(
		io.NewDefaultIO(io.NORMAL, optMap, argMap),
		config,
		repo,
	)
	opts := map[string]string{"value-of": "message-file", "default": "xxx"}
	args := &Args{context: ctx}
	result := args.Replacement(opts)
	if result != expected {
		t.Errorf("Replacement didn't work, got: %s, want: %s.", result, expected)
	}
}
