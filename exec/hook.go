package exec

import (
	"fmt"
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/events"
	"github.com/captainhook-go/captainhook/exec/printer"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

type HookRunner struct {
	hook            string
	appIO           io.IO
	config          *config.Configuration
	repo            *git.Repository
	eventDispatcher *events.Dispatcher
	actionLog       *hooks.ActionLog
}

func NewHookRunner(hook string, appIO io.IO, config *config.Configuration, repo *git.Repository) *HookRunner {
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
	errActions := h.runActions()
	if errActions != nil {
		return errActions
	}
	return nil
}

func (h *HookRunner) runActions() error {
	var err error
	hookConfig := h.config.HookConfig(h.hook)

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

	if err != nil {
		// TODO: trigger failed event
		h.appIO.Write(err.Error(), true, io.NORMAL)
		return err
	}

	err = h.eventDispatcher.DispatchHookSucceededEvent(
		events.NewHookSucceededEvent(app.NewContext(h.appIO, h.config, h.repo), hookConfig, h.actionLog),
	)
	// TODO: trigger failed event
	if err != nil {
		h.appIO.Write(err.Error(), true, io.NORMAL)
	}
	return err
}

func (h *HookRunner) runActionsFailFast(hookConfig *config.Hook) error {
	for _, action := range hookConfig.GetActions() {
		err := h.runAction(action)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *HookRunner) runActionsFailLate(hookConfig *config.Hook) error {
	for _, action := range hookConfig.GetActions() {
		err := h.runAction(action)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}

func (h *HookRunner) runAction(action *config.Action) error {
	actionRunner := NewActionRunner(h.appIO, h.config, h.repo, h.eventDispatcher, h.actionLog)
	errAction := actionRunner.Run(h.hook, action)
	if errAction != nil {
		return errAction
	}
	return nil
}
