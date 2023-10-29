package exec

import (
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/io"
)

type HookRunner struct {
	hook         string
	appIO        io.IO
	config       *config.Configuration
	repo         *git.Repository
	beforeHook   func(appIO io.IO, config *config.Configuration, repo *git.Repository) error
	beforeAction func(appIO io.IO, config *config.Action, repo *git.Repository) error
	afterAction  func(appIO io.IO, config *config.Action, repo *git.Repository) error
	afterHook    func(appIO io.IO, config *config.Configuration, repo *git.Repository) error
}

func NewHookRunner(hook string, appIO io.IO, config *config.Configuration, repo *git.Repository) *HookRunner {
	h := HookRunner{hook: hook, appIO: appIO, config: config, repo: repo}
	return &h
}

func (h *HookRunner) Run() error {
	errBefore := h.BeforeHook()
	if errBefore != nil {
		return errBefore
	}

	errActions := h.runActions()
	if errActions != nil {
		return errActions
	}

	errAfter := h.AfterHook()
	if errAfter != nil {
		return errAfter
	}
	return nil
}

func (h *HookRunner) BeforeHook() error {
	return h.beforeHook(h.appIO, h.config, h.repo)
}

func (h *HookRunner) AfterHook() error {
	return h.afterHook(h.appIO, h.config, h.repo)
}

func (h *HookRunner) runActions() error {
	var err error
	hookConfig := h.config.HookConfig(h.hook)

	if h.config.FailOnFirstError() {
		err = h.runActionsFailFirst(hookConfig)
	} else {
		err = h.runActionsFailLast(hookConfig)
	}

	return err
}

func (h *HookRunner) runActionsFailFirst(hookConfig *config.Hook) error {
	for _, action := range hookConfig.GetActions() {
		err := h.runAction(action)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *HookRunner) runActionsFailLast(hookConfig *config.Hook) error {
	for _, action := range hookConfig.GetActions() {
		err := h.runAction(action)
		if err != nil {

		}
	}
	return nil
}

func (h *HookRunner) runAction(action *config.Action) error {
	errBefore := h.beforeAction(h.appIO, action, h.repo)
	if errBefore != nil {
		return errBefore
	}
	h.appIO.Write(action.Action(), true, io.NORMAL)

	errAfter := h.afterAction(h.appIO, action, h.repo)
	if errAfter != nil {
		return errAfter
	}
	return nil
}
