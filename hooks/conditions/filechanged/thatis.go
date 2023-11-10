package filechanged

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

type ThatIs struct {
	hookBundle *hooks.HookBundle
}

func (c *ThatIs) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *ThatIs) IsTrue(condition *configuration.Condition) bool {
	return false
}

func NewThatIs(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Condition {
	return &ThatIs{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PrePush}),
	}
}
