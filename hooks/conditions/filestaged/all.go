package filestaged

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/util"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

// All makes sure an action is only executed if all the configured files are staged.
// Only applicable for 'pre-commit' hooks.
//
// Example configuration:
//
//	{
//	  "run": "echo ARRRRRR",
//	  "conditions": [
//	    {
//	      "run": "CaptainHook::FilesStaged.All",
//	      "options": {
//	        "files": ["foo.txt", "bar.txt"]
//	      }
//	    }
//	  ]
//	}
type All struct {
	hookBundle *hooks.HookBundle
}

func (c *All) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *All) IsTrue(condition *configuration.Condition) bool {
	stagedFiles, err := c.hookBundle.Repo.StagedFiles()
	if err != nil {
		c.hookBundle.AppIO.Write("Condition All failed: "+err.Error(), true, io.NORMAL)
		return false
	}
	mustContainAll := condition.Options().AsSliceOfStrings("files")
	return util.ContainsAllStrings(stagedFiles, mustContainAll)
}

func NewAll(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Condition {
	return &All{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PreCommit}),
	}
}
