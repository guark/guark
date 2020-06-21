// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package window

import (
	// "bytes"
	"fmt"
	// "io"
	// "mime"
	"net"
	// "net/http"
	"os"
	// "path/filepath"
	// "strings"
	"runtime"

	"github.com/guark/guark/app"
	"github.com/sirupsen/logrus"
)

type Server struct {
	App     *app.App
	Log     *logrus.Entry
	ln      net.Listener
	started bool
	window  *Window
}

func (s *Server) Start() {

	if s.started {
		s.Log.Panic("App already started!")
	}

	s.started = true

	if s.App.IsDev() == false {
		s.serve()
	}

	s.Log.Debug("Starting new window.")

	s.window = NewWindow(s)
	s.window.Webview.Run()
}

func (s Server) Addr() string {

	if s.ln != nil {
		return fmt.Sprintf("http://%s", s.ln.Addr().String())
	} else if runtime.GOOS == "windows" {
		return "fs"
	}

	return fmt.Sprintf("http://127.0.0.1:%s", os.Getenv("GUARK_DEBUG_PORT"))
}

func (s *Server) Exit() {
	s.window.Webview.Terminate()
}

func (s *Server) Close() {

	s.window.Webview.Destroy()

	if s.ln != nil {
		s.ln.Close()
	}
}
