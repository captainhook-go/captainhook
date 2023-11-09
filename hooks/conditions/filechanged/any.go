package filechanged

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

type Any struct {
	hookBundle *hooks.HookBundle
}

func (c *Any) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *Any) IsTrue(condition *configuration.Condition) bool {
	return false
}

func NewAny(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Condition {
	return &Any{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PrePush}),
	}
}
