// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

type (
	// Guark plugin interface.
	Plugin interface {

		// Init the plugin (called before starting the window).
		Init(App)

		// Get plugin name.
		GetName() string

		// Get plugin version.
		GetVersion() string

		// Get plugin exposed funcs to guark js API.
		GetFuncs() map[string]Func
	}

	// App plugins.
	Plugins map[string]Plugin
)
