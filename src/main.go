package main

import (
	"flag"
	"core"
	"strings"
	"log"
	"fmt"
	"time"
)

func main() {
	var host string
	var iterations int
	var specFile string
	var concurrency int
	var minDelayMs int
	var maxDelayMs int
	var paths string

	flag.StringVar(&host, "host", "", "the host to be tested")
	flag.IntVar(&iterations, "iterations", 0, "the number of iterations")
	flag.StringVar(&specFile, "spec-file", "", "path to the spec file")
	flag.IntVar(&concurrency, "concurrency", 1, "the number of concurrent HTTP requests that can be made")
	flag.IntVar(&minDelayMs, "min-delay", 0, "the minimum amount of delay in milliseconds")
	flag.IntVar(&maxDelayMs, "max-delay", 0, "the maximum amount of delay in milliseconds")
	flag.StringVar(&paths, "paths", "", "comma separated collection of paths to test")

	flag.Parse();

	var spec *core.TestSpec

	if specFile != "" {
		var err error
		spec, err = core.ReadJSONFile(specFile)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		spec = &core.TestSpec{
			Host: host, 
			Paths: strings.Split(paths, ","),
			Iterations: iterations, 
			Concurrency: concurrency,
			MinDelayMs: minDelayMs,
			MaxDelayMs: maxDelayMs,
		}
	}

	err := spec.Validate()

	if err != nil {
		log.Fatal(err)
	}

	start := time.Now().Unix()

	result, err := core.Run(spec)

	if err != nil {
		log.Fatal(err)
	}

	end := time.Now().Unix()

	fmt.Printf("tests finished in %v seconds\n\n", end - start)

	result.Print()
}