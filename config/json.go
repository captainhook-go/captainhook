package config

type JsonAction struct {
	Action     *string
	Settings   *JsonActionConfig `load:"config,omitempty"`
	Conditions *[]JsonCondition  `load:"conditions,omitempty"`
	Options    *JsonOptions      `load:"options ons,omitempty"`
}

type JsonActionConfig struct {
	AllowFailure *bool `load:"allow-failure,omitempty"`
	RunAsync     *bool `load:"run-async,omitempty"`
	WorkingDir   *bool `load:"working-dir,omitempty"`
}

type JsonCondition struct {
	Exec *string
	Args *[]string
}

type JsonConfiguration struct {
	Settings JsonSettings         `load:"config,omitempty"`
	Hooks    map[string]*JsonHook `load:"hooks,omitempty"`
}

type JsonHook struct {
	Actions []*JsonAction `load:"actions,omitempty"`
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

type JsonSettings struct {
	AllowFailure     *bool              `load:"allow-failure,omitempty"`
	AnsiColors       *bool              `load:"ansi-colors,omitempty"`
	Custom           *map[string]string `load:"custom,omitempty"`
	FailOnFirstError *bool              `load:"fail-on-first-error,omitempty"`
	GitDirectory     *string            `load:"git-directory,omitempty"`
	Includes         *[]string          `load:"includes,omitempty"`
	IncludeLevel     *int               `load:"includes-level,omitempty"`
	Verbosity        *string            `load:"verbosity,omitempty"`
}
