package message

import "github.com/captainhook-go/captainhook/git/types"

type Rulebook struct {
	rules []Rule
}

func (r *Rulebook) AddRule(rules ...Rule) {
	for _, rule := range rules {
		r.rules = append(r.rules, rule)
	}
}

func (r *Rulebook) IsFollowedBy(message *types.CommitMessage) (bool, []string) {
	var problems []string
	for _, rule := range r.rules {
		ok, hint := rule.IsFollowedBy(message)
		if !ok {
			problems = append(problems, hint)
		}
	}
	return len(problems) == 0, problems
}

func NewRulebook() *Rulebook {
	return &Rulebook{}
}
