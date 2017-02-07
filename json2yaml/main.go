package main

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	var object interface{}
	err := json.NewDecoder(os.Stdin).Decode(&object)
	if err != nil {
		panic(err)
	}

	bs, err := yaml.Marshal(&object)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(bs)
}
