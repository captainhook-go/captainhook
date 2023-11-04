package config

type JsonConfiguration struct {
	Settings *JsonAppSettings      `json:"config,omitempty"`
	Hooks    *map[string]*JsonHook `json:"hooks,omitempty"`
}

type JsonHook struct {
	Actions []*JsonAction `load:"actions,omitempty"`
}

type JsonAction struct {
	Action     *string
	Conditions []*JsonCondition    `json:"conditions,omitempty"`
	Options    *JsonOptions        `json:"options,omitempty"`
	Settings   *JsonActionSettings `json:"config,omitempty"`
}

type JsonActionSettings struct {
	Label        *string `json:"label,omitempty"`
	AllowFailure *bool   `json:"allow-failure,omitempty"`
	RunAsync     *bool   `json:"run-async,omitempty"`
	WorkingDir   *string `json:"working-dir,omitempty"`
}

type JsonCondition struct {
	Exec *string
	Args *[]string
}

type JsonOptions struct {
	options map[string]interface{}
}

func (o JsonOptions) valueOf(option string, defaultValue string) interface{} {
	var value, ok = o.options[option]
	if ok {
		return value
	}
	return defaultValue
}

func (o JsonOptions) all() interface{} {
	return o.options
}

type JsonAppSettings struct {
	AllowFailure     *bool              `json:"allow-failure,omitempty"`
	AnsiColors       *bool              `json:"ansi-colors,omitempty"`
	Custom           *map[string]string `json:"custom,omitempty"`
	FailOnFirstError *bool              `json:"fail-on-first-error,omitempty"`
	GitDirectory     *string            `json:"git-directory,omitempty"`
	Includes         *[]string          `json:"includes,omitempty"`
	IncludeLevel     *int               `json:"includes-level,omitempty"`
	Verbosity        *string            `json:"verbosity,omitempty"`
}
