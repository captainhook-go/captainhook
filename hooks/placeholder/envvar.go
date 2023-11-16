package placeholder

import (
	"github.com/captainhook-go/captainhook/hooks/app"
	"os"
)

type EnvVar struct {
	context *app.Context
}

func (r *EnvVar) Replacement(options map[string]string) string {
	envName, okEnv := options["value-of"]
	if !okEnv {
		return ""
	}

	defaultValue, okDefault := options["default-value"]
	if !okDefault {
		defaultValue = ""
	}

	envVal := os.Getenv(envName)
	if envVal == "" {
		envVal = defaultValue
	}
	return envVal
}
