// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package engine

import (
	"fmt"
	"os"

	"github.com/guark/guark/app"
	"github.com/guark/guark/server"
)

func intVal(i interface{}, def int) int {

	if i == nil {
		return def
	}

	v := i.(int)
	if v == 0 {
		v = def
	}

	return v
}

func newServer(a *app.App) (srv *server.Server, addr string) {

	if a.IsDev() {

		addr = fmt.Sprintf("http://127.0.0.1:%s", os.Getenv("GUARK_DEV_PORT"))

	} else {

		srv = server.New(a)
		addr = srv.Addr()
	}

	return
}
