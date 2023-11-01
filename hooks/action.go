package hooks

import "github.com/captainhook-go/captainhook/config"

type Action interface {
	IsApplicableFor(hook string) bool
	Run(action *config.Action) error
}
