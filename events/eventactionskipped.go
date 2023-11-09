package events

import (
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/config"
)

type ActionSkipped struct {
	Context *app.Context
	Config  *config.Action
}

func NewActionSkippedEvent(context *app.Context, action *config.Action) *ActionSkipped {
	e := ActionSkipped{
		Context: context,
		Config:  action,
	}
	return &e
}

type ActionSkippedSubscriber interface {
	Handle(event *ActionSkipped) error
}
