// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.
//
// +build hybrid

package engine

import (
	"github.com/guark/guark/app"
	"github.com/zserge/lorca"
)

type HybridEngine struct {
	engine app.Engine
}

func (e HybridEngine) Init() error {
	return e.engine.Init()
}

func (e HybridEngine) Run() (err error) {
	return e.engine.Run()
}

func (e HybridEngine) Bind(name string, fn app.Func) error {
	return e.engine.Bind(name, fn)
}

func (e *HybridEngine) Eval(js string) {
	e.engine.Eval(js)
}

func (e HybridEngine) Quit() {
	e.engine.Quit()
}

func New(a *app.App) app.Engine {

	if lorca.LocateChrome() != "" {
		return newChrome(a)
	}

	return newWebview(a)
}
