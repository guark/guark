// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package window

import (
	"github.com/zserge/webview"
)

type Window struct {
	Webview webview.WebView
}

func NewWindow(s *Server) *Window {

	wv := webview.New(s.App.IsDev())
	wv.SetTitle(s.App.Config.Name)
	wv.SetSize(s.App.Config.Window.Width, s.App.Config.Window.Height, webview.Hint(s.App.Config.Window.Hint))
	wv.Navigate(s.Addr())

	// todo: add bindings... here

	return &Window{
		Webview: wv,
	}
}
