// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

import (
	"fmt"
)

// Embed struct
type Embed struct {

	// List of embeds
	Files map[string]*[]byte
}

// Get embeded data, it returns *[]byte and error
func (e Embed) Data(name string) (b *[]byte, err error) {

	b, ok := e.Files[name]

	if ok == false {
		err = fmt.Errorf("could not find: %s", name)
	}

	return
}

// Delete from embeds
func (e *Embed) Delete(name string) {

	delete(e.Files, name)
}
