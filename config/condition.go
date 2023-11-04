package config

type Condition struct {
	exec string
	args []string
}

func (c *Condition) Exec() string {
	return c.exec
}

func (c *Condition) Args() []string {
	return c.args
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
	var args []string

	if json.Args != nil {
		args = *json.Args
	}

	c := Condition{exec: *json.Exec, args: args}

	return &c
}
