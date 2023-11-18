package events

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/app"
	"time"
)

type HookSucceeded struct {
	Context  *app.Context
	Config   *configuration.Hook
	Log      *hooks.ActionLog
	ExecTime time.Duration
}

func NewHookSucceededEvent(
	context *app.Context,
	hook *configuration.Hook,
	log *hooks.ActionLog,
	execTime time.Duration,
) *HookSucceeded {
	e := HookSucceeded{
		Context:  context,
		Config:   hook,
		Log:      log,
		ExecTime: execTime,
	}
	return &e
}

type HookSucceededSubscriber interface {
	Handle(event *HookSucceeded) error
}
