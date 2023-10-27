package json

type ActionJson struct {
	Action     *string
	Settings   *ActionConfigJson `json:"config,omitempty"`
	Conditions *[]ConditionJson  `json:"conditions,omitempty"`
	Options    *OptionsJson      `json:"options ons,omitempty"`
}
