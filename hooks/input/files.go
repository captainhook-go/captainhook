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
func StagedOrChangedFiles(appIO io.IO, repo git.Repo) ([]string, error) {
	cmd := appIO.Argument(info.ArgCommand, "")
	if cmd == "pre-commit" {
		return repo.StagedFiles()
	}
	return ChangedFiles(appIO, repo)
}

// ChangedFiles will return a list of changed files
// It uses a Detector that depending on the executed hook will use different methods to
// detect the `from` ad `to` references.
func ChangedFiles(appIO io.IO, repo git.Repo) ([]string, error) {
	ranges := DetectRanges(appIO)
	if len(ranges) == 0 {
		return []string{}, fmt.Errorf("could not detect ranges")
	}
	return repo.ChangedFiles(ranges[0].From().Hash(), ranges[0].To().Hash())
}
