// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

type Engine interface {
	Init() error
	Run() error
	Bind(name string, fn Func) error
	Eval(js string)
	Quit()
}
