package core

import (
	"fmt"
	"net/http"
	"time"
	"math"
)

// Run the test
func Run(spec *TestSpec) {
	url := spec.URL
	iterations := spec.Iterations

	fmt.Printf("Target: %v\n", url)
	fmt.Printf("performing %v iterations\n", iterations)

	client := &http.Client{}

	fmt.Printf("request\tstatus code\telapsed\n")

	responseTimes := make([]float64, iterations)

	for i := 0; i < iterations; i++ {
		start := time.Now().UnixNano()

		resp, err := client.Get(url)

		end := time.Now().UnixNano()
		elapsedMs := (float64)((end - start) / 1000000)
		responseTimes[i] = elapsedMs

		if err != nil { 
			fmt.Print(err)
		} else {
			fmt.Printf("request %v:\t%v\t%vms\n", i + 1, resp.StatusCode, elapsedMs)
		}
	}

	fmt.Print("\n\n")

	min, max, avg, stdDev := getMin(responseTimes), getMax(responseTimes), getAvg(responseTimes), getStdDev(responseTimes)

	fmt.Printf("min response time: %vms\n", min)
	fmt.Printf("max response time: %vms\n", max)
	fmt.Printf("avg response time: %vms\n", avg)
	fmt.Printf("stdDev: %6.2fms\n", stdDev)
	fmt.Printf("max anticipated response time: %6.2fms\n", avg + stdDev)
}

func getMax(x []float64) float64 {
	var m float64

	for i := 0; i < len(x); i++ {
		if x[i] > m {
			m = x[i]
		}
	}

	return m
}

func getMin(x []float64) float64 {
	var m float64 = 10000000000 // arbitrarily large number

	for i := 0; i < len(x); i++ {
		if x[i] < m {
			m = x[i]
		}
	}

	return m
}

func getAvg(x []float64) float64 {
	l := (float64)(len(x))
	var sum float64

	for i := 0; i < len(x); i++ {
		sum += (float64)(x[i])
	}

	return sum / l
}

func getStdDev(x []float64) float64 {
	avg := getAvg(x)

	results := make([]float64, len(x))

	for i := 0; i < len(x); i++ {
		n := ((x[i]) - avg) * ((x[i]) - avg)
		results[i] = n
	}

	avg = getAvg(results)

	return math.Sqrt(avg)
}