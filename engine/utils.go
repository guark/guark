// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package engine

func intVal(i interface{}, def int) int {

	if i == nil {
		return def
	}

	v := i.(int)
	if v == 0 {
		v = def
	}

	return v
}
