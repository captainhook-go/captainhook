package hooks

import "github.com/captainhook-go/captainhook/configuration"

type Action interface {
	// IsApplicableFor answers if an `Action` can be used for a given hook
	// For example actions that verify the commit message are only applicable for
	// the `commit-message` hook.
	IsApplicableFor(hook string) bool

	// Run executes the action
	// This will either execute some functionality provided from CaptainHook
	// or some external executable.
	Run(action *configuration.Action) error
}
