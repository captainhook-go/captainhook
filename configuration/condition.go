package configuration

type Condition struct {
	run     string
	options *Options
}

func (c *Condition) Run() string {
	return c.run
}

func (c *Condition) Options() *Options {
	return c.options
}

func createConditionsFromJson(jsonConditions []*JsonCondition) []*Condition {
	var conditions []*Condition
	if jsonConditions == nil {
		return conditions
	}
	for _, condition := range jsonConditions {
		conditions = append(conditions, createConditionFromJson(condition))
	}
	return conditions
}

func createConditionFromJson(json *JsonCondition) *Condition {
	var o *Options

	if json.Options != nil {
		o = createOptionsFromJson(json.Options)
	}

	c := Condition{run: *json.Run, options: o}

	return &c
}
