package events

import (
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/config"
)

type ActionFailed struct {
	Context *app.Context
	Config  *config.Action
	Error   error
}

func NewActionFailedEvent(context *app.Context, action *config.Action, err error) *ActionFailed {
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
