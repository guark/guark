// +build linux

// Build guark app from linux machine
//
// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"fmt"
)

// // https://blog.svgames.pl/article/cross-compiling-c-programs-for-ms-windows-using-mingw
// TODO: add darwin cross compile build flags.
func getBuildFlagsAndEnvFor(target string, buildDir string) (flags []string, env []string) {

	switch target {
	case "windows":
		env = []string{fmt.Sprintf("GOOS=%s", target), "CGO_ENABLED=1", "CC=x86_64-w64-mingw32-gcc", "CXX=x86_64-w64-mingw32-g++"}
		flags = []string{"build", "-ldflags", "-H windowsgui", "-o", fmt.Sprintf("%s/app.exe", buildDir)}
		return

	case "linux":
		env = []string{fmt.Sprintf("GOOS=%s", target), "CGO_ENABLED=1"}
		flags = []string{"build", "-o", fmt.Sprintf("%s/app", buildDir)}
		return
	}

	panic(fmt.Sprintf("target %s not supported yet.", target))
}
