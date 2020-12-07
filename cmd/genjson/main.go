package main

import (
	"encoding/json"
	"github.com/LeakIX/l9format"
	"os"
)

func main() {
	event := &l9format.L9Event{}
	event.Service.Software.Modules = append(event.Service.Software.Modules, l9format.SoftwareModule{})
	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(event)
	if err != nil {
		panic(err)
	}
}
