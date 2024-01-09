package printer

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/events"
	"github.com/captainhook-go/captainhook/hooks/app"
	"github.com/captainhook-go/captainhook/test"
	"testing"
)

func TestDefaultHookSuccess(t *testing.T) {
	io := test.CreateFakeIO()
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()
	ctx := app.NewContext(io, conf, repo)
	hook := &configuration.Hook{}
	p := NewDefaultPrinter(io)

	event := events.NewHookStartedEvent(ctx, hook)

	p.HookStarted(event)

	if len(io.Out) < 1 {
		t.Errorf("Should have written something")
	}
}
