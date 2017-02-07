package y2j2y

import (
	"encoding/json"
	"io"

	"gopkg.in/yaml.v2"
)

type JSON2YAMLConverter struct{}

func (c JSON2YAMLConverter) Convert(r io.Reader, w io.Writer) error {
	var object interface{}
	err := json.NewDecoder(r).Decode(&object)
	if err != nil {
		return err
	}

	bs, err := yaml.Marshal(&object)
	if err != nil {
		return err
	}

	_, err = w.Write(bs)
	return err
}
