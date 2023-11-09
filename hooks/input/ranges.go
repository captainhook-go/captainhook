package input

import (
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/git/types"
	"github.com/captainhook-go/captainhook/io"
	"strings"
)

const (
	LocalRef   = 0
	LocalHash  = 1
	RemoteRef  = 2
	RemoteHash = 3
)

var (
	detectors = map[string]func(appIO io.IO) []*types.Range{
		"pre-push": func(appIO io.IO) []*types.Range {
			var ranges []*types.Range
			for _, line := range appIO.StandardInput() {
				if len(line) == 0 {
					continue
				}
				p := strings.Split(line, " ")

				if git.IsZeroHash(p[RemoteHash]) {
					continue
				}

				from := types.NewRef(p[RemoteRef], p[RemoteHash], git.ExtractBranchFromRefPath(p[RemoteRef]))
				to := types.NewRef(p[LocalRef], p[LocalHash], git.ExtractBranchFromRefPath(p[LocalRef]))
				ranges[] = types.NewRange(from, to)
			}
			return []*types.Range{}
		},
		"post-rewrite": func(appIO io.IO) []*types.Range {
			var ranges []*types.Range
			for _, input := range appIO.StandardInput() {
				if len(input) > 0 {
					parts := strings.Split(strings.TrimSpace(input), " ")
					id := "HEAD@{1}"
					if len(parts) > 1 {
						id = parts[1] + "^"
					}
					from := types.NewRef(id, "", "")
					to := types.NewRef("HEAD", "", "")
					ranges = append(ranges, types.NewRange(from, to))
				}
			}
			return ranges
		},
		"fallback": func(appIO io.IO) []*types.Range {
			var ranges []*types.Range
			r := types.NewRange(
				types.NewRef(appIO.Argument("previousHead", "HEAD@{1}"), "", ""),
				types.NewRef("HEAD", "", ""),
			)
			ranges = append(ranges, r)
			return ranges
		},
	}
)

func DetectRange(appIO io.IO) []*types.Range {
	command := appIO.Argument("command", "fallback")

	detector, ok := detectors[command]
	if !ok {
		detector = detectors["fallback"]
	}
	return detector(appIO)
}
