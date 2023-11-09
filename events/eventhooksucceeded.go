package events

import (
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/hooks"
)

type HookSucceeded struct {
	Context *app.Context
	Config  *config.Hook
	Log     *hooks.ActionLog
}

func NewHookSucceededEvent(context *app.Context, hook *config.Hook, log *hooks.ActionLog) *HookSucceeded {
	e := HookSucceeded{
		Context: context,
		Config:  hook,
		Log:     log,
	}
	return &e
}

type HookSucceededSubscriber interface {
	Handle(event *HookSucceeded) error
}
