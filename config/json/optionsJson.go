package json

type OptionsJson struct {
	options map[string]interface{}
}

func (o OptionsJson) valueOf(option string, defaultValue string) interface{} {
	var value, ok = o.options[option]
	if ok {
		return value
	}
	return defaultValue
}

func (o OptionsJson) all() interface{} {
	return o.options
}
