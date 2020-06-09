// +build windows

// Build guark app from windows machine
//
// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"fmt"
)

// TODO: add linux and darwin cross compile flags.
func getBuildFlagsAndEnvFor(target string, buildDir string) (flags []string, env []string) {

	switch target {
	case "windows":
		env = []string{fmt.Sprintf("GOOS=%s", target), "CGO_ENABLED=1"}
		flags = []string{"build", "-o", fmt.Sprintf("%s/app", buildDir)}
		return
	}

	panic(fmt.Sprintf("target %s not supported yet.", target))
}
