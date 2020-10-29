// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package builders

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Linux app builder.
type LinuxBuilder struct {

	// Main build.
	Build *Build
}

func (b LinuxBuilder) Before() error {

	b.Build.Log.Update("Building Gnu/Linux App...")
	return nil
}

// Build and compile linux app.
func (b LinuxBuilder) Run() error {

	var (
		flags []string
		env   []string = []string{"CGO_ENABLED=1", "GOOS=linux"}
		dest  string   = filepath.Join(b.Build.Dest, "linux", b.Build.Info.ID)
	)

	// Set ldflags
	if b.Build.Config.Linux.Ldflags != "" {
		flags = append(flags, "-ldflags", b.Build.Config.Linux.Ldflags)
	}

	flags = append(flags, "-o", dest)

	if b.Build.Config.Linux.CC != "" {
		env = append(env, fmt.Sprintf("CC=%s", b.Build.Config.Linux.CC))
	}

	if b.Build.Config.Linux.CXX != "" {
		env = append(env, fmt.Sprintf("CXX=%s", b.Build.Config.Linux.CXX))
	}

	if err := gobuild(flags, b.Build.Info.EngineName, env); err != nil {
		return err
	}

	if err := copyStaticFiles(b.Build.Dest, "linux"); err != nil {
		return err
	}

	b.Build.Log.Done("Build Gnu/Linux App 🐧")
	return nil
}

func (b LinuxBuilder) Cleanup() {

}

func gobuild(flags []string, engine string, env []string) error {

	cmd := exec.Command("go", append([]string{"build", "-tags", engine}, flags...)...)
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
