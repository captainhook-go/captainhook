package exec

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/io"
)

// ActionResult is the way to communicate the result of an action to the hook runner
// It contains the action output and errors
// It also knows if all events where dispatched correctly
// If you execute actions concurrently the channel is collection action results
type ActionResult struct {
	Config      *configuration.Action
	Status      int
	RunErr      error
	DispatchErr error
	Log         *io.CollectorIO
}

func NewActionResult(config *configuration.Action, status int, run, dispatch error, log *io.CollectorIO) *ActionResult {
	return &ActionResult{
		Config:      config,
		Status:      status,
		RunErr:      run,
		DispatchErr: dispatch,
		Log:         log,
	}
}
