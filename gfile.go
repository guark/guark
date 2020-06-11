package guark

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Unmarshal guark.yaml file.
func UnmarshalGuarkFile(file string, s interface{}) (err error) {

	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		return
	}

	err = yaml.Unmarshal(bytes, s)
	return
}
