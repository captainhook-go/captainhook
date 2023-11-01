package exec

import (
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/events"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/actions"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"os/exec"
	"strings"
)

type ActionRunner struct {
	appIO           io.IO
	conf            *config.Configuration
	repo            *git.Repository
	eventDispatcher *events.Dispatcher
	actionLog       *hooks.ActionLog
}

func NewActionRunner(appIO io.IO, conf *config.Configuration, repo *git.Repository, dispatcher *events.Dispatcher, log *hooks.ActionLog) *ActionRunner {
	a := ActionRunner{appIO: appIO, conf: conf, repo: repo, eventDispatcher: dispatcher, actionLog: log}
	return &a
}

func (a *ActionRunner) Run(hook string, action *config.Action) error {

	err := a.eventDispatcher.DispatchActionStartedEvent(events.NewActionStartedEvent(app.NewContext(a.appIO, a.conf, a.repo), action))
	if err != nil {
		return err
	}

	cIO := io.NewCollectorIO(a.appIO.Verbosity(), a.appIO.Arguments())
	err = a.runAction(hook, action, cIO)

	var status int

	if err != nil {
		err = a.eventDispatcher.DispatchActionFailedEvent(events.NewActionFailedEvent(app.NewContext(a.appIO, a.conf, a.repo), action, err))
		status = info.ACTION_FAILED
	} else {
		err = a.eventDispatcher.DispatchActionSucceededEvent(events.NewActionSucceededEvent(app.NewContext(a.appIO, a.conf, a.repo), action))
		status = info.ACTION_SUCCEEDED
	}
	a.appendActionLog(action, cIO, status)
	return err

}

func (a *ActionRunner) runAction(hook string, action *config.Action, cIO *io.CollectorIO) error {
	if strings.HasPrefix(action.Action(), "CaptainHook::") {
		return a.runInternalAction(hook, action, cIO)
	}

	return a.runExternalAction(hook, action, cIO)
}

func (a *ActionRunner) runInternalAction(hook string, action *config.Action, aIO *io.CollectorIO) error {
	actionPath := strings.Split(action.Action(), "::")[1]
	path := strings.Split(actionPath, ".")

	actionGenerator, err := actions.GetActionHookFunc(path)
	if err != nil {
		return err
	}

	var actionToExecute hooks.Action
	actionToExecute = actionGenerator(aIO, a.conf, a.repo)

	_, ok := actionToExecute.(events.Subscriber)
	if ok {

	}

	if !actionToExecute.IsApplicableFor(hook) {
		// TODO: trigger action skipped event
		a.appIO.Write("skipped", true, io.NORMAL)
		return nil
	}
	return actionToExecute.Run(action)
}

func (a *ActionRunner) runExternalAction(hook string, action *config.Action, aIO *io.CollectorIO) error {
	splits := strings.Split(action.Action(), " ")

	cmd := exec.Command(splits[0], splits[1:]...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		message := ""
		if len(out) > 0 {
			message = message + string(out)
		}
		message = message + err.Error()
		aIO.Write(message, true, io.NORMAL)
		return err
	}

	if len(out) > 0 {
		aIO.Write(string(out), false, io.VERBOSE)
	}
	return nil
}

func (a *ActionRunner) appendActionLog(action *config.Action, cIO *io.CollectorIO, status int) {
	a.actionLog.Add(hooks.NewActionLogItem(action, cIO, status))
}
