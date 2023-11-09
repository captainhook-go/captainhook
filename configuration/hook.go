package configuration

type Hook struct {
	name      string
	isEnabled bool
	actions   []*Action
}

func NewHook(name string, isEnabled bool) *Hook {
	return &Hook{name: name, isEnabled: isEnabled}
}

func (h *Hook) Name() string {
	return h.name
}

func (h *Hook) IsEnabled() bool {
	return h.isEnabled
}
func (h *Hook) AddAction(action *Action) {
	h.actions = append(h.actions, action)
}
func (h *Hook) GetActions() []*Action {
	return h.actions
}
