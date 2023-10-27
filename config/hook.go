package config

type Hook struct {
	isEnabled bool
	actions   []*Action
}

func NewHook(isEnabled bool) *Hook {
	return &Hook{isEnabled: isEnabled}
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
