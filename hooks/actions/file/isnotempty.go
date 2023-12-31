package file

import (
	"fmt"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

// IsNotEmpty is preventing you from committing empty files to the repository
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::File:IsNotEmpty",
//	  "options": {
//	    "files": ["README.md", "LICENSE.md"]
//	  }
//	}
type IsNotEmpty struct {
	hookBundle *hooks.HookBundle
}

func (a *IsNotEmpty) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *IsNotEmpty) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("checking if file is empty", true, io.VERBOSE)

	for _, file := range action.Options().AsSliceOfStrings("files") {
		content, err := io.ReadFile(file)
		if err != nil {
			return fmt.Errorf("file not found: %s", file)
		}
		if len(content) < 1 {
			return fmt.Errorf("file '%s' can't be empty", file)
		}
	}
	return nil
}

func NewIsNotEmpty(appIO io.IO, conf *configuration.Configuration, repo git.Repo) hooks.Action {
	a := IsNotEmpty{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
	return &a
}
