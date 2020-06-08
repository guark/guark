// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package generator

import (
	"compress/gzip"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/guark/guark/app/utils"
)

func Assets(src string, destDir string, destFile string, pkg string, root string) error {

	if utils.IsDir(destDir) == false {
		return fmt.Errorf("%s is not a valid directory", destDir)
	} else if utils.IsDir(src) == false {
		return fmt.Errorf("Invalid source directory path %s", src)
	}

	d, err := os.Create(destFile)

	if err != nil {
		return err
	}

	defer d.Close()

	tmpl := fmt.Sprintf(`package %s
//
// ------ AUTO GENERATED FILE (DO NOT EDIT) ------
//

import (
	"github.com/guark/guark/app"
)

var Assets app.Assets

func init() {
	Assets = app.Assets{
		Index: map[string]string{
			{{- range $name, $embed := .embeds }}
			"{{ $name }}": "{{ uuid $name $embed }}",
			{{- end }}
		},
	}
}
`, pkg)

	e := &EmbedGenerator{
		Root:     root,
		Template: tmpl,
		Funcs: map[string]func(file string, data []byte) string{
			"uuid": func(file string, data []byte) string {

				id := uuid.New().String()
				f, err := os.Create(fmt.Sprintf("%s/%s", destDir, id))

				if err != nil {
					panic(err)
				}

				w := gzip.NewWriter(f)
				w.Write(data)
				w.Close()

				return id
			},
		},
	}

	var files []string

	filepath.Walk(src, func(path string, i os.FileInfo, err error) error {

		if err != nil {
			panic(err)
		} else if i.IsDir() == false {
			files = append(files, path)
		}

		return nil
	})

	bytes, err := e.Build(files)

	if err != nil {
		return err
	}

	_, err = d.Write(bytes)

	return err
}
