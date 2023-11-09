package app

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/io"
)

type Context struct {
	appIO io.IO
	conf  *configuration.Configuration
	repo  *git.Repository
}

func (e *Context) IO() io.IO {
	return e.appIO
}

func (e *Context) Config() *configuration.Configuration {
	return e.conf
}

func (e *Context) Repository() *git.Repository {
	return e.repo
}

func NewContext(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) *Context {
	e := Context{
		appIO: appIO,
		conf:  conf,
		repo:  repo,
	}
	return &e
}
