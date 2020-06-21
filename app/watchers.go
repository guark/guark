// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

type Watcher interface {

	// Watcher routine.
	Watch(*App)
}
