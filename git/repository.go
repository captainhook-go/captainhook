package git

import (
	"fmt"
	"os"
)

type Repository struct {
	path      string
	operators map[string]string
}

func NewRepository(repoPath string) (*Repository, error) {
	if !isPathARepository(repoPath) {
		err := fmt.Errorf("repository not found in: %s", repoPath)
		return nil, err
	}
	r := Repository{path: repoPath}
	return &r, nil
}

func (r *Repository) Path() string {
	return r.path
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
