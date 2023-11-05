package config

import (
	"fmt"
	"strconv"
)

type Options struct {
	options map[string]interface{}
}

func (o Options) IntValueOf(option string, defaultValue int) int {
	var value, ok = o.options[option]
	if ok {
		switch value := value.(type) {
		case int:
			return value
		case int64:
		case float64:
			return int(value)
		case string:
			i, err := strconv.Atoi(value)
			if err != nil {
				return defaultValue
			}
			return i
		default:
			return defaultValue
		}
	}
	return defaultValue
}

func (o Options) StringValueOf(option string, defaultValue string) string {
	var value, ok = o.options[option]
	if ok {
		switch value := value.(type) {
		case int:
		case int64:
		case float64:
			return fmt.Sprintf("%d", value)
		case string:
			return value
		default:
			return defaultValue
		}
	}
	return defaultValue
}

func (o Options) All() map[string]interface{} {
	return o.options
}

func createOptionsFromJson(jsonOptions *map[string]interface{}) *Options {
	options := map[string]interface{}{}

	if jsonOptions != nil {
		options = *jsonOptions
	}
	o := Options{options}
	return &o
}
