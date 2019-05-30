package core

import (
	"fmt"
	"net/http"
	"time"
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
		resp, err := makeRequest(client, url)

		if err != nil { 
			fmt.Print(err)
		} else {
			elapsedMs := resp.ElapsedMs
			
			responseTimes[i] = elapsedMs
			fmt.Printf("request %v:\t%v\t%vms\n", i + 1, resp.StatusCode, elapsedMs)
		}
	}

	fmt.Print("\n\n")

	min, max, avg, stdDev := getMin(responseTimes), getMax(responseTimes), getAvg(responseTimes), getStdDev(responseTimes)

	fmt.Printf("min response time: %vms\n", min)
	fmt.Printf("max response time: %vms\n", max)
	fmt.Printf("avg response time: %vms\n", avg)
	fmt.Printf("expected response time: %6.2fms\n", avg + stdDev)
}

func makeRequest(client *http.Client, url string) (*TestResponse, error) {
	start := time.Now().UnixNano()

	resp, err := client.Get(url)

	end := time.Now().UnixNano()
	
	elapsedMs := (float64)((end - start) / 1000000)
	
	return &TestResponse{StatusCode: resp.StatusCode, ElapsedMs: elapsedMs}, err
}