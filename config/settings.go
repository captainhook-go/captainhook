package config

import "github.com/captainhook-go/captainhook/config/json"

type Settings struct {
	AllowFailure     bool
	AnsiColors       bool
	Custom           map[string]string
	FailOnFirstError bool
	GitDirectory     string
	Includes         []string
	IncludeLevel     int
	Verbosity        string
}

func createSettingsFromJson(json json.SettingsJson) *Settings {
	return &Settings{
		AllowFailure:     *json.AllowFailure,
		AnsiColors:       *json.AnsiColors,
		Custom:           *json.Custom,
		FailOnFirstError: *json.FailOnFirstError,
		GitDirectory:     *json.GitDirectory,
		Includes:         *json.Includes,
		IncludeLevel:     *json.IncludeLevel,
		Verbosity:        *json.Verbosity,
	}
}
