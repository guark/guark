// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.
//
// +build chrome hybrid

package engine

import (
	"fmt"

	"github.com/guark/guark/app"
	"github.com/guark/guark/server"
	"github.com/zserge/lorca"
)

type ChromeEngine struct {
	app    *app.App
	addr   string
	quited bool
	server *server.Server
	ui     lorca.UI
}

func (e *ChromeEngine) Init() error {

	profile, err := e.app.DataFile("_profile")
	if err != nil {
		return err
	}

	e.ui, err = lorca.New(
		fmt.Sprintf("data:text/html,<html><title>%s</title></html>", e.app.Name),
		profile,
		intVal(e.app.EngineConfig["window_width"], 900),
		intVal(e.app.EngineConfig["window_height"], 700),
	)
	return err
}

func (e ChromeEngine) Run() (err error) {

	func() {

		// For debuging let the app panic on dev mode.
		if app.APP_MODE != "dev" {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("ChromeEngine panic: %v", r)
				}
			}()
		}

		e.ui.Load(e.addr)
		<-e.ui.Done()
	}()

	return
}

func (e *ChromeEngine) Bind(name string, fn app.Func) error {
	return e.ui.Bind(fmt.Sprintf("__guark_func_%s", name), func(args map[string]interface{}) (interface{}, error) {
		return fn(app.NewContext(e.app, args))
	})
}

func (e *ChromeEngine) Eval(js string) {
	e.ui.Eval(js)
}

func (e *ChromeEngine) Quit() {

	if e.quited {
		return
	}

	e.quited = true
	e.ui.Close()

	if e.server != nil {
		e.server.Close()
	}
}

func newChrome(a *app.App) app.Engine {

	srv, addr := newServer(a)

	return &ChromeEngine{
		app:    a,
		addr:   addr,
		server: srv,
	}
}
