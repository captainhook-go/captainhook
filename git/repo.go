package git

import (
	"github.com/captainhook-go/captainhook/git/types"
)

type Repo interface {
	// Path returns the path to the repositories root directory
	Path() string

	// GitDir returns the path to the .git directory of the repository
	GitDir() string

	// HookExists tells you if a hook script for a given hook exists
	HookExists(hook string) bool

	// HooksDir returns the path to the repositories hooks directory
	HooksDir() string

	// CommitMessage creates a CommitMessage type from a message file and returns it
	CommitMessage(path string) (*types.CommitMessage, error)

	// PrepareCommitMessage will write the commit message back to the file
	PrepareCommitMessage(path string, msg *types.CommitMessage) error

	// ConfigValue will return a git config value
	ConfigValue(value string, defaultValue string) string

	// IsMerging tells you if the repository is in a merging state
	IsMerging() bool

	// StagedFiles returns a list of staged files
	StagedFiles() ([]string, error)

	// ChangedFiles returns a list of changed files
	ChangedFiles(from, to string) ([]string, error)

	// BranchName returns the current branch name
	BranchName() string

	// CommitsBetween returns a list of Commit between two hashes
	CommitsBetween(from string, to string) []*types.Commit
}
