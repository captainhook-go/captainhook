package exec

import (
	"fmt"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/events"
	"github.com/captainhook-go/captainhook/exec/printer"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"os"
	"sync"
	"time"
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
	var err error
	err = h.eventDispatcher.DispatchHookStartedEvent(
		events.NewHookStartedEvent(app.NewContext(h.appIO, h.config, h.repo), h.getHookConfig()),
	)
	if err != nil {
		h.appIO.Write(err.Error(), true, io.NORMAL)
		return err
	}

	if h.shouldHooksBeSkipped() {
		return nil
	}

	errActions := h.runActions()
	if errActions != nil {
		return errActions
	}
	return nil
}

func (h *HookRunner) shouldHooksBeSkipped() bool {
	for _, envName := range []string{"CAPTAINHOOK_SKIP_HOOKS", "CI"} {
		skip := os.Getenv(envName)
		if skip == "1" {
			h.appIO.Write(" - skipped because of ENV variable "+envName+" is set to 1", true, io.NORMAL)
			return true
		}
	}
	if h.hook == info.CommitMsg {
		msg, err := h.repo.CommitMessage(h.appIO.Argument(info.ArgCommitMsgFile, ""))
		if err != nil {
			return false
		}
		if msg.IsFixup() || msg.IsSquash() {
			h.appIO.Write(" - no message validation for fixup & squash commits: skipping all actions", true, io.NORMAL)
			return true
		}
	}
	return false
}

func (h *HookRunner) runActions() error {
	var err error
	start := time.Now()
	hookConfig := h.getHookConfig()

	if h.config.FailOnFirstError() {
		err = h.runActionsFailFast(hookConfig)
	} else if h.config.RunAsync() {
		err = h.runActionsAsync(hookConfig)
	} else {
		err = h.runActionsFailLate(hookConfig)
	}
	executionTime := time.Since(start)

	if err != nil {
		_ = h.eventDispatcher.DispatchHookFailedEvent(
			events.NewHookFailedEvent(
				app.NewContext(h.appIO, h.config, h.repo), hookConfig, h.actionLog, executionTime, err,
			),
		)
		return err
	}

	_ = h.eventDispatcher.DispatchHookSucceededEvent(
		events.NewHookSucceededEvent(
			app.NewContext(h.appIO, h.config, h.repo), hookConfig, h.actionLog, executionTime,
		),
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

func (h *HookRunner) runActionsAsync(hookConfig *configuration.Hook) error {
	channel := make(chan *ActionResult)
	var wg sync.WaitGroup
	for _, action := range hookConfig.GetActions() {
		wg.Add(1)
		go h.runActionAsync(action, channel, &wg)
	}
	go func() {
		wg.Wait()
		close(channel)
	}()

	failed := 0
	for result := range channel {
		h.actionLog.Add(hooks.NewActionLogItem(result.Config, result.Log, result.Status))
		if result.RunErr != nil {
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

func (h *HookRunner) runActionAsync(action *configuration.Action, channel chan *ActionResult, wg *sync.WaitGroup) {
	actionRunner := NewActionRunner(h.appIO, h.config, h.repo, h.eventDispatcher)
	actionRunner.RunAsync(h.hook, action, channel)
	defer wg.Done()
}

func (h *HookRunner) runAction(action *configuration.Action) error {
	actionRunner := NewActionRunner(h.appIO, h.config, h.repo, h.eventDispatcher)
	result := actionRunner.Run(h.hook, action)

	h.actionLog.Add(hooks.NewActionLogItem(action, result.Log, result.Status))

	if result.DispatchErr != nil {
		h.appIO.Write(fmt.Sprintf("error dispatching events: %s", result.DispatchErr.Error()), true, io.NORMAL)
	}
	return result.RunErr
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
