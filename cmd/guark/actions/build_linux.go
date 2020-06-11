// +build linux

// Build guark app from linux machine
//
// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"fmt"
)

// TODO: add darwin cross compile build flags.
func getBuildFlagsAndEnvFor(target string, buildDir string, id string) (flags []string, env []string, out string) {

	switch target {
	case "windows":
		out = fmt.Sprintf("%s/%s.exe", buildDir, id)
		env = []string{fmt.Sprintf("GOOS=%s", target), "CGO_ENABLED=1", "CC=x86_64-w64-mingw32-gcc", "CXX=x86_64-w64-mingw32-g++"}
		flags = []string{"build", "-ldflags", "-H windowsgui"}
		return

	case "linux",
		"darwin":
		out = fmt.Sprintf("%s/%s", buildDir, id)
		env = []string{fmt.Sprintf("GOOS=%s", target), "CGO_ENABLED=1"}
		flags = []string{"build"}
		return
	}

	panic(fmt.Sprintf("target %s not supported yet.", target))
}
