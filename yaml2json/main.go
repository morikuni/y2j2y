package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var object interface{}
	err = yaml.Unmarshal(bs, &object)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(os.Stdout).Encode(interface2Map(object))
	if err != nil {
		panic(err)
	}
}

func interface2Map(object interface{}) interface{} {
	switch t := object.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{})
		for k, v := range t {
			m[k.(string)] = interface2Map(v)
		}
		return m
	case []interface{}:
		for i, v := range t {
			t[i] = interface2Map(v)
		}
		return t
	default:
		return object
	}
}
