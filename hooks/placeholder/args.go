package placeholder

import (
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/io"
)

type Args struct {
	context *app.Context
}

func (r *Args) Replacement(options map[string]string) string {
	arg := io.MappedStringOrDefault(options, "value-of", "_")
	def := io.MappedStringOrDefault(options, "default", "")

	return r.context.IO().Argument(placeholderToArg(arg), def)
}
