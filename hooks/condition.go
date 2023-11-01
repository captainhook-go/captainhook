package hooks

import "github.com/captainhook-go/captainhook/config"

type Condition interface {
	IsApplicableFor(hook string) bool
	IsTrue(condition *config.Condition) bool
}
