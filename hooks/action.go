package hooks

import "github.com/captainhook-go/captainhook/configuration"

type Action interface {
	IsApplicableFor(hook string) bool
	Run(action *configuration.Action) error
}
