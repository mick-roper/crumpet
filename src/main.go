package main

import (
	"flag"
	"core"
)

func main() {
	var url string
	var iterations int

	flag.StringVar(&url, "url", "", "the url to be tested")
	flag.IntVar(&iterations, "iterations", 0, "the number of iterations")

	flag.Parse();

	spec := &core.TestSpec{URL: url, Iterations: iterations}

	core.Run(spec)
}