package core

import (
	"fmt"
	"net/http"
	"time"
	"sync"
	"sync/atomic"
	"math/rand"
	"bytes"
	"io/ioutil"
)

// Run the test
func Run(spec *TestSpec) {
	host := spec.Host
	iterations := spec.Iterations

	var counter uint64

	fmt.Printf("Host: %v\n", host)
	fmt.Printf("Performing %v iterations\n", iterations)

	testChan := make(chan string, iterations)
	printChan := make(chan string)
	responseTimes := make([]float64, iterations)

	// prep
	for i := 0; i < iterations; i++ {
		// construct a URL from the host and a random path
		ix := rand.Intn(len(spec.Paths))
		url := fmt.Sprintf("%v/%v", host, spec.Paths[ix])
		testChan <- url
	}

	// execution
	fmt.Printf("starting %v workers\n\n", spec.Concurrency)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(iterations)

	for i := 0; i < spec.Concurrency; i++ {
		go func (ix int, wg *sync.WaitGroup) {
			client := &http.Client{}

			for url := range testChan {
				resp, err := makeRequest(client, url, spec.Options)

				if err != nil { 
					panic(err)
				}

				elapsedMs := resp.ElapsedMs	
				responseTimes = append(responseTimes, elapsedMs)

				printChan <- fmt.Sprintf("runner %v\t%v\t%vms\t%v\t%v\n", ix, resp.StatusCode, elapsedMs, resp.Data, resp.URL)

				// atomically increment the counter
				atomic.AddUint64(&counter, 1)
				
				// delay a bit
				var delay int

				if spec.MaxDelayMs > 0 {
					delay = rand.Intn(spec.MaxDelayMs)
				}

				time.Sleep(time.Duration(delay) * time.Millisecond)

				wg.Add(1) // add for the benefit of the message writer
				wg.Done() // remove due to processing being complete
			}
		}(i, &waitGroup)
	}

	// message writer
	go func(wg *sync.WaitGroup) {
		for message := range printChan {
			fmt.Print(message)

			wg.Done()
		}
	}(&waitGroup)

	// progress writer
	go func() {
		for {
			time.Sleep(5 * time.Second)
			printChan <- fmt.Sprintf(" -- processed %v iterations --\n", atomic.LoadUint64(&counter))
		}
	}()

	waitGroup.Wait()

	fmt.Print("\n\n")

	min, max, avg, stdDev := getMin(responseTimes), getMax(responseTimes), getAvg(responseTimes), getStdDev(responseTimes)

	fmt.Printf("requests made: %v\n", atomic.LoadUint64(&counter))
	fmt.Printf("min response time: %vms\n", min)
	fmt.Printf("max response time: %vms\n", max)
	fmt.Printf("avg response time: %6.2fms\n", avg)
	fmt.Printf("std deviation: %6.2fms\n", stdDev)
}

func makeRequest(client *http.Client, url string, opts *TestSpecOptions) (*TestResponse, error) {
	buffer := &bytes.Buffer{}
	req, err := http.NewRequest("GET", url, buffer)

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	if opts != nil {
		for k, v := range opts.HTTPRequestHeaders {
			req.Header.Add(k, v)
		}
	}

	start := time.Now().UnixNano()

	resp, err := client.Do(req)

	end := time.Now().UnixNano()
	
	elapsedMs := (float64)((end - start) / 1000000)

	data := "OK"

	if resp.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			data = err.Error()
		} else {
			data = string(bodyBytes)
		}
	}
	
	return &TestResponse{URL: url, StatusCode: resp.StatusCode, ElapsedMs: elapsedMs, Data: data}, err
}