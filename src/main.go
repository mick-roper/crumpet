package main

import (
	"flag"
	"core"
)

func main() {
	var url string
	var iterations int
	var specFile string
	var concurrency int
	var maxDelayMs int

	flag.StringVar(&url, "url", "", "the url to be tested")
	flag.IntVar(&iterations, "iterations", 0, "the number of iterations")
	flag.StringVar(&specFile, "spec-file", "", "path to the spec file")
	flag.IntVar(&concurrency, "concurrency", 1, "the number of concurrent HTTP requests that can be made")
	flag.IntVar(&maxDelayMs, "max-delay", 0, "The maximum amount of delay in milliseconds")

	flag.Parse();

	var spec *core.TestSpec

	spec = &core.TestSpec{
		URL: url, 
		Iterations: iterations, 
		Concurrency: concurrency,
		MaxDelayMs: maxDelayMs,
	}

	core.Run(spec)
}