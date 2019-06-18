package core

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

// ReadJSONFile and parse it into a test spec
func ReadJSONFile(path string) (*ScenarioSpec, error) {
	fmt.Printf("Attempting to read file: %v\n", path)

	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	var spec ScenarioSpec

	fmt.Printf("Attemting to parse contents...\n")

	err = json.Unmarshal(bytes, &spec)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Spec file has been parsed successfully\n\n\n")

	return &spec, err
}