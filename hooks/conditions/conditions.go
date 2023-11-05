package conditions

import (
	"errors"
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/conditions/fileChanged"
	"github.com/captainhook-go/captainhook/hooks/conditions/fileStaged"
	"github.com/captainhook-go/captainhook/hooks/conditions/inConfig"
	"github.com/captainhook-go/captainhook/io"
	"strings"
)

func GetConditionFunc(path []string) (func(appIO io.IO, conf *config.Configuration, repo *git.Repository) hooks.Condition, error) {
	if len(path) != 2 {
		return nil, errors.New("invalid condition functionality")
	}

	for index, value := range path {
		path[index] = strings.ToLower(value)
	}

	data := map[string]map[string]func(appIO io.IO, conf *config.Configuration, repo *git.Repository) hooks.Condition{
		"inconfig": {
			"customvalueistruthy": inConfig.NewCustomValueIsTruthy,
			"customvalueisfalsy":  inConfig.NewCustomValueIsFalsy,
		},
		"filechanged": {
			"any":    fileChanged.NewAny,
			"all":    fileChanged.NewAll,
			"oftype": fileChanged.NewOfType,
		},
		"filestaged": {
			"all":         fileStaged.NewAll,
			"any":         fileStaged.NewAny,
			"indirectory": fileStaged.NewInDirectory,
			"oftype":      fileStaged.NewOfType,
			"thatis":      fileStaged.NewThatIs,
		},
	}

	group, ok := data[path[0]]
	if !ok {
		return nil, errors.New("invalid condition functionality group")
	}
	intFunc, ok := group[path[1]]
	if !ok {
		return nil, errors.New("invalid condition functionality")
	}
	return intFunc, nil
}
