// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package utils

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Unmarshal guark.yaml file.
func UnmarshalGuarkFile(dir string, s interface{}) (err error) {

	bytes, err := ioutil.ReadFile(filepath.Join(dir, "guark.yaml"))

	if err != nil {
		return
	}

	err = yaml.Unmarshal(bytes, s)
	return
}
