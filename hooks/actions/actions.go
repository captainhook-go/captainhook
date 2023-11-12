package actions

import (
	"errors"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/actions/branch"
	"github.com/captainhook-go/captainhook/hooks/actions/file"
	"github.com/captainhook-go/captainhook/hooks/actions/message"
	"github.com/captainhook-go/captainhook/io"
	"strings"
)

func GetActionHookFunc(path []string) (func(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action, error) {
	if len(path) != 2 {
		return nil, errors.New("invalid actions functionality")
	}

	for index, value := range path {
		path[index] = strings.ToLower(value)
	}

	data := map[string]map[string]func(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action{
		"branch": {
			"preventpushoffixupandsquashcommits": branch.NewPreventPushOfFixupAndSquashCommits,
		},
		"file": {
			"doesnotcontainregex": file.NewDoesNotContainRegex,
			"isnotempty":          file.NewIsNotEmpty,
			"maxsize":             file.NewMaxSize,
		},
		"message": {
			"cacheonfail":          message.NewCacheOnFail,
			"mustfollowbeamsrules": message.NewBeamsRules,
			"mustcontainsregex":    message.NewContainsRegex,
		},
	}

	group, ok := data[path[0]]
	if !ok {
		return nil, errors.New("invalid actions functionality group")
	}

	intFunc, ok := group[path[1]]
	if !ok {
		return nil, errors.New("invalid actions functionality")
	}

	return intFunc, nil

}
