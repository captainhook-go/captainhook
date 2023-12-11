package test

import (
	"errors"
	"github.com/captainhook-go/captainhook/git/types"
)

type RepoMock struct {
	triggerFileError bool
	path             string
	branch           string
	fileList         []string
	log              []*types.Commit
}

func (r *RepoMock) SetBranch(name string) *RepoMock {
	r.branch = name
	return r
}

func (r *RepoMock) SetFiles(files []string) *RepoMock {
	r.fileList = files
	return r
}

func (r *RepoMock) SetFilesError(triggerError bool) *RepoMock {
	r.triggerFileError = triggerError
	return r
}

func (r *RepoMock) files() ([]string, error) {
	if r.triggerFileError {
		return []string{}, errors.New("files error")
	}
	return r.fileList, nil
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
	return r.files()
}

func (r *RepoMock) ChangedFiles(from, to string) ([]string, error) {
	return r.files()
}

func (r *RepoMock) BranchName() string {
	return r.branch
}

func (r *RepoMock) CommitsBetween(from string, to string) []*types.Commit {
	return []*types.Commit{}
}
