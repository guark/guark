// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package embed

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/guark/guark/app/utils"
)

type EmbedGenerator struct {
	Root     string
	Template string
	Funcs    map[string]func(file string, data []byte) string
}

func (e EmbedGenerator) Build(files []string) ([]byte, error) {

	var (
		err    error
		bytes  []byte
		embeds = map[string][]byte{}
	)

	for i := range files {

		if utils.IsDir(files[i]) {
			continue
		}

		if bytes, err = ioutil.ReadFile(files[i]); err != nil {
			return nil, err
		}

		embeds[strings.Replace(files[i], e.Root, "", 1)] = bytes
	}

	return e.parse(embeds)
}

func (e EmbedGenerator) parse(embeds map[string][]byte) ([]byte, error) {

	funcs := map[string]interface{}{
		"stringify": stringify,
	}

	for i := range e.Funcs {
		funcs[i] = e.Funcs[i]
	}

	tmpl, err := template.New("guark.embed").Funcs(funcs).Parse(e.Template)

	if err != nil {
		return nil, err
	}

	var buff bytes.Buffer

	err = tmpl.Execute(&buff, map[string]interface{}{
		"embeds": embeds,
	})

	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

// Generate embedable code from list of files, it returns parsed template and error.
func Generate(files []string, root string, tmpl string) ([]byte, error) {

	e := &EmbedGenerator{
		Root:     root,
		Template: tmpl,
	}

	return e.Build(files)
}

// Convert bytes to embedable string code.
func stringify(bytes []byte) string {

	var parts []string

	for _, v := range bytes {
		parts = append(parts, fmt.Sprintf("%d", int(v)))
	}

	return fmt.Sprintf("[]byte{%s}", strings.Join(parts, ","))
}
