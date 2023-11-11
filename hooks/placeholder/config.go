package placeholder

import "github.com/captainhook-go/captainhook/app"

type Config struct {
	context *app.Context
}

func (r *Config) Replacement(options map[string]string) string {
	return "--CONFIG--"
}
