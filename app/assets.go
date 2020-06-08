// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

import (
	"fmt"
	"io/ioutil"
)

// Assets struct
type Assets struct {

	// Assets index.
	Index map[string]string

	// Assets path prefix.
	Prefix string
}

// Read asset by file name.
func (a Assets) ReadAll(file string) ([]byte, error) {

	if v, ok := a.Index[file]; ok {
		return ioutil.ReadFile(a.Prefix + v)
	}

	return nil, fmt.Errorf("could not find: %s", file)
}

// Check if assets has a file.
func (a Assets) Has(file string) bool {

	_, ok := a.Index[file]

	return ok
}
