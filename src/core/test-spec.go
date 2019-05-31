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