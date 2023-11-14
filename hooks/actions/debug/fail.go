package debug

import (
	"errors"
	"fmt"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

// Fail is a debug action to output the original hook args.
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::Debug.Fail"
//	}
type Fail struct {
	hookBundle *hooks.HookBundle
}

func (a *Fail) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *Fail) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("debug action", true, io.VERBOSE)
	a.hookBundle.AppIO.Write("<info>hook arguments</info>", true, io.NORMAL)
	for name, value := range a.hookBundle.AppIO.Arguments() {
		a.hookBundle.AppIO.Write(fmt.Sprintf("name: '%s' value: '%s'", name, value), true, io.NORMAL)
	}
	return errors.New("fail debug action")
}

func NewFail(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := Fail{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
	return &a
}
