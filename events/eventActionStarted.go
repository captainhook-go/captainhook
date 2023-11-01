package events

import (
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/config"
)

type ActionStarted struct {
	Context *app.Context
	Config  *config.Action
}

func NewActionStartedEvent(context *app.Context, action *config.Action) *ActionStarted {
	e := ActionStarted{
		Context: context,
		Config:  action,
	}
	return &e
}

type ActionStartedSubscriber interface {
	Handle(event *ActionStarted) error
}
