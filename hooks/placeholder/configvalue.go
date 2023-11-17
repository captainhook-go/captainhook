package placeholder

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/io"
)

type ConfigValue struct {
	context *app.Context
}

func (r *ConfigValue) Replacement(options map[string]string) string {
	replaceWith, okValue := options["value-of"]
	if !okValue {
		return ""
	}

	defaultValue := io.MappedStringOrDefault(options, "default", "")

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
