package placeholder

import (
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/hooks/util"
	"github.com/captainhook-go/captainhook/io"
	"strings"
)

type FileList struct {
	name    string
	context *app.Context
	files   []string
}

func (r *FileList) Replacement(options map[string]string) string {
	r.context.IO().Write("<comment>placeholder: "+r.name+"</comment>", true, io.VERBOSE)
	r.filterByType(options)
	r.filterByDirectory(options)

	separator, ok := options["separated-by"]
	if !ok {
		separator = " "
	}
	return strings.Join(r.files, separator)
}

func (r *FileList) filterByType(options map[string]string) {
	suffix, ok := options["of-type"]
	if ok {
		r.context.IO().Write("  filter files by type: "+suffix+"", true, io.DEBUG)
		r.files = util.FilterByType(r.files, "."+suffix)
	}
}

func (r *FileList) filterByDirectory(options map[string]string) {
	dir, ok := options["in-directory"]
	if ok {
		r.context.IO().Write("  filter files by directory: "+dir, true, io.DEBUG)
		r.files = util.FilterByDirectory(r.files, dir)
	}
}
