package printer

import (
	"fmt"
	"github.com/captainhook-go/captainhook/events"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"strings"
	"time"
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
}

func (p *DefaultPrinter) HookSucceeded(event *events.HookSucceeded) {
	p.printActionLog(event.Log)
	d := event.ExecTime.Round(2 * time.Millisecond)
	p.appIO.Write(fmt.Sprintf("<ok>captainhook successfully executed all actions in %s</ok>", d.String()), true, io.NORMAL)
}

func (p *DefaultPrinter) HookFailed(event *events.HookFailed) {
	p.printActionLog(event.Log)
	p.appIO.Write("<warning>captainhook failed: "+event.Error.Error()+"</warning>", true, io.NORMAL)
}

func (p *DefaultPrinter) ActionSuccess(event *events.ActionSucceeded) {
	p.appIO.Write(p.actionIntro(event.Config.Label())+"<ok>done</ok>", true, io.NORMAL)
}

func (p *DefaultPrinter) ActionSkipped(event *events.ActionSkipped) {
	p.appIO.Write(p.actionIntro(event.Config.Label())+"<comment>skipped</comment>", true, io.NORMAL)
}

func (p *DefaultPrinter) ActionFailed(event *events.ActionFailed) {
	p.appIO.Write(p.actionIntro(event.Config.Label())+"<warning>failed</warning>", true, io.NORMAL)
}

func (p *DefaultPrinter) RegisterSubscribers(dispatcher *events.Dispatcher) {
	dispatcher.RegisterHookStartedSubscribers(NewDefaultHookStartedSubscriber(p))
	dispatcher.RegisterHookSucceededSubscribers(NewDefaultHookSucceededSubscriber(p))
	dispatcher.RegisterHookFailedSubscribers(NewDefaultHookFailedSubscriber(p))
	dispatcher.RegisterActionSucceededSubscribers(NewDefaultActionSucceededSubscriber(p))
	dispatcher.RegisterActionSkippedSubscribers(NewDefaultActionSkippedSubscriber(p))
	dispatcher.RegisterActionFailedSubscribers(NewDefaultActionFailedSubscriber(p))
}

func (p *DefaultPrinter) actionIntro(label string) string {
	actionLength := len(label)
	action := " - <info>"
	if actionLength > 58 {
		action = action + label[0:58] + "..."
	} else {
		action = action + label + strings.Repeat(" ", 61-actionLength)
	}
	return action + "</info> : "
}

func (p *DefaultPrinter) printActionLog(log *hooks.ActionLog) {
	if log.HasLogs() {
		for _, log := range log.Logs() {
			opening := "<ok>"
			closing := "</ok>"
			if log.Status == info.ActionFailed {
				opening = "<warning>"
				closing = "</warning>"
			}
			if log.Status == info.ActionSkipped {
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
