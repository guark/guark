// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package guark

import (
	"fmt"
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

	var logLevel logrus.Level

	if g.App.IsDev() {

		err = UnmarshalGuarkFile("guark.yaml", g.App)

		if err != nil {
			g.log.Error(err)
			return
		}

		logLevel = logrus.DebugLevel
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
		})

	} else {

		logLevel = logrus.WarnLevel
	}

	logrus.SetLevel(logLevel)

	g.log.Debug("config loaded.")

	g.srv.App = g.App
	g.srv.Root = g.App.Path("static")
	g.srv.Log = g.log

	g.log.Debug("starting guark server.")
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

func New(c *app.Config) *Guark {

	return &Guark{
		App: app.New(c, app.Funcs{
			"hook": func(c app.Context) (interface{}, error) {

				if c.Params.Has("name") == false {
					return nil, fmt.Errorf("could not find hook name in params")
				}

				return nil, c.App.Hooks.Run(c.Params.Get("name").(string), c.App)
			},
		}),
		log: logrus.WithFields(logrus.Fields{"context": "guark"}),
		srv: &window.Server{},
	}
}
