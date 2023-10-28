package exec

import (
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/io"
)

type Installer struct {
	appIO        io.IO
	config       *config.Configuration
	repo         *git.Repository
	force        bool
	skipExisting bool
}

func NewInstaller(appIO io.IO, config *config.Configuration, repo *git.Repository) *Installer {
	i := Installer{appIO: appIO, config: config, repo: repo, force: false, skipExisting: false}
	return &i
}

func (i *Installer) SkipExisting(skip bool) {
	i.skipExisting = skip
}

func (i *Installer) Force(force bool) {
	i.force = force
}

func (i *Installer) Run() error {
	i.appIO.Write("Installing CaptainHook START", true, io.NORMAL)
	i.appIO.Write(i.config.Path(), true, io.NORMAL)
	i.appIO.Write(i.repo.Path(), true, io.NORMAL)
	i.appIO.Write("Installing CaptainHook END", true, io.NORMAL)
	return nil
}
