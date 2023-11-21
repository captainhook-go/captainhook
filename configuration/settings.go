package configuration

type AppSettings struct {
	AllowFailure     bool
	AnsiColors       bool
	Custom           map[string]string
	FailOnFirstError bool
	GitDirectory     string
	Includes         []string
	IncludeLevel     int
	RunPath          string
	RunAsync         bool
	Verbosity        string
}

func NewDefaultAppSettings() *AppSettings {
	return &AppSettings{
		AllowFailure:     false,
		AnsiColors:       true,
		Custom:           map[string]string{},
		FailOnFirstError: true,
		GitDirectory:     ".git",
		Includes:         []string{},
		IncludeLevel:     1,
		RunAsync:         false,
		Verbosity:        "normal",
	}
}

type ActionSettings struct {
	AllowFailure bool
	WorkingDir   string
	Label        string
}

func NewDefaultActionSettings() *ActionSettings {
	return &ActionSettings{
		AllowFailure: false,
		WorkingDir:   "",
		Label:        "",
	}
}

func createActionSettingsFromJson(json *JsonActionSettings) *ActionSettings {
	a := NewDefaultActionSettings()
	if json == nil {
		return a
	}
	if json.AllowFailure != nil {
		a.AllowFailure = *json.AllowFailure
	}
	if json.WorkingDir != nil {
		a.WorkingDir = *json.WorkingDir
	}
	if json.Label != nil {
		a.Label = *json.Label
	}
	return a
}
