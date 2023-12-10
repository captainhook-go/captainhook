package hooks

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/io"
)

type HookBundle struct {
	AppIO       io.IO
	Conf        *configuration.Configuration
	Repo        git.Repo
	Restriction *Restriction
}

func NewHookBundle(appIO io.IO, conf *configuration.Configuration, repo git.Repo, hooks []string) *HookBundle {
	b := HookBundle{
		AppIO:       appIO,
		Conf:        conf,
		Repo:        repo,
		Restriction: NewRestriction(hooks),
	}
	return &b
}
