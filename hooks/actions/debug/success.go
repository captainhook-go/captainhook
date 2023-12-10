package debug

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

// Success is a debug action to output args and opts not failing the hook execution
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::Debug.Success",
//	  "option": {
//	    "foo": "bar"
//	  }
//	}
type Success struct{}

func NewSuccess(appIO io.IO, conf *configuration.Configuration, repo git.Repo) hooks.Action {
	return newDebug(
		hooks.NewHookBundle(appIO, conf, repo, []string{}),
		func() error {
			return nil
		},
	)
}
