package exec

import (
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/events"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/actions"
	"github.com/captainhook-go/captainhook/hooks/placeholder"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"os/exec"
	"strings"
)

type ActionRunner struct {
	appIO           io.IO
	conf            *configuration.Configuration
	repo            *git.Repository
	eventDispatcher *events.Dispatcher
	actionLog       *hooks.ActionLog
}

func NewActionRunner(
	appIO io.IO,
	conf *configuration.Configuration,
	repo *git.Repository,
	dispatcher *events.Dispatcher,
	log *hooks.ActionLog,
) *ActionRunner {
	return &ActionRunner{appIO: appIO, conf: conf, repo: repo, eventDispatcher: dispatcher, actionLog: log}
}

func (a *ActionRunner) Run(hook string, action *configuration.Action) *ActionResult {
	cIO := io.NewCollectorIO(a.appIO.Verbosity(), a.appIO.Input())

	errDispatchStart := a.eventDispatcher.DispatchActionStartedEvent(
		events.NewActionStartedEvent(app.NewContext(a.appIO, a.conf, a.repo), action),
	)
	if errDispatchStart != nil {
		a.appendActionLog(action, cIO, info.ACTION_FAILED)
		return NewActionResult(info.ACTION_FAILED, nil, errDispatchStart)
	}

	if !a.doConditionsApply(hook, action.Conditions(), cIO) {
		errDispatchSkipped := a.eventDispatcher.DispatchActionSkippedEvent(
			events.NewActionSkippedEvent(app.NewContext(a.appIO, a.conf, a.repo), action),
		)
		a.appendActionLog(action, cIO, info.ACTION_SKIPPED)
		return NewActionResult(info.ACTION_SKIPPED, nil, errDispatchSkipped)
	}

	errRun := a.runAction(hook, action, cIO)

	if errRun != nil {
		cIO.Write(errRun.Error(), true, io.NORMAL)
		errDispatchFailed := a.eventDispatcher.DispatchActionFailedEvent(
			events.NewActionFailedEvent(app.NewContext(a.appIO, a.conf, a.repo), action, errRun),
		)
		a.appendActionLog(action, cIO, info.ACTION_FAILED)
		return NewActionResult(info.ACTION_FAILED, errRun, errDispatchFailed)
	}
	errDispatchSuccess := a.eventDispatcher.DispatchActionSucceededEvent(
		events.NewActionSucceededEvent(app.NewContext(a.appIO, a.conf, a.repo), action),
	)

	a.appendActionLog(action, cIO, info.ACTION_SUCCEEDED)
	return NewActionResult(info.ACTION_SUCCEEDED, errRun, errDispatchSuccess)
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

	splits := strings.Split(commandToExecute, " ")
	cmd := exec.Command(splits[0], splits[1:]...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		if len(out) > 0 {
			aIO.Write("<comment>output:</comment>\n"+string(out), true, io.NORMAL)
		}
		return err
	}

	if len(out) > 0 {
		aIO.Write("<comment>output:</comment>\n"+string(out), false, io.VERBOSE)
	}
	return nil
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

func (a *ActionRunner) appendActionLog(action *configuration.Action, cIO *io.CollectorIO, status int) {
	a.actionLog.Add(hooks.NewActionLogItem(action, cIO, status))
}
