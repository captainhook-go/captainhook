package events

import (
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/config"
)

type HookStarted struct {
	Context *app.Context
	Config  *config.Hook
}

func NewHookStartedEvent(context *app.Context, hook *config.Hook) *HookStarted {
	e := HookStarted{
		Context: context,
		Config:  hook,
	}
	return &e
}

type HookStartedSubscriber interface {
	Handle(event *HookStarted) error
}
