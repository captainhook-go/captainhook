package inconfig

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

type CustomValueIsTruthy struct {
	hookBundle *hooks.HookBundle
}

func (c *CustomValueIsTruthy) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *CustomValueIsTruthy) IsTrue(condition *configuration.Condition) bool {

	valueToCheck := condition.Options().AsString("value", "")
	if valueToCheck == "" {
		c.hookBundle.AppIO.Write("Condition Config.CustomValueIsTruthy option 'value' is missing", true, io.NORMAL)
		return true
	}
	value, ok := c.hookBundle.Conf.CustomSettings()[valueToCheck]
	if !ok {
		c.hookBundle.AppIO.Write("Condition Config.CustomValueIsTruthy custom value not set", true, io.DEBUG)
		return false
	}
	return io.AnswerToBool(value)
}

func NewCustomValueIsTruthy(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Condition {
	return &CustomValueIsTruthy{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
}
