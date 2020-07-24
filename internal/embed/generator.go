// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package embed

import (
	"fmt"
	"os"
)

func GenerateEmbed(files []string, destFile string, pkg string, root string) error {

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

var Embeds = &app.Embed{
	Files: map[string]*[]byte{
		{{- range $embed := .embeds }}
		"{{ $embed.ID }}": &{{ stringify $embed.Data }},
		{{- end }}
	},
}

`, pkg)

	e := &Embed{
		Root:     root,
		Template: tmpl,
	}

	bytes, err := e.Build(files)

	if err != nil {
		return err
	}

	_, err = d.Write(bytes)

	return err
}
