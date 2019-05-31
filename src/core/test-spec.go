package core

import (
	"fmt"
)

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
	HTTPRequestHeaders map[string]string `json:"httpRequestHeaders"`
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
	requestCount uint64
	responses []float64
}

// Print the result
func (t *TestResult) Print() {
	fmt.Printf("requests:\t%v\n", t.RequestCount())
	fmt.Printf("min response time:\t%vms\n", t.MinElapsedMs())
	fmt.Printf("max response time:\t%vms\n", t.MaxElapsedMs())
	fmt.Printf("avg response time:\t%6.2fms\n", t.AverageElapsedMs())
	fmt.Printf("std deviation:\t%6.2fms\n", t.StandardDeviation())
	fmt.Printf("90th percentile:\t%6.2fms\n", t.AverageElapsedMs90thPc())
	fmt.Printf("95th percentile:\t%6.2fms\n", t.AverageElapsedMs95thPc())
	fmt.Printf("99th percentile:\t%6.2fms\n", t.AverageElapsedMs99thPc())
}

// RequestCount returns the number of requests
func (t *TestResult) RequestCount() uint64 {
	return t.requestCount
}

// AverageElapsedMs returns the overall average elapsed MS
func (t *TestResult) AverageElapsedMs() float64 {
	return getAvg(t.responses)
}

// AverageElapsedMs90thPc gets the 90th percentile average response time
func (t *TestResult) AverageElapsedMs90thPc() float64 {
	return getPercentileAverage(t.responses, 0.9)
}

// AverageElapsedMs95thPc gets the 95th percentile average response time
func (t *TestResult) AverageElapsedMs95thPc() float64 {
	return getPercentileAverage(t.responses, 0.95)
}

// AverageElapsedMs99thPc gets the 99th percentile average response time
func (t *TestResult) AverageElapsedMs99thPc() float64 {
	return getPercentileAverage(t.responses, 0.99)
}

// MaxElapsedMs returns the maximum elapsed ms
func (t *TestResult) MaxElapsedMs() float64 {
	return getMax(t.responses)
}

// MinElapsedMs returns the minimum elpased ms
func (t *TestResult) MinElapsedMs() float64 {
	return getMin(t.responses)
}

// StandardDeviation returns the standard deviation of the test
func (t *TestResult) StandardDeviation() float64 {
	return getStdDev(t.responses)
}