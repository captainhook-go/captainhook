package placeholder

import (
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/io"
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

	defaultValue := io.MappedStringOrDefault(options, "default", "")

	envVal := os.Getenv(envName)
	if envVal == "" {
		envVal = defaultValue
	}
	return envVal
}
