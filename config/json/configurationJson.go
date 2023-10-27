package json

type ConfigurationJson struct {
	Settings SettingsJson         `json:"config,omitempty"`
	Hooks    map[string]*HookJson `json:"hooks,omitempty"`
}
