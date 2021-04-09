// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.
//
// +build webview hybrid

package engine

import (
	"fmt"

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

func (e *WebviewEngine) Init() error {
	e.webview.SetTitle(e.app.Name)
	e.webview.SetSize(
		intVal(e.app.EngineConfig["window_width"], 900),
		intVal(e.app.EngineConfig["window_height"], 700),
		hint(e.app.EngineConfig["window_hint"]),
	)
	return nil
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

func (e *WebviewEngine) Eval(js string) {
	e.webview.Eval(js)
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

func newWebview(a *app.App) app.Engine {
	srv, addr := newServer(a)

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
