// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"os"
	"path/filepath"
)

var wdir string

func init() {

	var err error

	wdir, err = os.Getwd()

	if err != nil {
		panic(err)
	}
}

func path(elem ...string) string {
	return filepath.Join(append([]string{wdir}, elem...)...)
}
