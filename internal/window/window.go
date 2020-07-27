// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package window

import (
	"github.com/webview/webview"
)

type Window struct {
	Webview webview.WebView
}

func NewWindow(s *Server) *Window {

	wv := webview.New(s.App.IsDev())
	wv.SetTitle(s.App.Name)
	wv.SetSize(s.App.Window.Width, s.App.Window.Height, webview.Hint(s.App.Window.Hint))
	wv.Navigate(s.Addr())
	wv.Bind("__guark__call", func(fn string, args map[string]interface{}) (interface{}, error) {
		return s.App.Call(fn, args)
	})

	return &Window{
		Webview: wv,
	}
}
