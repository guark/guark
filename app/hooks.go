package app



import "fmt"

type (
	// App hook.
	Hook func(*App)
	// App hooks.
	Hooks map[string]Hook
)

// Run a hook.
func (h Hooks) Run(n string, a *App) error {
	if fn, ok := h[n]; ok {
		fn(a)
		return nil
	}

	return fmt.Errorf("could not find hook: %s", n)
}
