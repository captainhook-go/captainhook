package events

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/hooks/app"
)

type ActionFailed struct {
	Context *app.Context
	Config  *configuration.Action
	Error   error
}

func NewActionFailedEvent(context *app.Context, action *configuration.Action, err error) *ActionFailed {
	e := ActionFailed{
		Context: context,
		Config:  action,
		Error:   err,
	}
	return &e
}

type ActionFailedSubscriber interface {
	Handle(event *ActionFailed) error
}
