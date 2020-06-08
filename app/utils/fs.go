// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package utils

import (
	"os"
)

// Check path is as directory.
func IsDir(path string) bool {

	s, err := os.Stat(path)
	return err == nil && s.IsDir()
}

// Check path is a file.
func IsFile(path string) bool {

	s, err := os.Stat(path)
	return err == nil && s.IsDir() == false
}
