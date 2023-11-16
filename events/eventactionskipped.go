package events

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/hooks/app"
)

type ActionSkipped struct {
	Context *app.Context
	Config  *configuration.Action
}

func NewActionSkippedEvent(context *app.Context, action *configuration.Action) *ActionSkipped {
	e := ActionSkipped{
		Context: context,
		Config:  action,
	}
	return &e
}

type ActionSkippedSubscriber interface {
	Handle(event *ActionSkipped) error
}
