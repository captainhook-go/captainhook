package inConfig

import (
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

type CustomValueIsFalsy struct {
	hookBundle *hooks.HookBundle
}

func (c *CustomValueIsFalsy) IsApplicableFor(hook string) bool {
	return true
}

func (c *CustomValueIsFalsy) IsTrue(condition *config.Condition) bool {
	return false
}

func NewCustomValueIsFalsy(appIO io.IO, conf *config.Configuration, repo *git.Repository) hooks.Condition {
	return &CustomValueIsFalsy{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
}
