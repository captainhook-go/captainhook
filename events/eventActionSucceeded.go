package events

import (
	"github.com/captainhook-go/captainhook/app"
	"github.com/captainhook-go/captainhook/config"
)

type ActionSucceeded struct {
	Context *app.Context
	Config  *config.Action
}

func NewActionSucceededEvent(context *app.Context, action *config.Action) *ActionSucceeded {
	e := ActionSucceeded{
		Context: context,
		Config:  action,
	}
	return &e
}

type ActionSucceededSubscriber interface {
	Handle(event *ActionSucceeded) error
}
