package events

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/hooks/app"
)

type ActionSucceeded struct {
	Context *app.Context
	Config  *configuration.Action
}

func NewActionSucceededEvent(context *app.Context, action *configuration.Action) *ActionSucceeded {
	e := ActionSucceeded{
		Context: context,
		Config:  action,
	}
	return &e
}

type ActionSucceededSubscriber interface {
	Handle(event *ActionSucceeded) error
}
