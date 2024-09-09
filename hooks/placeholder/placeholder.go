package placeholder

import (
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/hooks/input"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"regexp"
	"slices"
	"strings"
)

var (
	placeholders = map[string]func(aContext *app.Context) Replacer{
		"ARG": func(aContext *app.Context) Replacer {
			return &Args{context: aContext}
		},
		"CONFIG": func(aContext *app.Context) Replacer {
			return &ConfigValue{context: aContext}
		},
		"CHANGED_FILES": func(aContext *app.Context) Replacer {
			return &FileList{name: "CHANGED_FILES", context: aContext, files: collectAllChangedFiles(aContext)}
		},
		"ENV": func(aContext *app.Context) Replacer {
			return &EnvVar{context: aContext}
		},
		"STAGED_FILES": func(aContext *app.Context) Replacer {
			files, _ := aContext.Repository().StagedFiles()
			return &FileList{name: "STAGED_FILES", context: aContext, files: files}
		},
		"STDIN": func(aContext *app.Context) Replacer {
			return &StdIn{context: aContext}
		},
	}
)

type Replacer interface {
	Replacement(options map[string]string) string
}

func ReplacePlaceholders(aContext *app.Context, input string) string {

	placeholderRegex := regexp.MustCompile(`{\$([A-Za-z_\-]+)(\|[^{}]+)*}`)

	// Replace placeholders in the input string
	result := placeholderRegex.ReplaceAllStringFunc(input, func(match string) string {
		options := map[string]string{}
		matches := placeholderRegex.FindStringSubmatch(match)
		if len(matches) == 3 {
			name := matches[1]
			optsAll := matches[2]

			// to conveniently access arguments directly like:
			// {$MESSAGE_FILE}
			// we have to check for this case
			// and inject the correct replacer syntax
			// {$ARG|value-of:MESSAGE_FILE}
			if optsAll == "" {
				if isValidArg(name) {
					optsAll = "|value-of:" + name
					name = "ARG"
				}
			}

			for _, opt := range strings.Split(io.SubString(optsAll, 1, 0), "|") {
				optParts := strings.Split(opt, ":")
				if len(optParts) == 2 {
					options[optParts[0]] = optParts[1]
				}
			}

			replacerCreationFunc, ok := placeholders[name]
			if !ok {
				aContext.IO().Write("no replacer found for type: "+name, true, io.VERBOSE)
				return ""
			}
			replacer := replacerCreationFunc(aContext)
			return replacer.Replacement(options)
		}
		aContext.IO().Write("placeholder not valid: "+match, true, io.VERBOSE)
		return ""
	})

	return result
}

func isValidArg(placeholder string) bool {
	return slices.Contains(
		info.AllHookArguments(),
		placeholderToArg(placeholder),
	)
}

func argToPlaceholder(arg string) string {
	return strings.Replace(strings.ToUpper(arg), "-", "_", -1)
}

func placeholderToArg(placeholder string) string {
	return strings.Replace(strings.ToLower(placeholder), "_", "-", -1)
}

func collectAllChangedFiles(aContext *app.Context) []string {
	var files []string
	unique := map[string]bool{}
	for _, aRange := range input.DetectRanges(aContext.IO()) {
		changed, _ := aContext.Repository().ChangedFiles(aRange.From().Id(), aRange.To().Id())
		for _, file := range changed {
			unique[file] = true
		}
	}
	for file := range unique {
		files = append(files, file)
	}
	return files
}
