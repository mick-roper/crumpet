package main

import "core"

func main() {
	spec := &core.TestSpec{
		URL: "https://api.marcopolo.acc.dazn-dev.com/v1/override/1bab7192-92f6-4ca0-b0f0-29a67b275537/geofence",
		Iterations: 100,
	}

	core.Run(spec)
}