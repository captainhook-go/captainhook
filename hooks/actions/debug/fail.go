package debug

import (
	"errors"
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
type Fail struct{}

func NewFail(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	return newDebug(
		hooks.NewHookBundle(appIO, conf, repo, []string{}),
		func() error {
			return errors.New("fail debug action")
		},
	)
}
