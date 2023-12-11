package test

import "github.com/captainhook-go/captainhook/git/types"

type RepoMock struct {
	path   string
	branch string
	files  []string
	log    []*types.Commit
}

func (r *RepoMock) setBranch(name string) *RepoMock {
	r.branch = name
	return r
}

func (r *RepoMock) setFiles(files []string) *RepoMock {
	r.files = files
	return r
}

func (r *RepoMock) Path() string {
	return r.path
}

func (r *RepoMock) GitDir() string {
	return r.Path() + "/.git"
}

func (r *RepoMock) HookExists(hook string) bool {
	return true
}

func (r *RepoMock) HooksDir() string {
	return r.GitDir() + "/hooks"
}

func (r *RepoMock) CommitMessage(path string) (*types.CommitMessage, error) {
	return types.NewCommitMessage("", "#"), nil
}

func (r *RepoMock) PrepareCommitMessage(path string, msg *types.CommitMessage) error {
	return nil
}

func (r *RepoMock) ConfigValue(value string, defaultValue string) string {
	return ""
}

func (r *RepoMock) IsMerging() bool {
	return false
}

func (r *RepoMock) StagedFiles() ([]string, error) {
	return r.files, nil
}

func (r *RepoMock) ChangedFiles(from, to string) ([]string, error) {
	return r.files, nil
}

func (r *RepoMock) BranchName() string {
	return r.branch
}

func (r *RepoMock) CommitsBetween(from string, to string) []*types.Commit {
	return []*types.Commit{}
}
