package exec

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
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
	conditionToExecute, err := c.crateCondition(condition)

	if err != nil {
		return false
	}

	if !conditionToExecute.IsApplicableFor(hook) {
		c.cIO.Write("ConditionRunner: "+condition.Run()+" nor applicable for hook "+hook, true, io.VERBOSE)
		return true
	}
	return conditionToExecute.IsTrue(condition)
}

func (c *ConditionRunner) crateCondition(condition *configuration.Condition) (hooks.Condition, error) {
	if isInternalFunctionality(condition.Run()) {
		return c.createInternalCondition(condition)
	}
	return c.createExternalCondition()
}

func (c *ConditionRunner) createInternalCondition(condition *configuration.Condition) (hooks.Condition, error) {
	path := splitInternalPath(condition.Run())
	conditionGenerator, err := conditions.ConditionCreationFunc(path)
	if err != nil {
		c.cIO.Write("ConditionRunner: "+condition.Run()+"\n"+err.Error(), true, io.NORMAL)
		return nil, err
	}
	return conditionGenerator(c.cIO, c.conf, c.repo), nil
}

func (c *ConditionRunner) createExternalCondition() (hooks.Condition, error) {
	return conditions.NewExternalCommand(c.cIO, c.conf, c.repo), nil
}
