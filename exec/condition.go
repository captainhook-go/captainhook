package exec

import (
	"errors"
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks/conditions"
	"github.com/captainhook-go/captainhook/io"
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

func (c *Condition) Run(hook string, condition *config.Condition) bool {
	if isInternalFunctionality(condition.Exec()) {
		return c.runInternalCondition(hook, condition, c.cIO)
	}
	err := c.runExternalCondition(hook, condition, c.cIO)
	if err != nil {
		return false
	}
	return true
}

func (c *Condition) runInternalCondition(hook string, condition *config.Condition, cIO io.IO) bool {
	path := splitInternalPath(condition.Exec())

	conditionGenerator, err := conditions.GetConditionFunc(path)
	if err != nil {
		cIO.Write("Condition: "+condition.Exec()+"\n"+err.Error(), true, io.NORMAL)
		return false
	}

	conditionToExecute := conditionGenerator(cIO, c.conf, c.repo)
	if !conditionToExecute.IsApplicableFor(hook) {
		cIO.Write("Condition: "+condition.Exec()+" nor applicable for hook "+hook, true, io.VERBOSE)
		return true
	}
	return conditionToExecute.IsTrue(condition)
}

func (c *Condition) runExternalCondition(hook string, condition *config.Condition, cIO io.IO) error {
	return errors.New("Condition: " + condition.Exec() + " failed")
}
