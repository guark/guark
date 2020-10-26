// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package builders

import (
	"os"
	"path/filepath"

	"github.com/guark/guark/cmd/guark/utils"
	"github.com/guark/guark/embed"
)

// Embeded files builder.
type EmbedBuilder struct {

	// Main build.
	Build *Build

	// output dir.
	dir string
}

func (b EmbedBuilder) Before() error {

	b.Build.Log.Update("Building Embeded Files...")
	return nil
}

// Build embed.go file.
func (b EmbedBuilder) Run() error {

	files := []string{"guark.yaml"}
	err := filepath.Walk(filepath.Join(b.Build.Temp, "ui"), func(path string, i os.FileInfo, err error) error {

		if err != nil {

			return err

		} else if i.IsDir() == false {

			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return err
	}

	if err = embed.GenerateEmbed(files, utils.Path("lib", "embed.go"), "lib", filepath.Join(b.Build.Temp, "ui")); err != nil {
		return err
	}

	b.Build.Log.Done("Build Embeded Files 🗃️")
	return nil
}

func (b EmbedBuilder) Cleanup() {

}
