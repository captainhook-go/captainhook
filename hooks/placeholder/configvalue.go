package placeholder

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/hooks/app"
)

type ConfigValue struct {
	context *app.Context
}

func (r *ConfigValue) Replacement(options map[string]string) string {
	replaceWith, okValue := options["value-of"]
	if !okValue {
		return ""
	}

	defaultValue, okDefault := options["default-value"]
	if !okDefault {
		defaultValue = ""
	}

	switch replaceWith {
	case "git-directory":
		return r.context.Repository().Path()
	case "verbosity":
		return configuration.UnMapVerbosity(r.context.Config().Verbosity())
	default:
		val, okCustom := r.context.Config().CustomSettings()[replaceWith]
		if !okCustom {
			return defaultValue
		}
		return val
	}
}
