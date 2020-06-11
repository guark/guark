// +build darwin

// Build guark app from darwin machine
//
// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"fmt"
)

// TODO: add build flags for linux, and windows.
func getBuildFlagsAndEnvFor(target string, buildDir string, id string) (flags []string, env []string, out string) {

	switch target {
	case "darwin",
		"linux":
		out = fmt.Sprintf("%s/%s", buildDir, id)
		env = []string{fmt.Sprintf("GOOS=%s", target), "CGO_ENABLED=1"}
		flags = []string{"build"}
		return
	case "windows":
		out = fmt.Sprintf("%s/%s.exe", buildDir, id)
		env = []string{fmt.Sprintf("GOOS=%s", target), "CGO_ENABLED=1"}
		flags = []string{"build", "-ldflags", "-H windowsgui"}
		return
	}

	panic(fmt.Sprintf("target %s not supported yet.", target))
}
