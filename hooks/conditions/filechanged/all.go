package filechanged

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
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
	changedFiles, err := c.hookBundle.Repo.ChangedFiles("from", "to")
	if err != nil {
		c.hookBundle.AppIO.Write("Condition All failed: "+err.Error(), true, io.NORMAL)
		return false
	}
	files := condition.Options().StringValueOf("files", "")
	mustContain := strings.Split(files, ",")
	return util.ContainsAllStrings(changedFiles, mustContain)
}

func NewAll(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Condition {
	return &All{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PrePush}),
	}
}
