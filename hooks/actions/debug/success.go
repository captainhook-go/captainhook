package debug

import (
	"fmt"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

// Success is a debug action to output args and opts
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::Debug.Success",
//	  "option": {
//	    "foo": "bar"
//	  }
//	}
type Success struct {
	hookBundle *hooks.HookBundle
}

func (a *Success) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *Success) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("debug action", true, io.VERBOSE)
	a.hookBundle.AppIO.Write("<info>Hook Arguments</info>", true, io.NORMAL)
	for name, value := range a.hookBundle.AppIO.Arguments() {
		a.hookBundle.AppIO.Write(fmt.Sprintf("Name: %s Value: %s", name, value), true, io.NORMAL)
	}
	return nil
}

func NewSuccess(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := Success{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
	return &a
}
