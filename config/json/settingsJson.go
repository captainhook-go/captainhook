package json

type SettingsJson struct {
	AllowFailure     *bool              `json:"allow-failure,omitempty"`
	AnsiColors       *bool              `json:"ansi-colors,omitempty"`
	Custom           *map[string]string `json:"custom,omitempty"`
	FailOnFirstError *bool              `json:"fail-on-first-error,omitempty"`
	GitDirectory     *string            `json:"git-directory,omitempty"`
	Includes         *[]string          `json:"includes,omitempty"`
	IncludeLevel     *int               `json:"includes-level,omitempty"`
	Verbosity        *string            `json:"verbosity,omitempty"`
}
