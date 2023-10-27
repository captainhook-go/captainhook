package config

type Condition struct {
	Exec string
	Args []string
}

func NewCondition(exec string, args []string) *Condition {
	return &Condition{Exec: exec, Args: args}
}
