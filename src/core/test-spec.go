package core

// TestSpec that a test runner will process
type TestSpec struct {
	URL string
	Iterations int
	Concurrency int
	MaxDelay int
}

// TestResponse produced by a single test run
type TestResponse struct {
	StatusCode int
	ElapsedMs float64
}

type TestResult struct {
	AverageElapsed float64
	MaxElapsedMs float64
	MinElapsedMs float64
	StandardDeviation float64
	Responses []*TestResponse
}