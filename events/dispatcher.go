package events

type Dispatcher struct {
	hookStartedHandlers     []HookStartedSubscriber
	hookSucceededHandlers   []HookSucceededSubscriber
	hookFailedHandlers      []HookFailedSubscriber
	actionStartedHandlers   []ActionStartedSubscriber
	actionSkippedHandlers   []ActionSkippedSubscriber
	actionFailedHandlers    []ActionFailedSubscriber
	actionSucceededHandlers []ActionSucceededSubscriber
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) RegisterHookStartedSubscribers(handlers ...HookStartedSubscriber) {
	for _, handler := range handlers {
		d.hookStartedHandlers = append(d.hookStartedHandlers, handler)
	}
}

func (d *Dispatcher) DispatchHookStartedEvent(event *HookStarted) error {
	for _, handler := range d.hookStartedHandlers {
		err := handler.Handle(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dispatcher) RegisterHookFailedSubscribers(handlers ...HookFailedSubscriber) {
	for _, handler := range handlers {
		d.hookFailedHandlers = append(d.hookFailedHandlers, handler)
	}
}

func (d *Dispatcher) DispatchHookFailedEvent(event *HookFailed) error {
	for _, handler := range d.hookFailedHandlers {
		err := handler.Handle(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dispatcher) RegisterHookSucceededSubscribers(handlers ...HookSucceededSubscriber) {
	for _, handler := range handlers {
		d.hookSucceededHandlers = append(d.hookSucceededHandlers, handler)
	}
}

func (d *Dispatcher) DispatchHookSucceededEvent(event *HookSucceeded) error {
	for _, handler := range d.hookSucceededHandlers {
		err := handler.Handle(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dispatcher) RegisterActionStartedSubscribers(handlers ...ActionStartedSubscriber) {
	for _, handler := range handlers {
		d.actionStartedHandlers = append(d.actionStartedHandlers, handler)
	}
}

func (d *Dispatcher) DispatchActionStartedEvent(event *ActionStarted) error {
	for _, handler := range d.actionStartedHandlers {
		err := handler.Handle(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dispatcher) RegisterActionSkippedSubscribers(handlers ...ActionSkippedSubscriber) {
	for _, handler := range handlers {
		d.actionSkippedHandlers = append(d.actionSkippedHandlers, handler)
	}
}

func (d *Dispatcher) DispatchActionSkippedEvent(event *ActionSkipped) error {
	for _, handler := range d.actionSkippedHandlers {
		err := handler.Handle(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dispatcher) RegisterActionFailedSubscribers(handlers ...ActionFailedSubscriber) {
	for _, handler := range handlers {
		d.actionFailedHandlers = append(d.actionFailedHandlers, handler)
	}
}

func (d *Dispatcher) DispatchActionFailedEvent(event *ActionFailed) error {
	for _, handler := range d.actionFailedHandlers {
		err := handler.Handle(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dispatcher) RegisterActionSucceededSubscribers(handlers ...ActionSucceededSubscriber) {
	for _, handler := range handlers {
		d.actionSucceededHandlers = append(d.actionSucceededHandlers, handler)
	}
}

func (d *Dispatcher) DispatchActionSucceededEvent(event *ActionSucceeded) error {
	for _, handler := range d.actionSucceededHandlers {
		err := handler.Handle(event)
		if err != nil {
			return err
		}
	}
	return nil
}
