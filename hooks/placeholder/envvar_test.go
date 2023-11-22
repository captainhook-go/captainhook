package placeholder

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/io"
	"os"
	"testing"
)

func TestEnvVarFound(t *testing.T) {
	expected := "BAR"

	os.Setenv("FOO", "BAR")

	config := configuration.NewConfiguration("foo", false)
	repo, _ := git.NewRepository(".git")
	ctx := app.NewContext(
		io.NewDefaultIO(io.NORMAL, map[string]string{}),
		config,
		repo,
	)
	opts := map[string]string{"value-of": "FOO", "default": "xxx"}
	e := &EnvVar{context: ctx}
	result := e.Replacement(opts)

	os.Remove("FOO")

	if result != expected {
		t.Errorf("Replacement didn't work, got: %s, want: %s.", result, expected)
	}
}

func TestEnvVarDefault(t *testing.T) {
	expected := "xxx"

	config := configuration.NewConfiguration("foo", false)
	repo, _ := git.NewRepository(".git")
	ctx := app.NewContext(
		io.NewDefaultIO(io.NORMAL, map[string]string{}),
		config,
		repo,
	)
	opts := map[string]string{"value-of": "X_X_X", "default": "xxx"}
	e := &EnvVar{context: ctx}
	result := e.Replacement(opts)

	if result != expected {
		t.Errorf("Replacement didn't work, got: %s, want: %s.", result, expected)
	}
}
