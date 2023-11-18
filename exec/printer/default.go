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

func (p *DefaultPrinter) HookStarted(event *events.HookStarted) {
	p.appIO.Write("<comment>"+event.Config.Name()+":</comment>", true, io.NORMAL)
	if len(event.Config.GetActions()) == 0 {
		p.appIO.Write(" - no actions to execute", true, io.NORMAL)
	}
}

func (p *DefaultPrinter) HookSucceeded(event *events.HookSucceeded) {
	p.printActionLog(event.Log)
	p.appIO.Write("<ok>captainhook successfully executed all actions</ok>", true, io.NORMAL)
}

func (p *DefaultPrinter) HookFailed(event *events.HookFailed) {
	p.printActionLog(event.Log)
	p.appIO.Write("<warning>captainhook failed: "+event.Error.Error()+"</warning>", true, io.NORMAL)
}

func (p *DefaultPrinter) ActionSuccess(event *events.ActionSucceeded) {
	p.printActionIntro(event.Config.Label())
	p.appIO.Write("<ok>done</ok>", true, io.NORMAL)
}

func (p *DefaultPrinter) ActionSkipped(event *events.ActionSkipped) {
	p.printActionIntro(event.Config.Label())
	p.appIO.Write("<comment>skipped</comment>", true, io.NORMAL)
}

func (p *DefaultPrinter) ActionFailed(event *events.ActionFailed) {
	p.printActionIntro(event.Config.Label())
	p.appIO.Write("<warning>failed</warning>", true, io.NORMAL)
}

func (p *DefaultPrinter) RegisterSubscribers(dispatcher *events.Dispatcher) {
	dispatcher.RegisterHookStartedSubscribers(NewDefaultHookStartedSubscriber(p))
	dispatcher.RegisterHookSucceededSubscribers(NewDefaultHookSucceededSubscriber(p))
	dispatcher.RegisterHookFailedSubscribers(NewDefaultHookFailedSubscriber(p))
	dispatcher.RegisterActionSucceededSubscribers(NewDefaultActionSucceededSubscriber(p))
	dispatcher.RegisterActionSkippedSubscribers(NewDefaultActionSkippedSubscriber(p))
	dispatcher.RegisterActionFailedSubscribers(NewDefaultActionFailedSubscriber(p))
}

func (p *DefaultPrinter) printActionIntro(label string) {
	actionLength := len(label)
	action := " - <info>"
	if actionLength > 58 {
		action = action + label[0:58] + "..."
	} else {
		action = action + label + strings.Repeat(" ", 61-actionLength)
	}
	action = action + "</info> : "
	p.appIO.Write(action, false, io.NORMAL)
}

func (p *DefaultPrinter) printActionLog(log *hooks.ActionLog) {
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
			if log.CollectorIO.HasCollectedMessagesForVerbosity(p.appIO.Verbosity()) {
				p.appIO.Write("", true, io.NORMAL)
				p.appIO.Write(fmt.Sprintf("%sAction: "+log.Conf.Run()+"%s", opening, closing), true, io.NORMAL)
				for _, message := range log.CollectorIO.Messages() {
					p.appIO.Write(message.Message, false, message.Verbosity)
				}
			}
		}
	}
}
