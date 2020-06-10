package app

type (
	// Guark plugin interface.
	Plugin interface {

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
