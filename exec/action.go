package exec

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/events"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/actions"
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/hooks/placeholder"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

type ActionRunner struct {
	appIO           io.IO
	conf            *configuration.Configuration
	repo            *git.Repository
	eventDispatcher *events.Dispatcher
}

func NewActionRunner(
	appIO io.IO,
	conf *configuration.Configuration,
	repo *git.Repository,
	dispatcher *events.Dispatcher,
) *ActionRunner {
	return &ActionRunner{appIO: appIO, conf: conf, repo: repo, eventDispatcher: dispatcher}
}

func (a *ActionRunner) RunAsync(hook string, action *configuration.Action, channel chan *ActionResult) {
	channel <- a.Run(hook, action)
}

func (a *ActionRunner) Run(hook string, action *configuration.Action) *ActionResult {
	cIO := io.NewCollectorIO(a.appIO.Verbosity(), a.appIO.Input())

	errDispatchStart := a.eventDispatcher.DispatchActionStartedEvent(
		events.NewActionStartedEvent(app.NewContext(a.appIO, a.conf, a.repo), action),
	)
	if errDispatchStart != nil {
		return NewActionResult(action, info.ACTION_FAILED, nil, errDispatchStart, cIO)
	}

	if !a.doConditionsApply(hook, action.Conditions(), cIO) {
		errDispatchSkipped := a.eventDispatcher.DispatchActionSkippedEvent(
			events.NewActionSkippedEvent(app.NewContext(a.appIO, a.conf, a.repo), action),
		)
		return NewActionResult(action, info.ACTION_SKIPPED, nil, errDispatchSkipped, cIO)
	}

	errRun := a.runAction(hook, action, cIO)

	if errRun != nil {
		cIO.Write(errRun.Error(), true, io.NORMAL)
		errDispatchFailed := a.eventDispatcher.DispatchActionFailedEvent(
			events.NewActionFailedEvent(app.NewContext(a.appIO, a.conf, a.repo), action, errRun),
		)
		return NewActionResult(action, info.ACTION_FAILED, errRun, errDispatchFailed, cIO)
	}
	errDispatchSuccess := a.eventDispatcher.DispatchActionSucceededEvent(
		events.NewActionSucceededEvent(app.NewContext(a.appIO, a.conf, a.repo), action),
	)
	return NewActionResult(action, info.ACTION_SUCCEEDED, errRun, errDispatchSuccess, cIO)
}

func (a *ActionRunner) runAction(hook string, action *configuration.Action, cIO *io.CollectorIO) error {
	if isInternalFunctionality(action.Run()) {
		return a.runInternalAction(hook, action, cIO)
	}
	return a.runExternalAction(hook, action, cIO)
}

func (a *ActionRunner) runInternalAction(hook string, action *configuration.Action, cIO *io.CollectorIO) error {
	path := splitInternalPath(action.Run())

	actionGenerator, err := actions.ActionCreationFunc(path)
	if err != nil {
		return err
	}

	var actionToExecute hooks.Action
	actionToExecute = actionGenerator(cIO, a.conf, a.repo)
	if value, ok := interface{}(actionToExecute).(events.EventSubscriber); ok {
		subErr := value.Subscribe(a.eventDispatcher, action)
		if subErr != nil {
			return subErr
		}
	}

	if !actionToExecute.IsApplicableFor(hook) {
		cIO.Write("action not applicable for hook: "+hook, true, io.VERBOSE)
		return a.eventDispatcher.DispatchActionSkippedEvent(events.NewActionSkippedEvent(app.NewContext(a.appIO, a.conf, a.repo), action))
	}
	return actionToExecute.Run(action)
}

func (a *ActionRunner) runExternalAction(hook string, action *configuration.Action, aIO *io.CollectorIO) error {
	commandToExecute := placeholder.ReplacePlaceholders(app.NewContext(aIO, a.conf, a.repo), action.Run())
	// if there were placeholders replaced
	if commandToExecute != action.Run() {
		aIO.Write("<comment>cmd:</comment>\n"+commandToExecute, true, io.VERBOSE)
	}
	return ExecuteCommand(aIO, commandToExecute)
}

func (a *ActionRunner) doConditionsApply(hook string, conditions []*configuration.Condition, cIO *io.CollectorIO) bool {
	conditionRunner := NewConditionRunner(cIO, a.conf, a.repo)
	for _, condition := range conditions {
		if !conditionRunner.Run(hook, condition) {
			return false
		}
	}
	return true
}
