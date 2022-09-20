package tests

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func entries(entity string) []interface{} {
	jsonFile, err := os.Open("test_entries.json")
	defer jsonFile.Close()

	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

	var data map[string][]interface{}

	json.Unmarshal(content, &data)

	return data[entity]
}
