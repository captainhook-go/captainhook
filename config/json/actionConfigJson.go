package json

type ActionConfigJson struct {
	AllowFailure *bool `json:"allow-failure,omitempty"`
	RunAsync     *bool `json:"run-async,omitempty"`
	WorkingDir   *bool `json:"working-dir,omitempty"`
}
