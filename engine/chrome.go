// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.
//
// +build chrome

package engine

import (
	"fmt"
	"os"

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

func New(a *app.App) app.Engine {

	var (
		srv  *server.Server
		addr string
	)

	if a.IsDev() {

		addr = fmt.Sprintf("http://127.0.0.1:%s", os.Getenv("GUARK_DEV_PORT"))

	} else {

		srv = server.New(a)
		addr = srv.Addr()
	}

	ui, err := lorca.New("", "/tmp/profilss", 900, 700)
	if err != nil {
		panic(err)
	}

	return &ChromeEngine{
		ui:     ui,
		app:    a,
		addr:   addr,
		server: srv,
	}
}
