package filechanged

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/input"
	"github.com/captainhook-go/captainhook/hooks/util"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

// ThatIs makes sure an action is only executed if a file with the given configuration is changed.
// Only applicable for 'pre-push' hooks.
//
// Example configuration:
//
//	{
//	  "run": "echo ARRRRRR",
//	  "conditions": [
//	    {
//	      "run": "CaptainHook::FilesChanged.ThatIs",
//	      "options": {
//	        "of-type": "go",
//	        "in-directory": "app"
//	      }
//	    }
//	  ]
//	}
type ThatIs struct {
	hookBundle *hooks.HookBundle
}

func (c *ThatIs) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *ThatIs) IsTrue(condition *configuration.Condition) bool {
	c.hookBundle.AppIO.Write("Condition: FileChanged.ThatIs", true, io.VERBOSE)
	ranges := input.DetectRanges(c.hookBundle.AppIO)
	if len(ranges) == 0 {
		return false
	}
	changedFiles, err := c.hookBundle.Repo.ChangedFiles(ranges[0].From().Id(), ranges[0].To().Id())
	if err != nil {
		c.hookBundle.AppIO.Write("Condition FileChanged.ThatIs failed: "+err.Error(), true, io.NORMAL)
		return false
	}
	files := changedFiles
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

func NewThatIs(appIO io.IO, conf *configuration.Configuration, repo git.Repo) hooks.Condition {
	return &ThatIs{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PrePush}),
	}
}
