package printer

import (
	"github.com/captainhook-go/captainhook/events"
	"github.com/captainhook-go/captainhook/io"
)

type DefaultHookStartedSubscriber struct {
	printer *DefaultPrinter
}

func NewDefaultHookStartedSubscriber(printer *DefaultPrinter) *DefaultHookStartedSubscriber {
	s := DefaultHookStartedSubscriber{
		printer: printer,
	}
	return &s
}

func (s *DefaultHookStartedSubscriber) Handle(event *events.HookStarted) error {
	s.printer.HookStarted(event)
	return nil
}

type DefaultHookSucceededSubscriber struct {
	printer *DefaultPrinter
}

func NewDefaultHookSucceededSubscriber(printer *DefaultPrinter) *DefaultHookSucceededSubscriber {
	s := DefaultHookSucceededSubscriber{
		printer: printer,
	}
	return &s
}

func (s *DefaultHookSucceededSubscriber) Handle(event *events.HookSucceeded) error {
	s.printer.HookSucceeded(event)
	return nil
}

type DefaultHookFailedSubscriber struct {
	printer *DefaultPrinter
}

func NewDefaultHookFailedSubscriber(printer *DefaultPrinter) *DefaultHookFailedSubscriber {
	s := DefaultHookFailedSubscriber{
		printer: printer,
	}
	return &s
}

func (s *DefaultHookFailedSubscriber) Handle(event *events.HookFailed) error {
	s.printer.HookFailed(event)
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
	return nil
}

type DefaultActionSucceededSubscriber struct {
	printer *DefaultPrinter
}

func NewDefaultActionSucceededSubscriber(printer *DefaultPrinter) events.ActionSucceededSubscriber {
	s := DefaultActionSucceededSubscriber{
		printer: printer,
	}
	return &s
}

func (s *DefaultActionSucceededSubscriber) Handle(event *events.ActionSucceeded) error {
	s.printer.ActionSuccess(event)
	return nil
}

type DefaultActionSkippedSubscriber struct {
	printer *DefaultPrinter
}

func NewDefaultActionSkippedSubscriber(printer *DefaultPrinter) events.ActionSkippedSubscriber {
	s := DefaultActionSkippedSubscriber{
		printer: printer,
	}
	return &s
}

func (s *DefaultActionSkippedSubscriber) Handle(event *events.ActionSkipped) error {
	s.printer.ActionSkipped(event)
	return nil
}

type DefaultActionFailedSubscriber struct {
	printer *DefaultPrinter
}

func NewDefaultActionFailedSubscriber(printer *DefaultPrinter) events.ActionFailedSubscriber {
	s := DefaultActionFailedSubscriber{
		printer: printer,
	}
	return &s
}

func (s *DefaultActionFailedSubscriber) Handle(event *events.ActionFailed) error {
	s.printer.ActionFailed(event)
	return nil
}
