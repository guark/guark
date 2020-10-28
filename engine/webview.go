// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.
//
// +build webview !chrome !hybrid

package engine

import (
	"fmt"
	"os"

	"github.com/guark/guark/app"
	"github.com/guark/guark/server"
	"github.com/webview/webview"
)

type WebviewEngine struct {
	app     *app.App
	addr    string
	quited  bool
	server  *server.Server
	webview webview.WebView
}

func (e WebviewEngine) Run() (err error) {

	func() {

		// For debuging let the app panic on dev mode.
		if app.APP_MODE != "dev" {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("WebviewEngine panic: %v", r)
				}
			}()
		}

		e.webview.SetSize(
			intVal(e.app.EngineConfig.Options["window_width"], 900),
			intVal(e.app.EngineConfig.Options["window_height"], 700),
			hint(e.app.EngineConfig.Options["window_hint"]),
		)
		e.webview.SetTitle(e.app.Name)
		e.webview.Navigate(e.addr)
		e.webview.Run()
	}()

	return
}

func (e *WebviewEngine) Bind(name string, fn app.Func) error {
	return e.webview.Bind(fmt.Sprintf("__guark_func_%s", name), func(args map[string]interface{}) (interface{}, error) {
		return fn(app.NewContext(e.app, args))
	})
}

func (e *WebviewEngine) Quit() {

	if e.quited {
		return
	}

	e.quited = true
	e.webview.Destroy()

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

	return &WebviewEngine{
		app:     a,
		addr:    addr,
		server:  srv,
		webview: webview.New(a.IsDev()),
	}
}

func hint(h interface{}) webview.Hint {

	var hv string

	if h != nil {
		hv = h.(string)
	}

	switch hv {
	case "min":
		return webview.HintMin
	case "max":
		return webview.HintMax
	case "fix":
		return webview.HintFixed
	}

	return webview.HintNone
}
