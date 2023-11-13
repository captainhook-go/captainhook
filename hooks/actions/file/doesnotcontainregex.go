package file

import (
	"errors"
	"fmt"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"regexp"
)

// DoesNotContainRegex blocks commits if a file is containing a string matching the given regex.
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::File.DoesNotContainRegex",
//	  "options: {
//	    "regex": "shouldNotContainThis"
//	  }
//	}
type DoesNotContainRegex struct {
	hookBundle *hooks.HookBundle
}

func (a *DoesNotContainRegex) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *DoesNotContainRegex) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("checking if file does not contain regex", true, io.VERBOSE)

	reg := action.Options().AsString("regex", "")
	if reg == "" {
		return errors.New("the 'regex' option is missing or empty")
	}
	files, err := a.hookBundle.Repo.StagedFiles()
	if err != nil {
		return err
	}
	for _, file := range files {
		content, _ := io.ReadFile(file)
		matched, readErr := regexp.MatchString(reg, string(content))
		if readErr != nil {
			return readErr
		}
		if matched {
			return fmt.Errorf("the file '%s' contains the regex '%s'", file, reg)
		}
	}
	return nil
}

func NewDoesNotContainRegex(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := DoesNotContainRegex{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PreCommit}),
	}
	return &a
}
