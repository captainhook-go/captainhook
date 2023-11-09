package printer

import (
	"fmt"
	"github.com/captainhook-go/captainhook/events"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"strings"
)

type DefaultPrinter struct {
	appIO io.IO
}

func NewDefaultPrinter(appIO io.IO) *DefaultPrinter {
	p := DefaultPrinter{
		appIO: appIO,
	}
	return &p
}

func (p *DefaultPrinter) RegisterSubscribers(dispatcher *events.Dispatcher) {
	dispatcher.RegisterHookStartedSubscribers(NewDefaultHookStartedSubscriber(p.appIO))
	dispatcher.RegisterHookSucceededSubscribers(NewDefaultHookSucceededSubscriber(p.appIO))
	dispatcher.RegisterHookFailedSubscribers(NewDefaultHookFailedSubscriber(p.appIO))
	dispatcher.RegisterActionStartedSubscribers(NewDefaultActionStartedSubscriber(p.appIO))
	dispatcher.RegisterActionSucceededSubscribers(NewDefaultActionSucceededSubscriber(p.appIO))
	dispatcher.RegisterActionSkippedSubscribers(NewDefaultActionSkippedSubscriber(p.appIO))
	dispatcher.RegisterActionFailedSubscribers(NewDefaultActionFailedSubscriber(p.appIO))
}

type DefaultHookStartedSubscriber struct {
	AppIO io.IO
}

func NewDefaultHookStartedSubscriber(appIO io.IO) *DefaultHookStartedSubscriber {
	s := DefaultHookStartedSubscriber{
		AppIO: appIO,
	}
	return &s
}

func (s *DefaultHookStartedSubscriber) Handle(event *events.HookStarted) error {
	s.AppIO.Write("<comment>"+event.Config.Name()+":</comment>", true, io.NORMAL)
	if len(event.Config.GetActions()) == 0 {
		s.AppIO.Write(" - no actions to execute", true, io.NORMAL)
	}
	return nil
}

type DefaultHookSucceededSubscriber struct {
	AppIO io.IO
}

func NewDefaultHookSucceededSubscriber(appIO io.IO) *DefaultHookSucceededSubscriber {
	s := DefaultHookSucceededSubscriber{
		AppIO: appIO,
	}
	return &s
}

func (s *DefaultHookSucceededSubscriber) Handle(event *events.HookSucceeded) error {
	PrintActionLog(s.AppIO, event.Log)
	return nil
}

type DefaultHookFailedSubscriber struct {
	AppIO io.IO
}

func NewDefaultHookFailedSubscriber(appIO io.IO) *DefaultHookFailedSubscriber {
	s := DefaultHookFailedSubscriber{
		AppIO: appIO,
	}
	return &s
}

func (s *DefaultHookFailedSubscriber) Handle(event *events.HookFailed) error {
	PrintActionLog(s.AppIO, event.Log)
	s.AppIO.Write("\n<warning>captainhook failed: "+event.Error.Error()+"</warning>", true, io.NORMAL)
	return nil
}

type DefaultActionStartedSubscriber struct {
	AppIO io.IO
}

func NewDefaultActionStartedSubscriber(appIO io.IO) events.ActionStartedSubscriber {
	s := DefaultActionStartedSubscriber{
		AppIO: appIO,
	}
	return &s
}

func (s *DefaultActionStartedSubscriber) Handle(event *events.ActionStarted) error {
	actionLength := len(event.Config.Run())
	action := " - <info>"
	if actionLength > 60 {
		action = action + event.Config.Run()[0:60] + "..."
	} else {
		action = action + event.Config.Run() + strings.Repeat(" ", 60-actionLength)
	}
	action = action + "</info> : "
	s.AppIO.Write(action, false, io.NORMAL)
	return nil
}

type DefaultActionSucceededSubscriber struct {
	AppIO io.IO
}

func NewDefaultActionSucceededSubscriber(appIO io.IO) events.ActionSucceededSubscriber {
	s := DefaultActionSucceededSubscriber{
		AppIO: appIO,
	}
	return &s
}

func (s *DefaultActionSucceededSubscriber) Handle(event *events.ActionSucceeded) error {
	s.AppIO.Write("<ok>done</ok>", true, io.NORMAL)
	return nil
}

type DefaultActionSkippedSubscriber struct {
	AppIO io.IO
}

func NewDefaultActionSkippedSubscriber(appIO io.IO) events.ActionSkippedSubscriber {
	s := DefaultActionSkippedSubscriber{
		AppIO: appIO,
	}
	return &s
}

func (s *DefaultActionSkippedSubscriber) Handle(event *events.ActionSkipped) error {
	s.AppIO.Write("<comment>skipped</comment>", true, io.NORMAL)
	return nil
}

type DefaultActionFailedSubscriber struct {
	AppIO io.IO
}

func NewDefaultActionFailedSubscriber(appIO io.IO) events.ActionFailedSubscriber {
	s := DefaultActionFailedSubscriber{
		AppIO: appIO,
	}
	return &s
}

func (s *DefaultActionFailedSubscriber) Handle(event *events.ActionFailed) error {
	s.AppIO.Write("<warning>failed</warning>", true, io.NORMAL)
	return nil
}

func PrintActionLog(appIO io.IO, log *hooks.ActionLog) {
	if log.HasLogs() {
		for _, log := range log.Logs() {
			opening := "<ok>"
			closing := "</ok>"
			if log.Status == info.ACTION_FAILED {
				opening = "<warning>"
				closing = "</warning>"
			}
			if log.Status == info.ACTION_SKIPPED {
				opening = "<comment>"
				closing = "</comment>"
			}
			if log.CollectorIO.HasCollectedMessagesForVerbosity(appIO.Verbosity()) {
				appIO.Write("", true, io.NORMAL)
				appIO.Write(fmt.Sprintf("%sAction: "+log.Conf.Run()+"%s", opening, closing), true, io.NORMAL)
				for _, message := range log.CollectorIO.Messages() {
					appIO.Write(message.Message, false, message.Verbosity)
				}
			}
		}
	}
}