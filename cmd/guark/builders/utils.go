// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package builders

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/otiai10/copy"
)

func gobuild(flags []string, engine string, env []string) error {

	cmd := exec.Command("go", append([]string{"build", "-tags", engine}, flags...)...)
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func copyStaticFiles(dest string) error {
	return copy.Copy("statics", filepath.Join(dest, "statics"))
}
