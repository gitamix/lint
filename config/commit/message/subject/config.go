package subject

import "github.com/gitamix/lint/config/value"

// Config represents the configuration of the commit message subject.
type Config struct {
	// types stores the subject types with their issue type levels.
	types *value.Strings
}

// NewConfig creates a new subject config
// with the provided functional options.
func NewConfig(opts ...Option) *Config {
	c := &Config{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Types returns the types for the subject config.
func (c *Config) Types() *value.Strings {
	return c.types
}
