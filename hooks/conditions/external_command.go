package conditions

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/exec/ext"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/hooks/placeholder"
	"github.com/captainhook-go/captainhook/io"
)

type ExternalCommand struct {
	hookBundle *hooks.HookBundle
}

func (c *ExternalCommand) IsApplicableFor(hook string) bool {
	return c.hookBundle.Restriction.IsApplicableFor(hook)
}

func (c *ExternalCommand) IsTrue(condition *configuration.Condition) bool {
	commandToExecute := placeholder.ReplacePlaceholders(
		app.NewContext(c.hookBundle.AppIO, c.hookBundle.Conf, c.hookBundle.Repo),
		condition.Run(),
	)
	// if there were placeholders replaced
	if commandToExecute != condition.Run() {
		c.hookBundle.AppIO.Write("<comment>cmd:</comment>\n"+commandToExecute, true, io.VERBOSE)
	}
	err := ext.ExecuteCommand(c.hookBundle.AppIO, condition.Run())
	if err != nil {
		c.hookBundle.AppIO.Write(err.Error(), true, io.NORMAL)
		return false
	}
	return true
}

func NewExternalCommand(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Condition {
	return &ExternalCommand{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
}
