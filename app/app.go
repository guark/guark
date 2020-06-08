// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

import (
	"os"
	"path/filepath"

	"github.com/guark/guark/app/platform"
	"github.com/sirupsen/logrus"
)

// App!
type App struct {

	// App config.
	Config *Config

	// App assets
	Assets *Assets

	// App log.
	Log *logrus.Entry
}

// Check if app running in dev mode.
func (a App) IsDev() bool {
	return APP_MODE == "dev"
}

// Get file path from appdata dir.
func (a App) Path(elem ...string) string {

	if a.IsDev() {

		cwd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		return filepath.Join(append([]string{cwd, "res"}, elem...)...)
	}

	return filepath.Join(append([]string{platform.DATA_DIR, a.Config.ID}, elem...)...)
}
