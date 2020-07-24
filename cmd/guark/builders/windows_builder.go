// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package builders

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/otiai10/copy"
	"github.com/zserge/webview"
)

// Windows app builder.
type WindowsBuilder struct {

	// Main build.
	Build *Build
}

func (b WindowsBuilder) Before() error {

	b.Build.Log.Update("Building for widnows...")
	return nil
}

// Build and compile windows app.
func (b WindowsBuilder) Run() error {

	var (
		flags []string
		env   []string = []string{"CGO_ENABLED=1", "GOOS=windows"}
		dest  string   = filepath.Join(b.Build.Dest, "windows", fmt.Sprintf("%s.exe", b.Build.Info.ID))
	)

	// Set ldflags
	if b.Build.Config.Windows.Ldflags != "" {
		flags = append(flags, "-ldflags", b.Build.Config.Windows.Ldflags)
	}

	flags = append(flags, "-o", dest)

	if b.Build.Config.Windows.CC != "" {
		env = append(env, fmt.Sprintf("CC=%s", b.Build.Config.Windows.CC))
	}

	if b.Build.Config.Windows.CXX != "" {
		env = append(env, fmt.Sprintf("CXX=%s", b.Build.Config.Windows.CXX))
	}

	if err := compile(flags, env); err != nil {
		return err
	}

	if err := copy.Copy(getDlls(), filepath.Join(b.Build.Dest, "windows")); err != nil {
		return err
	}

	b.Build.Log.Done("Guark windows app compiled 🙉")
	return nil
}

func (b WindowsBuilder) Cleanup() {

}

// Get windows dlls path.
func getDlls() string {

	arch := "x86"

	if os.Getenv("GOARCH") == "amd64" {
		arch = "x64"
	}

	return filepath.Join(os.Getenv("GOPATH"), "src", pkgPath(webview.New(true)), "dll", arch)
}

// this function code was stolen from:
// https://stackoverflow.com/a/60846213/5834438
func pkgPath(v interface{}) string {
	if v == nil {
		return ""
	}

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		return val.Elem().Type().PkgPath()
	}
	return val.Type().PkgPath()
}
