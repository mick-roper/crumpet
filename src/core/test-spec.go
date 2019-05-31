package core

// TestSpec that a test runner will process
type TestSpec struct {
	Host string `json:"host"`
	Paths []string `json:"paths"`
	Iterations int `json:"iterations"`
	Concurrency int `json:"concurrency"`
	MaxDelayMs int `json:"maxDelayMs"`
	Options *TestSpecOptions `json:"options"`
}

// TestSpecOptions that help to further describe a test spec
type TestSpecOptions struct {
	HTTPRequestHeaders map[string]string `json:"httpReqHeaders"`
}

// TestResponse produced by a single test run
type TestResponse struct {
	URL string
	StatusCode int
	ElapsedMs float64
	Data string
}

// TestResult from an executed spec
type TestResult struct {
	AverageElapsed float64
	MaxElapsedMs float64
	MinElapsedMs float64
	StandardDeviation float64
	Responses []*TestResponse
}