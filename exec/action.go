package exec

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/events"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/actions"
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
)

// ActionRunner executes an action
// Besides executing the action the runner triggers the relevant events
type ActionRunner struct {
	appIO           io.IO
	conf            *configuration.Configuration
	repo            git.Repo
	eventDispatcher *events.Dispatcher
}

// RunAsync run the action concurrently and send the result to a channel
func (a *ActionRunner) RunAsync(hook string, action *configuration.Action, channel chan *ActionResult) {
	channel <- a.Run(hook, action)
}

// Run executes the action and returns an ActionResult struct
func (a *ActionRunner) Run(hook string, action *configuration.Action) *ActionResult {
	cIO := io.NewCollectorIO(a.appIO.Verbosity(), a.appIO.Input())

	errDispatchStart := a.eventDispatcher.DispatchActionStartedEvent(
		events.NewActionStartedEvent(app.NewContext(a.appIO, a.conf, a.repo), action),
	)
	if errDispatchStart != nil {
		return NewActionResult(action, info.ActionFailed, nil, errDispatchStart, cIO)
	}

	if !a.doConditionsApply(hook, action.Conditions(), cIO) {
		errDispatchSkipped := a.eventDispatcher.DispatchActionSkippedEvent(
			events.NewActionSkippedEvent(app.NewContext(a.appIO, a.conf, a.repo), action),
		)
		return NewActionResult(action, info.ActionSkipped, nil, errDispatchSkipped, cIO)
	}

	errRun := a.runAction(hook, action, cIO)

	if errRun != nil {
		cIO.Write(errRun.Error(), true, io.NORMAL)
		errDispatchFailed := a.eventDispatcher.DispatchActionFailedEvent(
			events.NewActionFailedEvent(app.NewContext(a.appIO, a.conf, a.repo), action, errRun),
		)
		return NewActionResult(action, info.ActionFailed, errRun, errDispatchFailed, cIO)
	}
	errDispatchSuccess := a.eventDispatcher.DispatchActionSucceededEvent(
		events.NewActionSucceededEvent(app.NewContext(a.appIO, a.conf, a.repo), action),
	)
	return NewActionResult(action, info.ActionSucceeded, errRun, errDispatchSuccess, cIO)
}

// runAction checks if the action is not restricted from running during this hook
// If the action is applicable it executes it
// If not it triggers an ActionSkippedEvent
func (a *ActionRunner) runAction(hook string, action *configuration.Action, cIO *io.CollectorIO) error {
	actionToExecute, err := a.createAction(action, cIO)
	if err != nil {
		return err
	}
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

// createAction creates the Action struct
// Either some special Action for some internal functionality or the generic ExternalCommand Action to execute some
// custom executable.
func (a *ActionRunner) createAction(action *configuration.Action, cIO *io.CollectorIO) (hooks.Action, error) {
	if isInternalFunctionality(action.Run()) {
		return a.createInternalAction(action, cIO)
	}
	return a.createExternalAction(cIO)
}

// createInternalAction creates one of CaptainHooks integrated actions
func (a *ActionRunner) createInternalAction(action *configuration.Action, cIO *io.CollectorIO) (hooks.Action, error) {
	path := splitInternalPath(action.Run())
	actionGenerator, err := actions.ActionCreationFunc(path)
	if err != nil {
		return nil, err
	}
	return actionGenerator(cIO, a.conf, a.repo), nil
}

// createExternalAction creates the Action that executes external executables
func (a *ActionRunner) createExternalAction(cIO *io.CollectorIO) (hooks.Action, error) {
	return actions.NewExternalCommand(cIO, a.conf, a.repo), nil
}

// doConditionsApply answers if an Action should be executed for a specific hook
// In order to be executed all conditions have to be true.
// The only exception are conditions within a logic OR condition which is handled a layer blow.
func (a *ActionRunner) doConditionsApply(hook string, conditions []*configuration.Condition, cIO *io.CollectorIO) bool {
	return DoAllConditionsApply(app.NewContext(cIO, a.conf, a.repo), conditions, hook)
}

func NewActionRunner(
	appIO io.IO,
	conf *configuration.Configuration,
	repo git.Repo,
	dispatcher *events.Dispatcher,
) *ActionRunner {
	return &ActionRunner{appIO: appIO, conf: conf, repo: repo, eventDispatcher: dispatcher}
}
