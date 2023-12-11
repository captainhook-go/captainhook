package events

import (
	"github.com/captainhook-go/captainhook/test"
	"testing"
)

type HookStartedSubscriberMock struct {
	HandleWasCalled bool
}

func (m *HookStartedSubscriberMock) Handle(event *HookStarted) error {
	m.HandleWasCalled = true
	return nil
}

func TestHookStart(t *testing.T) {
	mock := &HookStartedSubscriberMock{false}

	dispatcher := NewDispatcher()
	dispatcher.RegisterHookStartedSubscribers(mock)

	inOut := test.CreateFakeIO()
	config := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()
	context := test.CreateFakeHookContext(inOut, config, repo)

	event := NewHookStartedEvent(context, config.HookConfig("pre-commit"))
	err := dispatcher.DispatchHookStartedEvent(event)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if !mock.HandleWasCalled {
		t.Errorf("Subscriber should have been executed")
	}
}
