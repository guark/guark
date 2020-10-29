// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
)

// Embed struct
type Embed struct {

	// List of embeds
	Files map[string]*[]byte
}

// Get embeded data, it returns *[]byte and error
func (e Embed) Data(name string) (b *[]byte, err error) {

	b, ok := e.Files[name]
	if !ok {
		err = fmt.Errorf("could not find: %s", name)
	}

	return
}

func (e Embed) UngzipData(name string) ([]byte, error) {

	b, err := e.Data(name)
	if err != nil {
		return nil, err
	}

	reader, err := gzip.NewReader(bytes.NewReader(*b))
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(reader)
}

// Delete from embeds
func (e *Embed) Delete(name string) {

	delete(e.Files, name)
}
