package events

import (
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/configuration"
)

type ActionStarted struct {
	Context *app.Context
	Config  *configuration.Action
}

func NewActionStartedEvent(context *app.Context, action *configuration.Action) *ActionStarted {
	e := ActionStarted{
		Context: context,
		Config:  action,
	}
	return &e
}

type ActionStartedSubscriber interface {
	Handle(event *ActionStarted) error
}
