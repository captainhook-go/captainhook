package fileStaged

import (
	"github.com/captainhook-go/captainhook/config"
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

func (c *ThatIs) IsTrue(condition *config.Condition) bool {
	ofType := condition.Options().StringValueOf("of-type", "")
	c.hookBundle.AppIO.Write("Condition: "+ofType, true, io.VERBOSE)
	c.hookBundle.AppIO.Write("Condition: FileStages.ThatIs", true, io.VERBOSE)
	return false
}

func NewThatIs(appIO io.IO, conf *config.Configuration, repo *git.Repository) hooks.Condition {
	return &ThatIs{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PRE_COMMIT}),
	}
}
