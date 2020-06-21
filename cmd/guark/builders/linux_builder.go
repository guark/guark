// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package builders

import (
	"os"
	"os/exec"
	"path/filepath"
	// "github.com/otiai10/copy"
)

// Embeded files generator.
type LinuxBuilder struct {

	// Main build.
	Build *Build
}

func (b LinuxBuilder) Before() error {

	b.Build.Log.Update("Building for linux...")
	return nil
}

// Build and compile linux app.
func (b LinuxBuilder) Run() error {

	var (
		flags []string
		dest  string = filepath.Join(b.Build.Dest, "linux", b.Build.Info.ID)
	)

	// Set ldflags
	if b.Build.Config.Linux.Ldflags != "" {
		flags = append(flags, "-ldflags", b.Build.Config.Linux.Ldflags)
	}

	flags = append(flags, "-o", dest)

	if err := compile(flags, []string{}); err != nil {
		return err
	}

	b.Build.Log.Done("Guark linux app compiled 🙉")
	return nil
}

func (b LinuxBuilder) Cleanup() {

}

func compile(flags []string, env []string) error {

	cmd := exec.Command("go", append([]string{"build"}, flags...)...)
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
