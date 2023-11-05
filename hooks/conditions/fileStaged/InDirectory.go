package fileStaged

import (
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

type InDirectory struct {
	hookBundle *hooks.HookBundle
}

func (c *InDirectory) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *InDirectory) IsTrue(condition *config.Condition) bool {
	return false
}

func NewInDirectory(appIO io.IO, conf *config.Configuration, repo *git.Repository) hooks.Condition {
	return &InDirectory{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PRE_COMMIT}),
	}
}
