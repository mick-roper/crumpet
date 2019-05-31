package core

import (
	"fmt"
	"net/http"
	"time"
	"sync"
	"sync/atomic"
	"math/rand"
)

// Run the test
func Run(spec *TestSpec) {
	url := spec.URL
	iterations := spec.Iterations

	var counter uint64

	fmt.Printf("Target: %v\n", url)
	fmt.Printf("performing %v iterations\n", iterations)
	fmt.Printf("request\tstatus code\telapsed\n")

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(iterations)

	testQueue := make(chan string, iterations)
	writeQueue := make(chan string)

	responseTimes := make([]float64, iterations)

	// prep
	for i := 0; i < iterations; i++ {
		testQueue <- url
	}

	// execute
	fmt.Printf("starting %v workers", spec.Concurrency)

	for i := 0; i < spec.Concurrency; i++ {
		go func (ix int, wg *sync.WaitGroup) {
			client := &http.Client{}

			for url := range testQueue {
				resp, err := makeRequest(client, url)

				if err != nil { 
					fmt.Print(err)
				} else {
					elapsedMs := resp.ElapsedMs
					
					responseTimes = append(responseTimes, elapsedMs)

					writeQueue <- fmt.Sprintf("runner %v\t%v\t%vms\n", ix, resp.StatusCode, elapsedMs)

					atomic.AddUint64(&counter, 1)
					
					// todo: delay a bit
					var delay int

					if spec.MaxDelay > 0 {
						delay = rand.Intn(spec.MaxDelay)
					}

					time.Sleep(time.Duration(delay) * time.Millisecond)

					wg.Add(1) // add for the benefit of the message writer
					wg.Done() // remove due to processing being complete
				}
			}
		}(i, &waitGroup)
	}

	// message writer
	go func(wg *sync.WaitGroup) {
		for message := range writeQueue {
			fmt.Print(message)

			wg.Done()
		}
	}(&waitGroup)

	waitGroup.Wait()

	fmt.Print("\n\n")

	min, max, avg, stdDev := getMin(responseTimes), getMax(responseTimes), getAvg(responseTimes), getStdDev(responseTimes)

	fmt.Printf("requests made: %v\n", atomic.LoadUint64(&counter))
	fmt.Printf("min response time: %vms\n", min)
	fmt.Printf("max response time: %vms\n", max)
	fmt.Printf("avg response time: %6.2fms\n", avg)
	fmt.Printf("std deviation: %6.2fms\n", stdDev)
}

func makeRequest(client *http.Client, url string) (*TestResponse, error) {
	start := time.Now().UnixNano()

	resp, err := client.Get(url)

	if err != nil {
		defer resp.Body.Close()
	}

	end := time.Now().UnixNano()
	
	elapsedMs := (float64)((end - start) / 1000000)
	
	return &TestResponse{StatusCode: resp.StatusCode, ElapsedMs: elapsedMs}, err
}