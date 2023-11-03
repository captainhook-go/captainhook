package config

type AppSettings struct {
	AllowFailure     bool
	AnsiColors       bool
	Custom           map[string]string
	FailOnFirstError bool
	GitDirectory     string
	Includes         []string
	IncludeLevel     int
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
		IncludeLevel:     0,
		Verbosity:        "normal",
	}
}

func createAppSettingsFromJson(json *JsonAppSettings) *AppSettings {
	s := NewDefaultAppSettings()
	if json == nil {
		return s
	}

	if json.AllowFailure != nil {
		s.AllowFailure = *json.AllowFailure
	}
	if json.AnsiColors != nil {
		s.AnsiColors = *json.AnsiColors
	}
	if (json.Custom) != nil {
		s.Custom = *json.Custom
	}
	if json.FailOnFirstError != nil {
		s.FailOnFirstError = *json.FailOnFirstError
	}
	if json.GitDirectory != nil {
		s.GitDirectory = *json.GitDirectory
	}
	if json.Verbosity != nil {
		s.Verbosity = *json.Verbosity
	}
	return s
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
