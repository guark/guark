// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
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

	// App log level.
	LogLevel string `yaml:"logLevel"`

	// App log.
	Log log.Logger

	// App embeds
	Embed *Embed

	Hooks Hooks

	Funcs Funcs

	// App plugins.
	Plugins Plugins

	EngineConfig map[string]interface{} `yaml:"engineConfig"`

	engine Engine
}

// Check if app running in dev mode.
func (a App) IsDev() bool {
	return APP_MODE == "dev"
}

func (a *App) Run() error {
	return a.engine.Run()
}

func (a *App) Use(engine Engine) error {

	a.engine = engine

	cfg, err := a.Embed.UngzipData("guark.yaml")
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(cfg, a); err != nil {
		return err
	}

	a.Log.SetLevel(a.LogLevel)

	return a.init()
}

func (a App) DataFile(name string) (string, error) {
	return xdg.DataFile(filepath.Join(a.ID, name))
}

func (a App) CacheFile(name string) (string, error) {
	return xdg.CacheFile(filepath.Join(a.ID, name))
}

func (a App) ConfigFile(name string) (string, error) {
	return xdg.ConfigFile(filepath.Join(a.ID, name))
}

// Get file path from appdata dir.
func (a App) Path(elem ...string) string {

	if a.IsDev() {

		cwd, err := os.Getwd()

		if err != nil {
			a.Log.Panic(err)
		}

		return filepath.Join(append([]string{cwd, "statics"}, elem...)...)
	}

	return filepath.Join(append([]string{platform.DATA_DIR, a.ID, "statics"}, elem...)...)
}

func (a App) Quit() {
	a.engine.Quit()
}

func (a *App) init() error {

	if err := a.engine.Init(); err != nil {
		return err
	}

	a.engine.Bind("exit", func(c Context) (interface{}, error) {
		a.Quit()
		return nil, nil
	})

	a.engine.Bind("hook", func(c Context) (interface{}, error) {

		if !c.Has("name") {
			return nil, fmt.Errorf("hook name required!")
		}

		return nil, a.Hooks.Run(c.Get("name").(string), a)
	})

	a.engine.Bind("env", func(c Context) (interface{}, error) {
		return map[string]interface{}{
			"app_id":      a.ID,
			"app_name":    a.Name,
			"dev_mode":    a.IsDev(),
			"app_version": a.Version,
		}, nil
	})

	// Bind app functions.
	for name, fn := range a.Funcs {
		if err := a.engine.Bind(name, fn); err != nil {
			return err
		}
	}

	// Bind plugin functions.
	for id, plugin := range a.Plugins {

		// Init the plugin.
		plugin.Init(*a)

		for name, fn := range plugin.GetFuncs() {
			if err := a.engine.Bind(fmt.Sprintf("%s$%s", id, name), fn); err != nil {
				return err
			}
		}
	}

	return nil
}
