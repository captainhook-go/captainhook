package filestaged

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/util"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

type ThatIs struct {
	hookBundle *hooks.HookBundle
}

func (c *ThatIs) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *ThatIs) IsTrue(condition *configuration.Condition) bool {
	c.hookBundle.AppIO.Write("Condition: FileStaged.ThatIs", true, io.VERBOSE)
	stagedFiles, err := c.hookBundle.Repo.StagedFiles()
	if err != nil {
		c.hookBundle.AppIO.Write("Condition ThatIs failed: "+err.Error(), true, io.NORMAL)
		return false
	}
	files := stagedFiles
	files = c.filterByType(condition, files)
	files = c.filterByDirectory(condition, files)

	if len(files) > 0 {
		return true
	}
	return false
}

func (c *ThatIs) filterByType(condition *configuration.Condition, files []string) []string {
	ofType := condition.Options().AsString("of-type", "")
	if ofType == "" {
		return files
	}

	c.hookBundle.AppIO.Write("  filter by type: '"+ofType+"'", true, io.DEBUG)
	ofType = "." + ofType
	return util.FilterByType(files, ofType)
}

func (c *ThatIs) filterByDirectory(condition *configuration.Condition, files []string) []string {
	inDir := condition.Options().AsString("in-directory", "")
	if inDir == "" {
		return files
	}
	return util.FilterByDirectory(files, inDir)
}

func NewThatIs(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Condition {
	return &ThatIs{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PreCommit}),
	}
}
