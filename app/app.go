// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/guark/guark/log"
	"github.com/guark/guark/platform"
)

// App!
type App struct {

	// Config format version.
	Ver string `yaml:"guark"`

	// App id must be unique!
	ID string `yaml:"id"`

	// App name.
	Name string `yaml:"name"`

	// App version.
	Version string `yaml:"version"`

	// App license.
	License string `yaml:"license"`

	// App window state.
	Window struct {
		Width  int
		Height int
		Hint   int
	} `yaml:"window"`

	// App log level.
	LogLevel string `yaml:"logLevel"`

	EngineName string `yaml:"engine"`

	// App log.
	Log log.Log

	// App embeds
	Embed *Embed

	Hooks Hooks

	Funcs Funcs

	// App plugins.
	Plugins Plugins

	Engine Engine
}

// Check if app running in dev mode.
func (a App) IsDev() bool {
	return APP_MODE == "dev"
}

func (a App) Run() error {

	if err := a.init(); err != nil {
		return err
	}

	return a.Engine.Run()
}

func (a App) Quit() {
	a.Engine.Quit()
}

func (a *App) init() error {

	a.Engine.Bind("exit", func(c Context) (interface{}, error) {
		a.Quit()
		return nil, nil
	})

	a.Engine.Bind("hook", func(c Context) (interface{}, error) {

		if !c.Has("name") {
			return nil, fmt.Errorf("hook name required!")
		}

		return nil, a.Hooks.Run(c.Get("name").(string), a)
	})

	// Bind app functions.
	for name, fn := range a.Funcs {
		if err := a.Engine.Bind(name, fn); err != nil {
			return err
		}
	}

	// Bind plugin functions.
	for id, plugin := range a.Plugins {
		for name, fn := range plugin.GetFuncs() {
			if err := a.Engine.Bind(fmt.Sprintf("%s$%s", id, name), fn); err != nil {
				return err
			}
		}
	}

	return nil
}

// Get file path from appdata dir.
func (a App) Path(elem ...string) string {

	if a.IsDev() {

		cwd, err := os.Getwd()

		if err != nil {
			a.Log.Panic(err)
		}

		return filepath.Join(append([]string{cwd, "res", "static"}, elem...)...)
	}

	return filepath.Join(append([]string{platform.DATA_DIR, a.ID}, elem...)...)
}
