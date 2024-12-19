package configuration

type Condition struct {
	run        string
	options    *Options
	conditions []*Condition
}

func (c *Condition) Run() string {
	return c.run
}

func (c *Condition) Options() *Options {
	return c.options
}

func (c *Condition) Conditions() []*Condition {
	return c.conditions
}

func NewCondition(cmd string, o *Options, c []*Condition) *Condition {
	return &Condition{run: cmd, options: o, conditions: c}
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
	var c []*Condition

	if json.Options != nil {
		o = createOptionsFromJson(json.Options)
	}
	if json.Conditions != nil {
		c = createConditionsFromJson(json.Conditions)
	}
	return NewCondition(json.Run, o, c)
}
