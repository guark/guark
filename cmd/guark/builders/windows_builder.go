// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package builders

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"reflect"

	ico "github.com/Kodeworks/golang-image-ico"
	"github.com/akavel/rsrc/rsrc"
	"github.com/webview/webview"
)

// Windows app builder.
type WindowsBuilder struct {

	// Main build.
	Build *Build
}

func (b WindowsBuilder) Before() error {

	b.Build.Log.Update("Building Windows App...")

	if err := buildIcon(filepath.Join(b.Build.Temp, "app.ico")); err != nil {
		return err
	} else if err := buildManifest(b.Build, filepath.Join(b.Build.Temp, "app.manifest")); err != nil {
		return err
	}

	arch := os.Getenv("GOARCH")
	if arch == "" {
		arch = "amd64"
	}

	return rsrc.Embed("guark.syso", arch, filepath.Join(b.Build.Temp, "app.manifest"), filepath.Join(b.Build.Temp, "app.ico"))
}

// Build and compile windows app.
func (b WindowsBuilder) Run() error {

	var (
		err   error
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

	if err = gobuild(flags, b.Build.Info.EngineName, env); err != nil {
		return err
	}

	if err = copyStaticFiles(b.Build.Dest, "windows"); err != nil {
		return err
	}

	bundlerConfig, err := os.Create("bundler.yaml")
	if err != nil {
		return err
	}
	defer bundlerConfig.Close()

	if err = writeMetafile(b.Build, bundlerConfig, "guark-bundle.yaml"); err != nil {
		return err
	}

	b.Build.Log.Done("Build Windows App   🗔")

	return nil
}

func (b WindowsBuilder) Cleanup() {
	os.Remove("guark.syso")
}

func buildIcon(name string) error {

	f, err := os.Open(filepath.Join("res", "icon.png"))
	if err != nil {
		return err
	}
	defer f.Close()

	icon, err := png.Decode(f)
	if err != nil {
		return err
	}

	i, err := os.Create(name)
	if err != nil {
		return err
	}
	defer i.Close()

	return ico.Encode(i, icon)
}

func buildManifest(b *Build, name string) error {

	f, err := os.Create(name)
	if err != nil {
		return err
	}

	return writeMetafile(b, f, "app.manifest")
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
