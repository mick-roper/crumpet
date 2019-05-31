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

	flag.StringVar(&url, "url", "", "the url to be tested")
	flag.IntVar(&iterations, "iterations", 0, "the number of iterations")
	flag.StringVar(&specFile, "spec-file", "", "path to the spec file")
	flag.IntVar(&concurrency, "concurrency", 0, "the number of concurrent HTTP requests that can be made")

	flag.Parse();

	var spec *core.TestSpec
 
	// todo: find a better way to provide args

	spec = &core.TestSpec{URL: url, Iterations: iterations, Concurrency: concurrency, MaxDelay: 250}

	core.Run(spec)
}