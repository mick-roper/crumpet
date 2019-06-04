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
	"strings"
)

// Run the test
func Run(spec *TestSpec) (*TestResult, error) {
	host := spec.Host
	iterations := spec.Iterations

	var counter uint64

	fmt.Printf("Host: %v\n", host)
	fmt.Printf("Performing %v iterations\n", iterations)

	testChan := make(chan string, iterations)
	printChan := make(chan string)
	responseTimes := make([]float64, 0)

	// prep
	for i := 0; i < iterations; i++ {
		// construct a URL from the host and a random path
		ix := rand.Intn(len(spec.Paths))
		a, b := strings.Trim(host, "/"), strings.Trim(spec.Paths[ix], "/")

		url := fmt.Sprintf("%v/%v", a, b)
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
				// delay a bit
				delay(spec)

				// make the request
				resp, err := makeRequest(client, url, spec.Options)

				if err != nil {
					// abandon this loop
					continue
				}

				elapsedMs := resp.ElapsedMs	
				responseTimes = append(responseTimes, elapsedMs)

				printChan <- fmt.Sprintf("runner %v\t%v\t%vms\t%v\t%v\n", ix, resp.StatusCode, elapsedMs, resp.Data, resp.URL)

				wg.Done() // remove due to request processing being complete

				// atomically increment the counter
				atomic.AddUint64(&counter, 1)
			}
		}(i, &waitGroup)
	}

	// message writer
	go func(wg *sync.WaitGroup) {
		for message := range printChan {
			fmt.Print(message)
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

	time.Sleep(500 * time.Millisecond)

	result := &TestResult{
		requestCount: atomic.LoadUint64(&counter),
		responses: responseTimes,
	}

	return result, nil
}

func makeRequest(client *http.Client, url string, opts *TestSpecOptions) (*TestResponse, error) {
	buffer := &bytes.Buffer{}
	req, err := http.NewRequest("GET", url, buffer)

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	req.Header.Add("User-Agent", "Crumpet-v0.0.0");

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

func delay(spec *TestSpec) {
	var delay int

	delta := spec.MaxDelayMs - spec.MinDelayMs

	if spec.MaxDelayMs > 0 {
		delay = spec.MinDelayMs + rand.Intn(delta)
	}

	time.Sleep(time.Duration(delay) * time.Millisecond)
}