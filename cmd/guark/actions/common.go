// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"os"
	"fmt"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/guark/guark/app/utils"
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


func CheckWorkingDir(c *cli.Context) (err error) {

	if utils.IsFile("guark.yaml") == false {
		err = fmt.Errorf("could not find: guark.yaml, cd to a guark project!")
	}

	return
}
