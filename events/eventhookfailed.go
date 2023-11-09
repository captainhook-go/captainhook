package events

import (
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/hooks"
)

type HookFailed struct {
	Context *app.Context
	Config  *configuration.Hook
	Log     *hooks.ActionLog
	Error   error
}

func NewHookFailedEvent(context *app.Context, hook *configuration.Hook, log *hooks.ActionLog, err error) *HookFailed {
	e := HookFailed{
		Context: context,
		Config:  hook,
		Log:     log,
		Error:   err,
	}
	return &e
}

type HookFailedSubscriber interface {
	Handle(event *HookFailed) error
}
