package placeholder

import (
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/io"
)

type StdIn struct {
	context *app.Context
}

func (r *StdIn) Replacement(options map[string]string) string {
	esc := "'"
	if !io.AnswerToBool(io.MappedStringOrDefault(options, "escaped", "true")) {
		esc = ""
	}
	return esc + r.context.IO().Option("input", "") + esc
}
