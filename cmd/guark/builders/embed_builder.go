// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package builders

import (
	"os"
	"path/filepath"

	"github.com/guark/guark/cmd/guark/utils"
	"github.com/guark/guark/internal/embed"
)

// Embeded files generator.
type EmbedBuilder struct {

	// Main builder.
	Main *Build

	// output dir.
	dir string
}

func (b EmbedBuilder) Before() error {

	b.Main.Log.Update("Embedding...")

	b.dir = filepath.Join(b.Main.Dest, "assets")

	return os.Mkdir(b.dir, 0754)
}

// Build embed.go file.
func (b EmbedBuilder) Build() error {

	files := []string{"guark.yaml"}
	src := filepath.Join(b.Main.Temp, "ui")
	err := filepath.Walk(src, func(path string, i os.FileInfo, err error) error {

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

	if err = embed.Embed(files, utils.Path("lib", "embed.go"), "lib", src); err != nil {
		return err
	}

	b.Main.Log.Done("Guark embed files generated 🙉")
	return nil
}

func (b *EmbedBuilder) Cleanup() {

}
