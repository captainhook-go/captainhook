package actions

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

func (a *ExternalCommand) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *ExternalCommand) Run(action *configuration.Action) error {
	commandToExecute := placeholder.ReplacePlaceholders(
		app.NewContext(a.hookBundle.AppIO, a.hookBundle.Conf, a.hookBundle.Repo),
		action.Run(),
	)
	// if there were placeholders replaced
	if commandToExecute != action.Run() {
		a.hookBundle.AppIO.Write("<comment>cmd:</comment>\n"+commandToExecute, true, io.VERBOSE)
	}
	return ext.ExecuteCommand(a.hookBundle.AppIO, commandToExecute)
}

func NewExternalCommand(appIO io.IO, conf *configuration.Configuration, repo git.Repo) hooks.Action {
	a := ExternalCommand{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
	return &a
}
