package config

type Action struct {
	action     string
	settings   *ActionSettings
	conditions []Condition
	options    Options
}

func CreateActionFromJson(json *JsonAction) *Action {
	return &Action{
		action:   *json.Action,
		settings: createActionSettingsFromJson(json.Settings),
	}
}

func (a *Action) Action() string {
	return a.action
}

func (a *Action) Options() Options {
	return a.options
}

func (a *Action) Label() string {
	if len(a.settings.Label) > 0 {
		return a.settings.Label
	}
	return a.Action()
}

func (a *Action) IsFailureAllowed() bool {
	return a.settings.AllowFailure
}

func (a *Action) WorkingDir() string {
	return a.settings.WorkingDir
}
