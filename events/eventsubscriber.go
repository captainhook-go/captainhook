package events

import "github.com/captainhook-go/captainhook/configuration"

type EventSubscriber interface {
	Subscribe(dispatcher *Dispatcher, action *configuration.Action) error
}
