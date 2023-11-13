package exec

import (
	"errors"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks/conditions"
	"github.com/captainhook-go/captainhook/io"
)

type ConditionRunner struct {
	cIO  io.IO
	conf *configuration.Configuration
	repo *git.Repository
}

func NewConditionRunner(cIO io.IO, conf *configuration.Configuration, repo *git.Repository) *ConditionRunner {
	c := ConditionRunner{
		cIO,
		conf,
		repo,
	}
	return &c
}

func (c *ConditionRunner) Run(hook string, condition *configuration.Condition) bool {
	if isInternalFunctionality(condition.Run()) {
		return c.runInternalCondition(hook, condition, c.cIO)
	}
	err := c.runExternalCondition(hook, condition, c.cIO)
	if err != nil {
		return false
	}
	return true
}

func (c *ConditionRunner) runInternalCondition(hook string, condition *configuration.Condition, cIO io.IO) bool {
	path := splitInternalPath(condition.Run())

	conditionGenerator, err := conditions.ConditionCreationFunc(path)
	if err != nil {
		cIO.Write("ConditionRunner: "+condition.Run()+"\n"+err.Error(), true, io.NORMAL)
		return false
	}

	conditionToExecute := conditionGenerator(cIO, c.conf, c.repo)
	if !conditionToExecute.IsApplicableFor(hook) {
		cIO.Write("ConditionRunner: "+condition.Run()+" nor applicable for hook "+hook, true, io.VERBOSE)
		return true
	}
	return conditionToExecute.IsTrue(condition)
}

func (c *ConditionRunner) runExternalCondition(hook string, condition *configuration.Condition, cIO io.IO) error {
	return errors.New("ConditionRunner: " + condition.Run() + " failed")
}
