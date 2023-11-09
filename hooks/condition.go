package hooks

import "github.com/captainhook-go/captainhook/configuration"

type Condition interface {
	IsApplicableFor(hook string) bool
	IsTrue(condition *configuration.Condition) bool
}
