package placeholder

import (
	"github.com/captainhook-go/captainhook/hooks/app"
)

type Config struct {
	context *app.Context
}

func (r *Config) Replacement(options map[string]string) string {
	return "--CONFIG--"
}
