package input

import (
	"github.com/captainhook-go/captainhook/test"
	"testing"
)

func TestStagedOrChangedFilesWithStaged(t *testing.T) {
	inOut := test.CreateFakeIO()
	inOut.SetArguments(map[string]string{"command": "pre-commit"})
	repo := test.CreateFakeRepo()
	repo.SetFiles([]string{"foo", "bar", "baz"})

	files, _ := StagedOrChangedFiles(inOut, repo)

	if len(files) != 3 {
		t.Errorf("Rego should have returned 3 files")
	}
}

func TestStagedOrChangedFilesWithChanged(t *testing.T) {
	inOut := test.CreateFakeIO()
	inOut.SetArguments(map[string]string{"command": "pre-push"})
	inOut.SetStdIn([]string{"refs/heads/main 12345 refs/heads/main 09876"})
	repo := test.CreateFakeRepo()
	repo.SetFiles([]string{"foo", "bar", "baz"})

	files, _ := StagedOrChangedFiles(inOut, repo)

	if len(files) != 3 {
		t.Errorf("Rego should have returned 3 files")
	}
}

func TestStagedOrChangedFilesWithBrokenInput(t *testing.T) {
	inOut := test.CreateFakeIO()
	inOut.SetArguments(map[string]string{"command": "pre-push"})
	inOut.SetStdIn([]string{""})
	repo := test.CreateFakeRepo()
	repo.SetFiles([]string{"foo", "bar", "baz"})

	files, _ := StagedOrChangedFiles(inOut, repo)

	if len(files) != 0 {
		t.Errorf("Rego should have returned 3 files")
	}
}
