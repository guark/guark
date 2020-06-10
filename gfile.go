package guark

import (
	"io/ioutil"

	"github.com/guark/guark/app"
	"gopkg.in/yaml.v2"
)

// Unmarshal guark.yaml file.
func UnmarshalGuarkFile(file string, app *app.App) (err error) {

	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		return
	}

	err = yaml.Unmarshal(bytes, app)
	return
}
