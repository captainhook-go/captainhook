package filechanged

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/input"
	"github.com/captainhook-go/captainhook/hooks/util"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"strings"
)

type All struct {
	hookBundle *hooks.HookBundle
}

func (c *All) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *All) IsTrue(condition *configuration.Condition) bool {
	c.hookBundle.AppIO.Write("Condition: FileChanged.All", true, io.VERBOSE)
	ranges := input.DetectRanges(c.hookBundle.AppIO)
	if len(ranges) == 0 {
		return false
	}
	changedFiles, err := c.hookBundle.Repo.ChangedFiles(ranges[0].From().Id(), ranges[0].To().Id())
	if err != nil {
		c.hookBundle.AppIO.Write("Condition FileChanged.ThatIs failed: "+err.Error(), true, io.NORMAL)
		return false
	}
	files := condition.Options().AsString("files", "")
	mustContain := strings.Split(files, ",")
	return util.ContainsAllStrings(changedFiles, mustContain)
}

func NewAll(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Condition {
	return &All{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PrePush}),
	}
}
