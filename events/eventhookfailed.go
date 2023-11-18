package events

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/hooks/app"
	"time"
)

type HookFailed struct {
	Context  *app.Context
	Config   *configuration.Hook
	Log      *hooks.ActionLog
	ExecTime time.Duration
	Error    error
}

func NewHookFailedEvent(
	context *app.Context,
	hook *configuration.Hook,
	log *hooks.ActionLog,
	execTime time.Duration,
	err error,
) *HookFailed {
	e := HookFailed{
		Context:  context,
		Config:   hook,
		Log:      log,
		ExecTime: execTime,
		Error:    err,
	}
	return &e
}

type HookFailedSubscriber interface {
	Handle(event *HookFailed) error
}
