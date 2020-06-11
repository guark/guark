// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package window

import (
	"fmt"
	"net"
	"os"
	"net/http"

	"github.com/guark/guark/app"
	"github.com/sirupsen/logrus"
)

type Server struct {
	App *app.App
	Log *logrus.Entry
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
		s.serve()
	}

	s.Log.Debug("Starting new window.")

	s.window = NewWindow(s)
	s.window.Webview.Run()
}

func (s Server) Addr() string {

	if s.ln != nil {
		return fmt.Sprintf("http://%s", s.ln.Addr().String())
	}

	return fmt.Sprintf("http://127.0.0.1:%s", os.Getenv("GUARK_DEBUG_PORT"))
}

func (s *Server) Stop() {

	s.window.Webview.Destroy()

	if s.App.IsDev() {
		return
	}

	s.ln.Close()
}

func (s *Server) serve() {

	var err error

	s.ln, err = net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		panic(err)
	}

	go http.Serve(s.ln, &srvHandler{assets: s.App.Assets})
}


type srvHandler struct {
	assets *app.Assets
}


func (h srvHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		r.URL.Path = "/index.html"
	}

	f, e := h.assets.ReadAll(r.URL.Path)

	if e != nil {
		w.Write([]byte(e.Error()))
		return
	}

	w.Write(f)
	w.Write([]byte("hello"))
}
