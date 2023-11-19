package placeholder

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/io"
	"testing"
)

func TestFileListDefault(t *testing.T) {
	files := []string{"x/foo.txt", "x/bar.md", "y/fiz.txt", "y/baz.md"}
	expected := "x/foo.txt x/bar.md y/fiz.txt y/baz.md"

	config, _ := configuration.NewConfiguration("foo", false, configuration.NewDefaultAppSettings())
	repo, _ := git.NewRepository(".git")
	ctx := app.NewContext(
		io.NewDefaultIO(io.NORMAL, map[string]string{}),
		config,
		repo,
	)
	opts := map[string]string{}
	p := &FileList{name: "StagedFiles", context: ctx, files: files}
	result := p.Replacement(opts)

	if result != expected {
		t.Errorf("Replacement didn't work, got: %s, want: %s.", result, expected)
	}
}

func TestFileListInDirectory(t *testing.T) {
	files := []string{"x/foo.txt", "x/bar.md", "y/fiz.txt", "y/baz.md"}
	expected := "x/foo.txt x/bar.md"

	config, _ := configuration.NewConfiguration("foo", false, configuration.NewDefaultAppSettings())
	repo, _ := git.NewRepository(".git")
	ctx := app.NewContext(
		io.NewDefaultIO(io.NORMAL, map[string]string{}),
		config,
		repo,
	)
	opts := map[string]string{"in-directory": "x", "separated-by": " "}
	p := &FileList{name: "StagedFiles", context: ctx, files: files}
	result := p.Replacement(opts)

	if result != expected {
		t.Errorf("Replacement didn't work, got: %s, want: %s.", result, expected)
	}
}

func TestFileListOfType(t *testing.T) {
	files := []string{"x/foo.txt", "x/bar.md", "y/fiz.txt", "y/baz.md"}
	expected := "x/foo.txt y/fiz.txt"

	config, _ := configuration.NewConfiguration("foo", false, configuration.NewDefaultAppSettings())
	repo, _ := git.NewRepository(".git")
	ctx := app.NewContext(
		io.NewDefaultIO(io.NORMAL, map[string]string{}),
		config,
		repo,
	)
	opts := map[string]string{"of-type": "txt"}
	p := &FileList{name: "StagedFiles", context: ctx, files: files}
	result := p.Replacement(opts)

	if result != expected {
		t.Errorf("Replacement didn't work, got: %s, want: %s.", result, expected)
	}
}

func TestFileListSeparatedBy(t *testing.T) {
	files := []string{"x/foo.txt", "x/bar.md", "y/fiz.txt", "y/baz.md"}
	expected := "x/foo.txt"

	config, _ := configuration.NewConfiguration("foo", false, configuration.NewDefaultAppSettings())
	repo, _ := git.NewRepository(".git")
	ctx := app.NewContext(
		io.NewDefaultIO(io.NORMAL, map[string]string{}),
		config,
		repo,
	)
	opts := map[string]string{"of-type": "txt", "in-directory": "x", "separated-by": ","}
	p := &FileList{name: "StagedFiles", context: ctx, files: files}
	result := p.Replacement(opts)

	if result != expected {
		t.Errorf("Replacement didn't work, got: %s, want: %s.", result, expected)
	}
}
