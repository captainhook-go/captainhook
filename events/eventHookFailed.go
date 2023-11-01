package events

import (
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/config"
)

type HookFailed struct {
	Context *app.Context
	Config  *config.Hook
	Error   error
}

func NewHookFailedEvent(context *app.Context, hook *config.Hook, err error) *HookFailed {
	e := HookFailed{
		Context: context,
		Config:  hook,
		Error:   err,
	}
	return &e
}

type HookFailedSubscriber interface {
	Handle(event *HookFailed) error
}
