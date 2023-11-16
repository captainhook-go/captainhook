package conditions

import (
	"errors"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/conditions/filechanged"
	"github.com/captainhook-go/captainhook/hooks/conditions/filestaged"
	"github.com/captainhook-go/captainhook/hooks/conditions/inconfig"
	"github.com/captainhook-go/captainhook/hooks/conditions/status"
	"github.com/captainhook-go/captainhook/io"
	"strings"
)

var (
	conditionCreationConfig = map[string]map[string]func(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Condition{
		"inconfig": {
			"customvalueistruthy": inconfig.NewCustomValueIsTruthy,
			"customvalueisfalsy":  inconfig.NewCustomValueIsFalsy,
		},
		"filechanged": {
			"any":    filechanged.NewAny,
			"all":    filechanged.NewAll,
			"thatis": filechanged.NewThatIs,
		},
		"filestaged": {
			"all":    filestaged.NewAll,
			"any":    filestaged.NewAny,
			"thatis": filestaged.NewThatIs,
		},
		"status": {
			"onbranch": status.NewOnBranch,
		},
	}
)

func ConditionCreationFunc(path []string) (func(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Condition, error) {
	if len(path) != 2 {
		return nil, errors.New("invalid condition functionality")
	}

	for index, value := range path {
		path[index] = strings.ToLower(value)
	}

	group, ok := conditionCreationConfig[path[0]]
	if !ok {
		return nil, errors.New("invalid condition functionality group")
	}
	intFunc, ok := group[path[1]]
	if !ok {
		return nil, errors.New("invalid condition functionality")
	}
	return intFunc, nil
}
