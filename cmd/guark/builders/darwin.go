// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package builders

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"

	"github.com/jackmordaunt/icns"
)

// Darwin app builder.
type DarwinBuilder struct {

	// Main build.
	Build *Build
}

func (b DarwinBuilder) Before() error {

	b.Build.Log.Update("Building Darwin App...")
	return nil
}

// Build and compile MacOS app.
func (b DarwinBuilder) Run() (err error) {

	var (
		flags []string
		env   []string = []string{"CGO_ENABLED=1", "GOOS=darwin"}
		dest  string   = filepath.Join(b.Build.Dest, "darwin", "Contents", "MacOS", b.Build.Info.ID)
	)

	// Set ldflags
	if b.Build.Config.Darwin.Ldflags != "" {
		flags = append(flags, "-ldflags", b.Build.Config.Darwin.Ldflags)
	}

	flags = append(flags, "-o", dest)

	if b.Build.Config.Darwin.CC != "" {
		env = append(env, fmt.Sprintf("CC=%s", b.Build.Config.Darwin.CC))
	}

	if b.Build.Config.Darwin.CXX != "" {
		env = append(env, fmt.Sprintf("CXX=%s", b.Build.Config.Darwin.CXX))
	}

	if err = gobuild(flags, b.Build.Info.EngineName, env); err != nil {
		return err
	}

	resourcesPath := filepath.Join(b.Build.Dest, "darwin", "Contents", "Resources")

	if err = copyStaticFiles(resourcesPath); err != nil {
		return err
	}

	if err = buildIcns(filepath.Join(resourcesPath, fmt.Sprintf("%s.icns", b.Build.Info.ID))); err != nil {
		return err
	}

	b.Build.Log.Done("Build Darwin App    🍎")
	return nil
}

func (b DarwinBuilder) Cleanup() {

}

func buildIcns(dest string) error {

	f, err := os.Open(filepath.Join("statics", "icon.png"))
	if err != nil {
		return err
	}
	defer f.Close()

	icon, err := png.Decode(f)
	if err != nil {
		return err
	}

	i, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer i.Close()

	return icns.Encode(i, icon)
}
