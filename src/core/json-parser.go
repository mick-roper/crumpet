package core

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

// ReadJSONFile and parse it into a test spec
func ReadJSONFile(path string) (*TestSpec, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	var spec TestSpec

	err = json.Unmarshal(bytes, &spec)

	return &spec, err
}