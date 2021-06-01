// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/cli/v2"
)

var (
	ServeFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  "dir",
			Usage: "Set dir to serve.",
			Value: ".",
		},
		&cli.StringFlag{
			Name:  "port",
			Usage: "Set server port.",
			Value: "3900",
		},
	}
)

func Serve(c *cli.Context) (err error) {
	addr := fmt.Sprintf("127.0.0.1:%s", c.String("port"))
	log.Printf("Server started at: http://%s", addr)
	return http.ListenAndServe(addr, http.FileServer(http.Dir(c.String("dir"))))
}
