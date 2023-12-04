package input

import (
	"fmt"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

// StagedOrChangedFiles will return a list of files
//   - For `pre-commit` hooks it will return the staged files.
//   - For `pre-push` hooks it will return the changed files.
func StagedOrChangedFiles(appIO io.IO, repo *git.Repository) ([]string, error) {
	cmd := appIO.Argument(info.ArgCommand, "")
	if cmd == "pre-commit" {
		return repo.StagedFiles()
	}
	return ChangedFiles(appIO, repo)
}

// ChangedFiles will return a lost of changed files
// Depending on the executed hook it will use different detecting methods to
// determine the `from` ad `to` references.
func ChangedFiles(appIO io.IO, repo *git.Repository) ([]string, error) {
	ranges := DetectRanges(appIO)
	if len(ranges) == 0 {
		return []string{}, fmt.Errorf("could not detect ranges")
	}
	return repo.ChangedFiles(ranges[0].From().Hash(), ranges[0].To().Hash())
}
