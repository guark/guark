package app

type (
	// JS API call params
	Params struct {
		values map[string]interface{}
	}

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
	return len(p.values)
}

// Get value by key.
func (p Params) Get(key string) interface{} {
	return p.values[key]
}

// Get if key exists, if not return default.
func (p Params) GetOr(key string, def interface{}) interface{} {

	if p.Has(key) {
		return p.values[key]
	}

	return def
}

// Check params if has a key.
func (p Params) Has(key string) bool {
	_, ok := p.values[key]
	return ok
}

func NewContext(a *App, params map[string]interface{}) Context {
	return Context{
		App: a,
		Params: Params{
			values: params,
		},
	}
}
