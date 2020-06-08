// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config file struct
type Config struct {

	// Config format version.
	Guark string `yaml:"guark"`

	// App id must be unique!
	ID string `yaml:"id"`

	// App name.
	Name string `yaml:"name"`

	// App version.
	Version string `yaml:"version"`

	// App license.
	License string `yaml:"license"`

	// App window state.
	Window struct {
		Width  int
		Height int
		Hint   int
	}

	// App log options.
	Log struct {

		// Log level.
		Level string

		// Output to.
		Output string
	} `yaml:"log"`
}

// Load config from file.
func LoadConfig(file string) (c *Config, err error) {

	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		return
	}

	err = yaml.Unmarshal(bytes, &c)
	return
}
