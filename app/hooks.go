// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

import "fmt"

type (
	// hook func.
	Hook func(*App) error

	// App hooks.
	Hooks map[string]Hook
)

// Run a hook.
func (h Hooks) Run(n string, a *App) error {
	if fn, ok := h[n]; ok {
		return fn(a)
	}

	return fmt.Errorf("could not find hook: %s", n)
}
