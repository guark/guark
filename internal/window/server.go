// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package window

import (
	"fmt"
	"net"
	"os"

	"github.com/guark/guark/app"
	"github.com/sirupsen/logrus"
)

type Server struct {
	App *app.App
	Log *logrus.Entry
	// Port   string
	Root   string
	ln     net.Listener
	ran    bool
	window *Window
}

func (s *Server) Start() {

	if s.ran {
		panic("App already started!")
	}

	s.ran = true

	if s.App.IsDev() == false {
		go serve(s)
	}

	s.Log.Debug("Starting new window.")

	s.window = NewWindow(s)
	s.window.Webview.Run()
}

func (s Server) Addr() string {

	if s.ln != nil {
		return s.ln.Addr().String()
	}

	return fmt.Sprintf("http://127.0.0.1:%s", os.Getenv("GUARK_DEBUG_PORT"))
}

func (s *Server) Stop() {

	s.window.Webview.Destroy()

	if s.App.IsDev() {
		return
	}
}

func serve(s *Server) {
	// todo..
}
