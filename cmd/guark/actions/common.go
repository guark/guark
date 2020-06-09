// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/guark/guark/app/utils"
	"github.com/urfave/cli/v2"
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

func getHost() string {

	cmd := exec.Command("go", "env", "GOHOSTOS")
	out, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(out))
}
