// +build windows

// Build guark app from windows machine
//
// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"fmt"
	"path/filepath"
)

// TODO: add linux and darwin cross compile flags.
func getBuildFlagsAndEnvFor(target string, buildDir string, id string) (flags []string, env []string, out string) {

	switch target {
	case "windows":
		out = filepath.Join(buildDir, fmt.Sprintf("%s.exe", id))
		env = []string{fmt.Sprintf("GOOS=%s", target), "CGO_ENABLED=1"}
		flags = []string{"build", "-ldflags", "-H windowsgui"}
		return
	case "linux",
		"darwin":
		out = filepath.Join(buildDir, id)
		env = []string{fmt.Sprintf("GOOS=%s", target), "CGO_ENABLED=1"}
		flags = []string{"build"}
		return
	}

	panic(fmt.Sprintf("target %s not supported yet.", target))
}
