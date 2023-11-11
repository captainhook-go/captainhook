package configuration

type JsonConfiguration struct {
	Settings *JsonAppSettings      `json:"config,omitempty"`
	Hooks    *map[string]*JsonHook `json:"hooks,omitempty"`
}

type JsonHook struct {
	Actions []*JsonAction `load:"actions,omitempty"`
}

type JsonAction struct {
	Run        *string                 `json:"run,omitempty"`
	Conditions []*JsonCondition        `json:"conditions,omitempty"`
	Options    *map[string]interface{} `json:"options,omitempty"`
	Settings   *JsonActionSettings     `json:"config,omitempty"`
}

type JsonActionSettings struct {
	Label        *string `json:"label,omitempty"`
	AllowFailure *bool   `json:"allow-failure,omitempty"`
	RunAsync     *bool   `json:"run-async,omitempty"`
	WorkingDir   *string `json:"working-dir,omitempty"`
}

type JsonCondition struct {
	Run     *string                 `json:"run,omitempty"`
	Options *map[string]interface{} `json:"options,omitempty"`
}

type JsonAppSettings struct {
	AllowFailure     *bool              `json:"allow-failure,omitempty"`
	AnsiColors       *bool              `json:"ansi-colors,omitempty"`
	Custom           *map[string]string `json:"custom,omitempty"`
	FailOnFirstError *bool              `json:"fail-on-first-error,omitempty"`
	GitDirectory     *string            `json:"git-directory,omitempty"`
	Includes         *[]string          `json:"includes,omitempty"`
	IncludeLevel     *int               `json:"includes-level,omitempty"`
	RunPath          *string            `json:"run-path,omitempty"`
	Verbosity        *string            `json:"verbosity,omitempty"`
}
