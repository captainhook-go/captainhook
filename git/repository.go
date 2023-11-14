package git

import (
	"fmt"
	"github.com/captainhook-go/captainhook/git/config"
	"github.com/captainhook-go/captainhook/git/diff"
	"github.com/captainhook-go/captainhook/git/log"
	"github.com/captainhook-go/captainhook/git/revparse"
	"github.com/captainhook-go/captainhook/git/types"
	"github.com/captainhook-go/captainhook/io"
	"os"
	"path"
	"regexp"
)

type Repository struct {
	root   string
	gitDir string
}

func NewRepository(gitDir string) (*Repository, error) {
	repoPath := path.Dir(gitDir)
	if !isPathARepository(gitDir) {
		err := fmt.Errorf("repository not found in: %s", gitDir)
		return nil, err
	}
	dotGitDir := gitDir
	if isGitSubmodule(dotGitDir) {
		// For submodules hooks are stored in the parents .git/modules directory
		dotGitContents, err := os.ReadFile(dotGitDir)
		if err != nil {
			return nil, err
		}
		r, _ := regexp.Compile("^gitdir:\\s*(.+)$")
		match := r.FindAllStringSubmatch(string(dotGitContents), -1)
		if len(match) > 0 {
			dotGitDir = fmt.Sprintf("%s/%s", repoPath, match[1])
		}
	}
	r := Repository{root: repoPath, gitDir: dotGitDir}
	return &r, nil
}

func (r *Repository) Path() string {
	return r.root
}

func (r *Repository) GitDir() string {
	return r.gitDir
}

func (r *Repository) HookExists(hook string) bool {
	return io.FileExists(r.HooksDir() + "/" + hook)
}

func (r *Repository) HooksDir() string {
	return r.root + "/hooks"
}

func (r *Repository) CommitMessage(path string) (*types.CommitMessage, error) {
	commentChar := r.ConfigValue("core.commentchar", "#")
	return types.NewCommitMessageFromFile(path, commentChar)
}

func (r *Repository) ConfigValue(value string, defaultValue string) string {
	// git config --get VALUE
	out, err := Config(config.Get(value))
	if err != nil {
		return defaultValue
	}
	return out
}

func (r *Repository) IsMerging() bool {
	for _, file := range []string{"MERGE_MSG", "MERGE_HEAD", "MERGE_MODE"} {
		if _, err := os.Stat(r.gitDir + "/" + file); err == nil {
			return true
		}
	}
	return false
}

func (r *Repository) StagedFiles() ([]string, error) {
	// git diff-index --diff-algorithm=myers --no-ext-diff --cached --name-status HEAD
	out, err := DiffIndex(
		diff.Algorithm("myers"),
		diff.NoExtDiff,
		diff.Cached,
		log.NameOnly,
		diff.To("HEAD"),
	)
	if err != nil {
		return nil, err
	}
	return io.SplitLines(out), nil
}

func (r *Repository) ChangedFiles(from, to string) ([]string, error) {
	// git diff-tree --diff-algorithm=myers --no-ext-diff --no-commit-id --name-only -r --diff-filter=ADM FROM TO
	out, err := DiffTree(
		diff.NoExtDiff,
		log.NameOnly,
		diff.Recursive,
		diff.Filter("ACMR"),
		diff.FromTo(from, to),
	)
	if err != nil {
		return []string{}, err
	}
	return io.SplitLines(out), nil
}

func (r *Repository) BranchName() string {
	// rev-parse --abbrev-ref HEAD
	out, err := RevParse(revparse.AbbrevRef)
	if err != nil {
		return ""
	}
	return out
}

func (r *Repository) CommitsBetween(from string, to string) []*types.Commit {
	// git log --abbrev-commit --no-merges FROM TO
	out, err := Log(
		log.Format(log.XmlFormat),
		log.AbbrevCommit,
		log.NoMerges,
		log.FromTo(from, to),
	)
	if err != nil {
		return []*types.Commit{}
	}

	// TODO: handle error
	commits, _ := log.ParseXML("<log>" + out + "</log>")
	return commits
}

func isPathARepository(repoPath string) bool {
	// check at least the two most obvious files
	pathSlice := []string{
		repoPath + "/config",
		repoPath + "/HEAD",
	}
	for _, path := range pathSlice {
		_, err := os.Stat(path)
		if err != nil {
			return false
		}
	}
	return true
}

func isGitSubmodule(dotGitDir string) bool {
	file, err := os.Open(dotGitDir)
	if err != nil {
		return false
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}
