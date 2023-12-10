package inconfig

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

// CustomValueIsFalsy prevents an action from being executed if a config 'custom' value is not falsy
//
// Example configuration:
//
//	{
//	  "run": "echo ARRRRRRR",
//	  "conditions": [
//	    {
//	      "run": "CaptainHook::InConfig.CustomValueIsFalsy",
//	      "options": {
//	        "value": "my-custom-value-that-must-be-falsy"
//	      }
//	    }
//	  ]
//	}
type CustomValueIsFalsy struct {
	hookBundle *hooks.HookBundle
}

func (c *CustomValueIsFalsy) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *CustomValueIsFalsy) IsTrue(condition *configuration.Condition) bool {
	valueToCheck := condition.Options().AsString("value", "")
	if valueToCheck == "" {
		c.hookBundle.AppIO.Write("Condition Config.CustomValueIsFalsy option 'value' is missing", true, io.NORMAL)
		return true
	}
	value, ok := c.hookBundle.Conf.CustomSettings()[valueToCheck]
	if !ok {
		c.hookBundle.AppIO.Write("Condition Config.CustomValueIsFalsy custom value not set", true, io.DEBUG)
		return true
	}
	return !io.AnswerToBool(value)
}

func NewCustomValueIsFalsy(appIO io.IO, conf *configuration.Configuration, repo git.Repo) hooks.Condition {
	return &CustomValueIsFalsy{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
}
