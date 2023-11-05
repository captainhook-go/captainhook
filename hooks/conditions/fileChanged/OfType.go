package fileChanged

import (
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

type OfType struct {
	hookBundle *hooks.HookBundle
}

func (c *OfType) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *OfType) IsTrue(condition *config.Condition) bool {
	return false
}

func NewOfType(appIO io.IO, conf *config.Configuration, repo *git.Repository) hooks.Condition {
	return &OfType{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PRE_PUSH}),
	}
}
