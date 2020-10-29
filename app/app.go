// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/guark/guark/log"
	"github.com/guark/guark/platform"
	"gopkg.in/yaml.v2"
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

	// App log.
	Log log.Log

	// App embeds
	Embed *Embed

	Hooks Hooks

	Funcs Funcs

	// App plugins.
	Plugins Plugins

	EngineConfig map[string]interface{} `yaml:"engine"`

	backend Engine
}

// Check if app running in dev mode.
func (a App) IsDev() bool {
	return APP_MODE == "dev"
}

func (a *App) Run() error {
	return a.backend.Run()
}

func (a App) Quit() {
	a.backend.Quit()
}

func (a *App) Use(eng Engine) error {

	a.backend = eng

	cfg, err := a.Embed.Data("guark.yaml")
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(*cfg, a); err != nil {
		return err
	}

	return a.init()
}

func (a *App) init() error {

	a.backend.Bind("exit", func(c Context) (interface{}, error) {
		a.Quit()
		return nil, nil
	})

	a.backend.Bind("hook", func(c Context) (interface{}, error) {

		if !c.Has("name") {
			return nil, fmt.Errorf("hook name required!")
		}

		return nil, a.Hooks.Run(c.Get("name").(string), a)
	})

	// Bind app functions.
	for name, fn := range a.Funcs {
		if err := a.backend.Bind(name, fn); err != nil {
			return err
		}
	}

	// Bind plugin functions.
	for id, plugin := range a.Plugins {

		// Init the plugin.
		plugin.Init(*a)

		for name, fn := range plugin.GetFuncs() {
			if err := a.backend.Bind(fmt.Sprintf("%s$%s", id, name), fn); err != nil {
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
