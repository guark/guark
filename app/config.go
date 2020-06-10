// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

// App config.
type Config struct {

	// App assets
	Assets *Assets

	// App functios
	Funcs Funcs

	// App hooks.
	Hooks Hooks

	// App plugins.
	Plugins Plugins

	// App watchers.
	Watchers []Watcher
}
