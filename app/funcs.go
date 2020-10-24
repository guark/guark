// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package app

type (
	// JS call params
	Params map[string]interface{}

	// Func context
	Context struct {
		App *App
		Params
	}

	// Type of a Func to expose to JS api.
	Func func(Context) (interface{}, error)

	// funcs to expose.
	Funcs map[string]Func
)

// Params len
func (p Params) Len() int {
	return len(p)
}

// Get value by key.
func (p Params) Get(key string) interface{} {
	return p[key]
}

// Get if key exists, if not return default.
func (p Params) GetOr(key string, def interface{}) interface{} {

	if p.Has(key) {
		return p[key]
	}

	return def
}

// Check params if has a key.
func (p Params) Has(key string) bool {
	_, ok := p[key]
	return ok
}

func NewContext(a *App, params map[string]interface{}) Context {
	return Context{
		App:    a,
		Params: params,
	}
}
