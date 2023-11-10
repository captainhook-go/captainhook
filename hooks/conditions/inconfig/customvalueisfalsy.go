package inconfig

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

type CustomValueIsFalsy struct {
	hookBundle *hooks.HookBundle
}

func (c *CustomValueIsFalsy) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *CustomValueIsFalsy) IsTrue(condition *configuration.Condition) bool {
	value, ok := c.hookBundle.Conf.CustomSettings()["foo"]
	if !ok {
		return true
	}
	if value == "false" {
		return true
	}
	return false
}

func NewCustomValueIsFalsy(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Condition {
	return &CustomValueIsFalsy{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
}
