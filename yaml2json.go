package y2j2y

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type YAML2JSONConverter struct{}

func (c YAML2JSONConverter) Convert(r io.Reader, w io.Writer) error {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	var object interface{}
	err = yaml.Unmarshal(bs, &object)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(interface2Map(object))
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
