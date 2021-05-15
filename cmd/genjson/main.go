package main

import (
	"encoding/json"
	"github.com/LeakIX/l9format"
	"os"
)

func main() {
	event := &l9format.L9Event{}
	// load current schema for extension :
	currentSchemaFile, err := os.Open("l9event.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(currentSchemaFile)
	err = decoder.Decode(event)
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(os.Stdout)
	err = event.UpdateFingerprint()
	if err != nil {
		panic(err)
	}
	err = encoder.Encode(event)
	if err != nil {
		panic(err)
	}
}
