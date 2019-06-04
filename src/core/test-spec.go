package core

import (
	"fmt"
	"errors"
)

// TestSpec that a test runner will process
type TestSpec struct {
	Host string `json:"host"`
	Paths []string `json:"paths"`
	Iterations int `json:"iterations"`
	Concurrency int `json:"concurrency"`
	MinDelayMs int `json:"minDelayMs"`
	MaxDelayMs int `json:"maxDelayMs"`
	Options *TestSpecOptions `json:"options"`
}

// Validate a TestSpec
func (t *TestSpec) Validate() error {
	if t.MaxDelayMs < t.MinDelayMs {
		return errors.New("'Max Delay' must be greater than or equal to 'Min Delay'")
	}

	return nil
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
	responses []*TestResponse
}

// Print the result
func (t *TestResult) Print() {
	fmt.Printf("requests:\t%v\n\n", t.RequestCount())
	fmt.Printf("median response time:\t%vms\n", t.MedianElapsedMs())
	fmt.Printf("min response time:\t%vms\n", t.MinElapsedMs())
	fmt.Printf("max response time:\t%vms\n", t.MaxElapsedMs())
	fmt.Printf("avg response time:\t%6.2fms\n", t.AverageElapsedMs())
	fmt.Printf("standard deviation:\t%6.2fms\n", t.StandardDeviation())
	fmt.Printf("90th percentile:\t%6.2fms\n", t.AverageElapsedMs90thPc())
	fmt.Printf("95th percentile:\t%6.2fms\n", t.AverageElapsedMs95thPc())
	fmt.Printf("99th percentile:\t%6.2fms\n", t.AverageElapsedMs99thPc())
	fmt.Printf("99.9th percentile:\t%6.2fms\n", t.AverageElapsedMs999thPc())
}

// RequestCount returns the number of requests
func (t *TestResult) RequestCount() uint64 {
	return t.requestCount
}

// AverageElapsedMs returns the overall average elapsed MS
func (t *TestResult) AverageElapsedMs() float64 {
	x := getElapsedMsArray(t)

	return getAvg(x)
}

// AverageElapsedMs90thPc gets the 90th percentile average response time
func (t *TestResult) AverageElapsedMs90thPc() float64 {
	x := getElapsedMsArray(t)

	return getPercentileAverage(x, 0.9)
}

// AverageElapsedMs95thPc gets the 95th percentile average response time
func (t *TestResult) AverageElapsedMs95thPc() float64 {
	x := getElapsedMsArray(t)

	return getPercentileAverage(x, 0.95)
}

// AverageElapsedMs99thPc gets the 99th percentile average response time
func (t *TestResult) AverageElapsedMs99thPc() float64 {
	x := getElapsedMsArray(t)

	return getPercentileAverage(x, 0.99)
}

// AverageElapsedMs999thPc gets the 99.9th percentile average response time
func (t *TestResult) AverageElapsedMs999thPc() float64 {
	x := getElapsedMsArray(t)

	return getPercentileAverage(x, 0.999)
}

// MaxElapsedMs returns the maximum elapsed ms
func (t *TestResult) MaxElapsedMs() float64 {
	x := getElapsedMsArray(t)

	return getMax(x)
}

// MinElapsedMs returns the minimum elpased ms
func (t *TestResult) MinElapsedMs() float64 {
	x := getElapsedMsArray(t)

	return getMin(x)
}

// StandardDeviation returns the standard deviation of the test
func (t *TestResult) StandardDeviation() float64 {
	x := getElapsedMsArray(t)

	return getStdDev(x)
}

// MedianElapsedMs returns the medium elapsed response time
func (t *TestResult) MedianElapsedMs() float64 {
	x := getElapsedMsArray(t)

	return getMedian(x)
}

func getElapsedMsArray(t *TestResult) []float64 {
	l := len(t.responses)
	array := make([]float64, l)

	for i := 0; i < l; i++ {
		array[i] = t.responses[i].ElapsedMs
	}

	return array
} 