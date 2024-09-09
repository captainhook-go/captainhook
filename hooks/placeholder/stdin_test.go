package placeholder

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/io"
	"testing"
)

func TestStdin(t *testing.T) {
	argMap := map[string]string{}
	optMap := map[string]string{"input": "foo bar"}
	expected := "'foo bar'"

	config := configuration.NewConfiguration("foo", false)
	repo, _ := git.NewRepository(".git")
	ctx := app.NewContext(
		io.NewDefaultIO(io.NORMAL, optMap, argMap),
		config,
		repo,
	)
	opts := map[string]string{}
	placeholder := &StdIn{context: ctx}
	result := placeholder.Replacement(opts)
	if result != expected {
		t.Errorf("Replacement didn't work, got: %s, want: %s.", result, expected)
	}
}

func TestStdinUnescaped(t *testing.T) {
	argMap := map[string]string{}
	optMap := map[string]string{"input": "foo bar"}
	expected := "foo bar"

	config := configuration.NewConfiguration("foo", false)
	repo, _ := git.NewRepository(".git")
	ctx := app.NewContext(
		io.NewDefaultIO(io.NORMAL, optMap, argMap),
		config,
		repo,
	)
	opts := map[string]string{"escaped": "false"}
	placeholder := &StdIn{context: ctx}
	result := placeholder.Replacement(opts)
	if result != expected {
		t.Errorf("Replacement didn't work, got: %s, want: %s.", result, expected)
	}
}
