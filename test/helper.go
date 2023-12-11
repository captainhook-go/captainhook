package test

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/io"
)

func CreateFakeIO() io.IO {
	return &IOMock{}
}

func CreateFakeConfig() *configuration.Configuration {
	return configuration.NewConfiguration("captain.json", false)
}

func CreateFakeRepo() *RepoMock {
	return &RepoMock{path: "./", branch: "main"}
}

func CreateFakeHookContext(inOut io.IO, conf *configuration.Configuration, repo git.Repo) *app.Context {
	return app.NewContext(inOut, conf, repo)
}
