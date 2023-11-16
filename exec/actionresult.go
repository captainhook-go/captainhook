package exec

const (
	ActionSuccessful = 0
	ActionSkipped    = 1
	ActionFailed     = 2
)

type ActionResult struct {
	State       int
	RunErr      error
	DispatchErr error
}

func NewActionResult(state int, run, dispatch error) *ActionResult {
	return &ActionResult{state, run, dispatch}
}
