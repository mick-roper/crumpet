package core

import "errors"

// ScenarioSpec that a test runner will process
type ScenarioSpec struct {
	Host string `json:"host"`
	Scenarios []*Scenario `json:"scenarios"`
}

// Scenario that the runner will process
type Scenario struct {
	Name string `json:"name"`
	Iterations int `json:"iterations"`
	Concurrency int `json:"concurrency"`
	MinDelayMs int `json:"minDelayMs"`
	MaxDelayMs int `json:"maxDelayMs"`
	Options *ScenarioSpecOptions `json:"options"`
	Steps []*ScenarioStep `json:"steps"`
}

// ScenarioStep that is processed in order
type ScenarioStep struct {
	Name string `json:"name"`
	Method string `json:"method"`
	Path string `json:"path"`
	Data interface{} `json:"data"`
}

// Validate a ScenarioSpec
func (t *ScenarioSpec) Validate() error {
	for _, s := range t.Scenarios {
		if s.MaxDelayMs < s.MinDelayMs {
			return errors.New("max delay must greater than or equal to min delay")
		}
	}

	// todo: make this better!
	return nil
}

// ScenarioSpecOptions that help to further describe a test spec
type ScenarioSpecOptions struct {
	HTTPTimeout int `json:"httpTimeout"`
	HTTPRequestHeaders map[string]string `json:"httpRequestHeaders"`
}