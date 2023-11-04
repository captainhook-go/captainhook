package exec

import (
	"errors"
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/io"
	"strings"
)

type Condition struct {
	cIO  io.IO
	conf *config.Configuration
	repo *git.Repository
}

func NewCondition(cIO io.IO, conf *config.Configuration, repo *git.Repository) *Condition {
	c := Condition{
		cIO,
		conf,
		repo,
	}
	return &c
}

func (c *Condition) Run(hook string, condition *config.Condition) error {
	if strings.HasPrefix(condition.Exec(), "CaptainHook::") {
		return c.runInternalCondition(hook, condition, c.cIO)
	}
	return c.runExternalAction(hook, condition, c.cIO)
}

func (c *Condition) runInternalCondition(hook string, condition *config.Condition, cIO io.IO) error {
	return errors.New("Condition: " + condition.Exec() + " failed")
}

func (c *Condition) runExternalAction(hook string, condition *config.Condition, cIO io.IO) error {
	return errors.New("Condition: " + condition.Exec() + " failed")
}
