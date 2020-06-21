// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/guark/guark/app/platform"
	"github.com/sirupsen/logrus"
)

// App!
type App struct {

	// Config format version.
	Guark string `yaml:"guark"`

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
	Log *logrus.Entry

	// App embeds
	Embed *Embed

	// App functios
	Funcs Funcs

	// App hooks.
	Hooks Hooks

	// App plugins.
	Plugins Plugins

	// App watchers.
	Watchers []Watcher

	// Builtin functions
	bFuncs Funcs
}

// Check if app running in dev mode.
func (a App) IsDev() bool {
	return APP_MODE == "dev"
}

// Get file path from appdata dir.
// TODO: refactor.
func (a App) Path(elem ...string) string {

	if a.IsDev() {

		cwd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		return filepath.Join(append([]string{cwd, "res"}, elem...)...)
	}

	return filepath.Join(append([]string{platform.DATA_DIR, a.ID, "datadir"}, elem...)...)
}

// Call a func.
func (a *App) Call(fn string, args map[string]interface{}) (interface{}, error) {

	name := strings.Split(fn, ".")

	if len(name) > 2 || fn == "" {

		return nil, fmt.Errorf("Invalid func name: %s", fn)

	} else if len(name) == 2 {

		if _, ok := a.Plugins[name[0]]; ok {

			funcs := a.Plugins[name[0]].GetFuncs()

			if fnc, ok := funcs[name[1]]; ok {
				return fnc(NewContext(a, args))
			}
		}

		return nil, fmt.Errorf("Could not find func name: %s", fn)

	} else if fnc, ok := a.bFuncs[fn]; ok {

		return fnc(NewContext(a, args))

	} else if fnc, ok := a.Funcs[fn]; ok {

		return fnc(NewContext(a, args))
	}

	return nil, fmt.Errorf("Invalid func call: %s", fn)
}

func New(c *Config, builtin Funcs) *App {
	return &App{
		Log:      logrus.WithFields(logrus.Fields{"context": "app"}),
		Funcs:    c.Funcs,
		Hooks:    c.Hooks,
		Embed:    c.Embed,
		Plugins:  c.Plugins,
		Watchers: c.Watchers,
		bFuncs:   builtin,
	}
}
