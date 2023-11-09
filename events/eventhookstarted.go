package events

import (
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/configuration"
)

type HookStarted struct {
	Context *app.Context
	Config  *configuration.Hook
}

func NewHookStartedEvent(context *app.Context, hook *configuration.Hook) *HookStarted {
	e := HookStarted{
		Context: context,
		Config:  hook,
	}
	return &e
}

type HookStartedSubscriber interface {
	Handle(event *HookStarted) error
}
