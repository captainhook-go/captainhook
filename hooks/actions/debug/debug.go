package debug

import (
	"fmt"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

type Debug struct {
	hookBundle *hooks.HookBundle
	returnFunc func() error
}

func (a *Debug) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *Debug) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("debug action", true, io.VERBOSE)
	a.hookBundle.AppIO.Write("<info>Hook Arguments</info>", true, io.NORMAL)
	for name, value := range a.hookBundle.AppIO.Arguments() {
		a.hookBundle.AppIO.Write(
			fmt.Sprintf("  name: <comment>%s</comment> value: <comment>%s</comment>", name, value),
			true,
			io.NORMAL,
		)
	}
	return a.returnFunc()
}

func NewDebug(bundle *hooks.HookBundle, rFunc func() error) hooks.Action {
	a := Debug{
		hookBundle: bundle,
		returnFunc: rFunc,
	}
	return &a
}
