package core

import (
	"fmt"
	"net/http"
	"time"
	"sync"
)

// Run the test
func Run(spec *TestSpec) {
	url := spec.URL
	iterations := spec.Iterations

	fmt.Printf("Target: %v\n", url)
	fmt.Printf("performing %v iterations\n", iterations)
	fmt.Printf("request\tstatus code\telapsed\n")

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(iterations)

	testQueue := make(chan string, iterations)
	writeQueue := make(chan string)
	responseTimesQueue := make(chan float64)

	responseTimes := make([]float64, iterations)

	// prep
	for i := 0; i < iterations; i++ {
		testQueue <- url
	}

	// execute
	fmt.Printf("starting %v workers", spec.Concurrency)

	for i := 0; i < spec.Concurrency; i++ {
		go func (wg *sync.WaitGroup) {
			client := &http.Client{}
			
			for url := range testQueue {
				resp, err := makeRequest(client, url)

				if err != nil { 
					fmt.Print(err)
				} else {
					elapsedMs := resp.ElapsedMs
					
					responseTimesQueue <- elapsedMs

					writeQueue <- fmt.Sprintf("%v\t%6.2f", resp.StatusCode, elapsedMs)

					wg.Done()
				}
			}
		}(&waitGroup)
	}

	go func() {
		for elapsedMs := range responseTimesQueue {
			responseTimes = append(responseTimes, elapsedMs)
		}
	}()

	// message writer
	go func() {
		for message := range writeQueue {
			fmt.Printf("%v\n", message)
		}
	}()

	waitGroup.Wait()

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