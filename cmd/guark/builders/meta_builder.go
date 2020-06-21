// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package builders

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/guark/guark/cmd/guark/utils"
)

// Meta files generator.
type MetaBuilder struct {

	// Main build.
	Build *Build
}

func (b MetaBuilder) Before() error {

	b.Build.Log.Update("Building meta files...")
	return nil
}

// Parse and build meta files.
func (b MetaBuilder) Run() (err error) {

	// Make dest targets dirs.
	for _, target := range b.Build.Targets {

		if err = meta(b, target, b.Build.Dest); err != nil {

			return
		}
	}

	b.Build.Log.Done("Guark meta files generated 🙉")
	return nil
}

func (b MetaBuilder) Cleanup() {

}

// Parse meta files. then move them in dest dir. and replace _id_ in path with App id.
func meta(b MetaBuilder, osbuild string, dest string) error {

	metaFilesDir := utils.Path("res", "meta", osbuild)
	metaFilesDest := filepath.Join(dest, osbuild)

	if osbuild == "darwin" {
		metaFilesDest = filepath.Join(metaFilesDest, "Contents")
	}

	return filepath.Walk(metaFilesDir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		metaFile := filepath.Join(metaFilesDest, strings.Replace(filepath.Base(path), "_id_", b.Build.Info.ID, -1))

		f, err := utils.Create(metaFile, 0754)

		if err != nil {
			return err
		}

		defer f.Close()

		fc, err := ioutil.ReadFile(path)

		if err != nil {
			return err
		}

		tmpl, err := template.New("meta").Parse(string(fc))

		if err != nil {
			return err
		}

		return tmpl.Execute(f, b.Build.Info)
	})
}
