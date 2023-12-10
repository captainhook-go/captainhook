package file

import (
	"fmt"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/input"
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
//	    "presets": ["Aws", "GitHub", "Stripe", "Google"],
//	    "blocked": ["pattern1", "pattern2"],
//	    "allowed": ["patternA"]
//	  }
//	}
type BlockSecrets struct {
	hookBundle *hooks.HookBundle
	presets    []string
	blocked    []string
	allowed    []string
}

func (a *BlockSecrets) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *BlockSecrets) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("checking if file does not contain regex", true, io.VERBOSE)

	a.presets = action.Options().AsSliceOfStrings("presets")
	a.blocked = action.Options().AsSliceOfStrings("blocked")
	a.allowed = action.Options().AsSliceOfStrings("allowed")

	files, err := input.StagedOrChangedFiles(a.hookBundle.AppIO, a.hookBundle.Repo)
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
	regs, err := a.getRegexPatterns()
	if err != nil {
		return false, err
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
	for _, reg := range a.allowed {
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

func (a *BlockSecrets) getRegexPatterns() ([]string, error) {
	regs := a.blocked
	for _, name := range a.presets {
		preset, err := GetPreset(name)
		if err != nil {
			return nil, err
		}
		regs = append(regs, preset...)
	}
	return regs, nil
}

func NewBlockSecrets(appIO io.IO, conf *configuration.Configuration, repo git.Repo) hooks.Action {
	return &BlockSecrets{hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PreCommit, info.PrePush})}
}
