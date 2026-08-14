package types

import "github.com/gitamix/lint/config/value"

// Config represents the configuration of the commit types.
type Config struct {
	// types stores the commit types with their issue type levels.
	types value.Strings
}

// NewConfig creates a new types config
// with the provided types value.
func NewConfig(types value.Strings) Config {
	return Config{
		types: types,
	}
}

// Types returns the commit types with their issue type levels.
func (c Config) Types() value.Strings {
	return c.types
}
