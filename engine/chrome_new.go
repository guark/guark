// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.
//
// +build chrome

package engine

import (
	"github.com/guark/guark/app"
)

func New(a *app.App) app.Engine {
	return newChrome(a)
}
