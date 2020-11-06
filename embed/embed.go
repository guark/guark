// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package embed

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"text/template"

	"github.com/guark/guark/utils"
)

type (
	Item struct {
		ID   string
		Path string
		Data []byte
	}

	Embed struct {
		Root     string
		Template string
		Funcs    map[string]func(file string, data []byte) string
	}
)

// Build embed file.
func (e Embed) Build(files []string) ([]byte, error) {

	var (
		err    error
		data   []byte
		embeds []*Item
	)

	for i := range files {

		if utils.IsDir(files[i]) {
			continue
		}

		if data, err = ioutil.ReadFile(files[i]); err != nil {
			return nil, err
		}

		var gzdata bytes.Buffer
		w := gzip.NewWriter(&gzdata)
		w.Write(data)
		w.Close()

		id := strings.Replace(files[i], e.Root, "", 1)
		if runtime.GOOS == "windows" {
			id = strings.Replace(id, `\`, "/", -1)
		}

		embeds = append(embeds, &Item{
			ID:   id,
			Path: files[i],
			Data: gzdata.Bytes(),
		})
	}

	return e.parse(embeds)
}

// Parse embed template and return bytes.
func (e Embed) parse(embeds []*Item) ([]byte, error) {

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

	e := &Embed{
		Root:     root,
		Template: tmpl,
	}

	return e.Build(files)
}

// Convert bytes to embedable string code.
func stringify(bs []byte) string {

	var b []string

	for _, v := range bs {
		b = append(b, fmt.Sprintf("%d", int(v)))
	}

	return fmt.Sprintf("[]byte{%s}", strings.Join(b, ","))
}
