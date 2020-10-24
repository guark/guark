// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/guark/guark/app"
)

type Server struct {
	ln net.Listener
}

func (s Server) Addr() string {
	return fmt.Sprintf("http://%s", s.ln.Addr().String())
}

func (s *Server) Close() {
	s.ln.Close()
}

func New(a *app.App) *Server {

	var (
		err error
		srv = &Server{}
	)

	srv.ln, err = net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		a.Log.Panic(err)
	}

	go http.Serve(srv.ln, &Handler{
		log:   a.Log,
		embed: a.Embed,
	})

	return srv
}
