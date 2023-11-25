package file

import (
	"fmt"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"regexp"
)

// BlockSecrets blocks commits if a file is containing a string matching any if the given regexes.
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::File.BlockSecrets",
//	  "options: {
//	    "use-defaults": true,
//	    "block": ["pattern1", "pattern2"],
//	    "allow": ["patternA"]
//	  }
//	}
type BlockSecrets struct {
	hookBundle       *hooks.HookBundle
	blockDefault     bool
	blockedByDefault []string
	blockedByUser    []string
	allowedByUser    []string
}

func (a *BlockSecrets) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *BlockSecrets) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("checking if file does not contain regex", true, io.VERBOSE)

	a.blockDefault = action.Options().AsBool("use-defaults", true)
	a.blockedByUser = action.Options().AsSliceOfStrings("blocked")
	a.allowedByUser = action.Options().AsSliceOfStrings("allowed")

	files, err := a.hookBundle.Repo.StagedFiles()
	if err != nil {
		return err
	}
	for _, file := range files {
		content, _ := io.ReadFile(file)
		containsBlocked, err := a.containsBlockedString(string(content))
		if containsBlocked {
			return fmt.Errorf("file '%s' %s", file, err.Error())
		}
	}
	return nil
}

func (a *BlockSecrets) containsBlockedString(cnt string) (bool, error) {
	regs := a.blockedByUser
	if a.blockDefault {
		regs = append(a.blockedByDefault, regs...)
	}
	for _, reg := range regs {
		r, _ := regexp.Compile(reg)
		matches := r.FindStringSubmatch(cnt)
		if len(matches) > 0 && !a.isAllowed(matches[0]) {
			return true, fmt.Errorf("contains %s", matches[0])
		}
	}
	return false, nil
}

func (a *BlockSecrets) isAllowed(blocked string) bool {
	for _, reg := range a.allowedByUser {
		matched, readErr := regexp.MatchString(reg, blocked)
		if readErr != nil {
			return false
		}
		if matched {
			return true
		}
	}
	return false
}

func NewBlockSecrets(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	aws := "(AWS|aws|Aws)?_?"
	quote := "(\"|')"
	optQuote := quote + "?"
	connect := "\\s*(:|=>|=|:=)\\s*"

	a := BlockSecrets{
		hookBundle:   hooks.NewHookBundle(appIO, conf, repo, []string{info.PreCommit}),
		blockDefault: true,
		blockedByDefault: []string{
			"(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}",
			optQuote + aws + "(SECRET|secret|Secret)?_?(ACCESS|access|Access)?_?(KEY|key|Key)" + optQuote + connect + optQuote + "[A-Za-z0-9/\\+=]{40}$" + optQuote,
			optQuote + aws + "(ACCOUNT|account|Account)_?(ID|id|Id)?" + optQuote + connect + optQuote + "[0-9]{4}\\-?[0-9]{4}\\-?[0-9]{4}" + optQuote,
			"(?i)password" + optQuote + connect + optQuote + "[a-z\\-_#/\\+0-9]{16,}" + optQuote,
		},
	}
	return &a
}
