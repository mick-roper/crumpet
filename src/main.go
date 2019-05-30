package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const url = "https://api.marcopolo.acc.dazn-dev.com/v1/override/1bab7192-92f6-4ca0-b0f0-29a67b275537/geofence"
	client := &http.Client{}

	fmt.Printf("request\tstatus code\telapsed\n")

	iterations := 100
	responseTimes := make([]int64, iterations)

	for i := 0; i < iterations; i++ {
		start := time.Now().UnixNano()

		resp, err := client.Get(url)

		end := time.Now().UnixNano()
		elapsedMs := (end - start) / 1000000
		responseTimes[i] = elapsedMs

		if err != nil { 
			fmt.Print(err)
		} else {
			fmt.Printf("request %v:\t%v\t%vms\n", i, resp.StatusCode, elapsedMs)
		}
	}

	fmt.Print("\n\n")

	min, max, avg := getMin(responseTimes), getMax(responseTimes), getAvg(responseTimes)

	fmt.Printf("min response time: %vms\n", min)
	fmt.Printf("max response time: %vms\n", max)
	fmt.Printf("avg response time: %vms\n", avg)
}

func getMax(x []int64) int64 {
	var m int64

	for i := 0; i < len(x); i++ {
		if x[i] > m {
			m = x[i]
		}
	}

	return m
}

func getMin(x []int64) int64 {
	var m int64 = 10000000000 // arbitrarily large number

	for i := 0; i < len(x); i++ {
		if x[i] < m {
			m = x[i]
		}
	}

	return m
}

func getAvg(x []int64) float64 {
	l := (float64)(len(x))
	var sum float64

	for i := 0; i < len(x); i++ {
		sum += (float64)(x[i])
	}

	return sum / l
}