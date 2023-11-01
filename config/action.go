package config

import (
	"strconv"
)

type Action struct {
	action     string
	settings   map[string]string
	conditions []Condition
	options    Options
}

func CreateActionFromJson(json JsonAction) *Action {
	return &Action{
		action: *json.Action,
	}
}

func (a Action) IsFailureAllowed() bool {
	return a.getSettingBool(SETTING_ALLOW_FAILURE, false)
}

func (a Action) getSettingBool(name string, defaultValue bool) bool {
	if a.settings == nil {
		return defaultValue
	}
	var valString, existed = a.settings[name]
	if existed {
		valBool, err := strconv.ParseBool(valString)
		if err == nil {
			return valBool
		}
	}
	return defaultValue
}

func (a Action) Action() string {
	return a.action
}

func (a Action) Options() Options {
	return a.options
}
