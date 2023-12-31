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

// Any makes sure an action is only executed if any of the configured files is changed.
// Only applicable for 'pre-push' hooks.
//
// Example configuration:
//
//	{
//	  "run": "echo ARRRRRR",
//	  "conditions": [
//	    {
//	      "run": "CaptainHook::FilesChanged.Any",
//	      "options": {
//	        "files": ["foo.txt", "bar.txt"]
//	      }
//	    }
//	  ]
//	}
type Any struct {
	hookBundle *hooks.HookBundle
}

func (c *Any) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *Any) IsTrue(condition *configuration.Condition) bool {
	c.hookBundle.AppIO.Write("Condition: FileChanged.Any", true, io.VERBOSE)
	ranges := input.DetectRanges(c.hookBundle.AppIO)
	if len(ranges) == 0 {
		return false
	}
	changedFiles, err := c.hookBundle.Repo.ChangedFiles(ranges[0].From().Id(), ranges[0].To().Id())
	if err != nil {
		c.hookBundle.AppIO.Write("Condition FileChanged.ThatIs failed: "+err.Error(), true, io.NORMAL)
		return false
	}
	needles := condition.Options().AsSliceOfStrings("files")
	return util.ContainsAnyString(changedFiles, needles)
}

func NewAny(appIO io.IO, conf *configuration.Configuration, repo git.Repo) hooks.Condition {
	return &Any{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PrePush}),
	}
}
