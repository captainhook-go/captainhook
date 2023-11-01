package events

type Subscriber interface {
	setEventHandlers(dispatcher Dispatcher)
}
