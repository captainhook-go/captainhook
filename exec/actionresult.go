package exec

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/io"
)

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
