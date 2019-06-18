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
	var specFile string

	flag.StringVar(&specFile, "spec-file", "", "path to the spec file")

	flag.Parse();

	if specFile == "" {
		log.Fatal("you must provide a -spec-file argument")
	}

	spec, err := core.ReadJSONFile(specFile)

	if err != nil {
		log.Fatal(err)
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