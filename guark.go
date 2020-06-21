// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package guark

import (
	"fmt"

	"github.com/guark/guark/app"
	"github.com/guark/guark/internal/window"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Guark struct {
	App *app.App
	srv *window.Server
	log *logrus.Entry
}

func (g *Guark) Run() (err error) {

	g.log.Debug("Starting guark window.")
	g.srv.Start()
	return
}

func (g *Guark) Close() {

	g.srv.Close()
}

func New(c *app.Config) *Guark {

	g := &Guark{
		log: logrus.WithFields(logrus.Fields{"context": "guark"}),
	}

	g.App = app.New(c, app.Funcs{
		"hook": func(c app.Context) (interface{}, error) {

			if c.Params.Has("name") == false {
				return nil, fmt.Errorf("could not find hook name in params")
			}

			return nil, c.App.Hooks.Run(c.Params.Get("name").(string), c.App)
		},

		"exit": func(c app.Context) (interface{}, error) {

			g.srv.Exit()
			return nil, nil
		},
	})

	// Load guark yaml file.
	bs, err := g.App.Embed.Data("guark.yaml")

	if err != nil {
		g.log.Panic(err)
	}

	err = yaml.Unmarshal(*bs, g.App)

	if err != nil {
		g.log.Panic(err)
	}

	if g.App.IsDev() {

		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
		})

	} else {

		logrus.SetLevel(logLevel(g.App.LogLevel))
	}

	g.srv = &window.Server{
		App: g.App,
		Log: g.log,
	}

	g.log.Debug("Config loaded.")

	// Initialize plugins
	for _, p := range g.App.Plugins {
		p.Init(*g.App)
	}

	g.log.Debug("Plugins Initialized.")

	return g
}

func logLevel(n string) logrus.Level {

	switch n {
	case "debug":
		return logrus.DebugLevel

	case "info":
		return logrus.InfoLevel

	case "warn":
		return logrus.WarnLevel

	case "error":
		return logrus.ErrorLevel

	case "fatal":
		return logrus.FatalLevel

	case "panic":
		return logrus.PanicLevel
	}

	return logrus.WarnLevel
}
