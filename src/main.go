package main

import (
	"log"
	"net/http"
	"time"
	"util"
)

func main() {
	const url = "https://api.marcopolo.acc.dazn-dev.com/v1/override/1bab7192-92f6-4ca0-b0f0-29a67b275537/geofence"
	client := &http.Client{}

	log.Printf("request\tstatus code\telapsed")

	iterations := 100
	responseTimes := make([]int64, iterations)

	for i := 0; i < iterations; i++ {
		log.Printf("making GET request %v", i)

		start := time.Now().UnixNano()

		resp, err := client.Get(url)

		end := time.Now().UnixNano()
		elapsedMs := (end - start) / 1000000
		responseTimes[i] = elapsedMs

		if err != nil { 
			log.Print(err)
		} else {
			log.Printf("request %v:\t%v\t%vms", i, resp.StatusCode, elapsedMs)
		}
	}

	log.Print("\n\n")

	min, max := getMin(responseTimes), getMax(responseTimes)

	log.Printf("min response time: %v", min)
	log.Printf("max response time: %v", max)
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