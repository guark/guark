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

		e.webview.Run()
	}()

	return
}

func (e WebviewEngine) Bind(name string, fn app.Func) error {
	if err := e.webview.Bind(fmt.Sprintf("__guark_func_%s", name), fn); err != nil {
		return err
	}
	return nil
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

		addr = fmt.Sprintf("http://127.0.0.1:%s", os.Getenv("GUARK_DEBUG_PORT"))

	} else {

		srv = server.New(a)
		addr = srv.Addr()
	}

	wv := webview.New(a.IsDev())
	wv.SetTitle(a.Name)
	wv.SetSize(1000, 700, webview.Hint(a.Window.Hint))
	wv.Navigate(addr)

	return &WebviewEngine{
		server:  srv,
		webview: wv,
	}
}
