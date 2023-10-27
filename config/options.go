package config

type Options struct {
	options map[string]interface{}
}

func (o Options) valueOf(option string, defaultValue string) interface{} {
	var value, ok = o.options[option]
	if ok {
		return value
	}
	return defaultValue
}

func (o Options) all() interface{} {
	return o.options
}
