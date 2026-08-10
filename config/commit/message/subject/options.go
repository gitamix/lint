package subject

import "github.com/gitamix/lint/config/value"

// Option is a functional option that modifies
// the subject config on its creation.
type Option func(*Config)

// WithTypes sets the types for the subject config.
func WithTypes(types value.Strings) Option {
	return func(c *Config) {
		c.types = types
	}
}

// WithScope sets the scope for the subject config.
func WithScope(v value.String) Option {
	return func(c *Config) {
		c.scope = v
	}
}
