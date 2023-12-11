package status

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

// OnBranch prevents an action from being executed you are not working on a specified branch.
//
// Example configuration:
//
//	{
//	  "run": "echo ARRRRRRR",
//	  "conditions": [
//	    {
//	      "run": "CaptainHook::Status.OnBranch",
//	      "options": {
//	        "name": "main"
//	      }
//	    }
//	  ]
//	}
type OnBranch struct {
	hookBundle *hooks.HookBundle
}

func (c *OnBranch) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *OnBranch) IsTrue(condition *configuration.Condition) bool {
	requiredBranch := condition.Options().AsString("name", "")
	if requiredBranch == "" {
		c.hookBundle.AppIO.Write("Condition Status.OnBranch option 'name' is missing", true, io.NORMAL)
		return false
	}
	currentBranch := c.hookBundle.Repo.BranchName()
	if requiredBranch != currentBranch {
		return false
	}
	return true
}

func NewOnBranch(appIO io.IO, conf *configuration.Configuration, repo git.Repo) hooks.Condition {
	return &OnBranch{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
}
