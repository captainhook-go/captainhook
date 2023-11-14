package git

import (
	"github.com/captainhook-go/captainhook/git/revparse"
	"regexp"
	"strings"
)

// IsZeroHash indicates if commit hash is a zero hash 0000000000000000000000000000000000000000
func IsZeroHash(hash string) bool {
	matched, _ := regexp.MatchString("^0+$", hash)
	return matched
}

func ExtractBranchFromRefPath(head string) string {
	parts := strings.Split(head, "/")
	return parts[len(parts)-1]
}

func DetectGitDir() (string, error) {
	out, err := RevParse(
		revparse.ShowTopLevel,
	)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out), nil
}
