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
	return false
}

func NewCustomValueIsTruthy(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Condition {
	return &CustomValueIsTruthy{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
}
