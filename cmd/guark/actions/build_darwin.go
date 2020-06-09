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
func getBuildFlagsAndEnvFor(target string, buildDir string) (flags []string, env []string) {

	switch target {
	case "darwin":
		env = []string{fmt.Sprintf("GOOS=%s", target), "CGO_ENABLED=1"}
		flags = []string{"build", "-o", fmt.Sprintf("%s/app", buildDir)}
		return
	}

	panic(fmt.Sprintf("target %s not supported yet.", target))
}
