package exec

import (
	"fmt"
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/events"
	"github.com/captainhook-go/captainhook/exec/printer"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"os"
)

type HookRunner struct {
	hook            string
	appIO           io.IO
	config          *configuration.Configuration
	repo            *git.Repository
	eventDispatcher *events.Dispatcher
	actionLog       *hooks.ActionLog
}

func NewHookRunner(hook string, appIO io.IO, config *configuration.Configuration, repo *git.Repository) *HookRunner {
	h := HookRunner{
		hook:            hook,
		appIO:           appIO,
		config:          config,
		repo:            repo,
		eventDispatcher: events.NewDispatcher(),
		actionLog:       hooks.NewActionLog(),
	}

	defaultPrinter := printer.NewDefaultPrinter(appIO)
	defaultPrinter.RegisterSubscribers(h.eventDispatcher)

	return &h
}

func (h *HookRunner) Run() error {
	if shouldHooksBeSkipped() {
		return nil
	}

	errActions := h.runActions()
	if errActions != nil {
		return errActions
	}
	return nil
}

func shouldHooksBeSkipped() bool {
	for _, envName := range []string{"CAPTAINHOOK_SKIP_HOOKS", "CI"} {
		skip := os.Getenv(envName)
		if skip == "1" {
			return true
		}
	}
	return false
}

func (h *HookRunner) runActions() error {
	var err error
	hookConfig := h.getHookConfig()

	err = h.eventDispatcher.DispatchHookStartedEvent(
		events.NewHookStartedEvent(app.NewContext(h.appIO, h.config, h.repo), hookConfig),
	)
	// TODO: trigger failed event
	if err != nil {
		h.appIO.Write(err.Error(), true, io.NORMAL)
		return err
	}

	if h.config.FailOnFirstError() {
		err = h.runActionsFailFast(hookConfig)
	} else {
		err = h.runActionsFailLate(hookConfig)
	}

	// TODO: handle dispatcher errors
	if err != nil {
		_ = h.eventDispatcher.DispatchHookFailedEvent(
			events.NewHookFailedEvent(app.NewContext(h.appIO, h.config, h.repo), hookConfig, h.actionLog, err),
		)
		return err
	}

	_ = h.eventDispatcher.DispatchHookSucceededEvent(
		events.NewHookSucceededEvent(app.NewContext(h.appIO, h.config, h.repo), hookConfig, h.actionLog),
	)
	return err
}

func (h *HookRunner) runActionsFailFast(hookConfig *configuration.Hook) error {
	for _, action := range hookConfig.GetActions() {
		err := h.runAction(action)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *HookRunner) runActionsFailLate(hookConfig *configuration.Hook) error {
	failed := 0
	for _, action := range hookConfig.GetActions() {
		err := h.runAction(action)
		if err != nil {
			failed++
		}
	}
	if failed > 0 {
		plural := ""
		if failed > 1 {
			plural = "s"
		}
		return fmt.Errorf("%d action%s failed", failed, plural)
	}
	return nil
}

func (h *HookRunner) runAction(action *configuration.Action) error {
	actionRunner := NewActionRunner(h.appIO, h.config, h.repo, h.eventDispatcher, h.actionLog)
	actionErr, dispatchErr := actionRunner.Run(h.hook, action)
	// TODO: propagate dispatch error
	if dispatchErr != nil {
		h.appIO.Write(fmt.Sprintf("error dispatching events: %s", dispatchErr.Error()), true, io.NORMAL)
	}
	return actionErr
}

func (h *HookRunner) getHookConfig() *configuration.Hook {
	hookConfig := h.config.HookConfig(h.hook)
	vHook, ok := info.VirtualHook(h.hook)
	if ok {
		vHookConfig := h.config.HookConfig(vHook)
		hookConfig = configuration.NewHook(
			h.hook+" ("+vHook+")",
			true,
		)
		for _, action := range vHookConfig.GetActions() {
			hookConfig.AddAction(action)
		}
	}

	return hookConfig
}
