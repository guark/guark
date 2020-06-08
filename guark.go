// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package guark

import (
	"os"

	"github.com/guark/guark/app"
	"github.com/guark/guark/internal/window"
	"github.com/sirupsen/logrus"
)

type Guark struct {
	App    *app.App
	srv    *window.Server
	log    *logrus.Entry
	exited bool
}

func (g *Guark) Run() (err error) {

	var logLevel logrus.Level = logrus.WarnLevel

	g.App = &app.App{
		Log: logrus.WithFields(logrus.Fields{"context": "app"}),
	}

	if g.App.IsDev() {

		logLevel = logrus.DebugLevel
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
		})
	}

	logrus.SetLevel(logLevel)

	g.log.Debug("loading config.")

	g.App.Config, err = app.LoadConfig("guark.yaml")

	if err != nil {
		g.log.Error(err)
		return
	}

	g.log.Debug("config loaded.")

	g.srv.App = g.App
	g.srv.Root = g.App.Path("static")

	if g.App.Assets != nil {
		g.App.Assets.Prefix = g.App.Path("assets")
	}

	// if g.Log == nil {
	// 	// setup default log here
	// }

	g.log.Debug("starting guark server.")

	g.srv.Log = g.log
	g.srv.Start()
	return
}

func (g *Guark) Exit() {

	if g.exited {
		return
	}

	g.exited = true
	g.srv.Stop()
}

func New() *Guark {

	return &Guark{
		log: logrus.WithFields(logrus.Fields{"context": "guark"}),
		srv: &window.Server{
			Port: os.Getenv("GUARK_DEBUG_PORT"), // this will be ignored on production
		},
	}
}
