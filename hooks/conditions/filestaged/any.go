package filestaged

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/util"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"strings"
)

type Any struct {
	hookBundle *hooks.HookBundle
}

func (c *Any) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *Any) IsTrue(condition *configuration.Condition) bool {
	stagedFiles, err := c.hookBundle.Repo.StagedFiles()
	if err != nil {
		c.hookBundle.AppIO.Write("Condition All failed: "+err.Error(), true, io.NORMAL)
		return false
	}
	files := condition.Options().StringValueOf("files", "")
	mustContain := strings.Split(files, ",")
	return util.ContainsAnyString(stagedFiles, mustContain)
}

func NewAny(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Condition {
	return &Any{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PreCommit}),
	}
}
