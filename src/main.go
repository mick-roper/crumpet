package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	const url = "https://api.marcopolo.acc.dazn-dev.com/v1/override/1bab7192-92f6-4ca0-b0f0-29a67b275537/geofence"
	client := &http.Client{}

	for i := 0; i < 10; i++ {
		log.Printf("making GET request %v", i)

		start := time.Now().UnixNano()

		client.Get(url)

		end := time.Now().UnixNano()

		log.Printf("request %v took %vms", i, end - start)
	}
}