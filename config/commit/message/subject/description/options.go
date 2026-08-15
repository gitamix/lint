package description

import "github.com/gitamix/lint/config/value"

// Option is a functional option that modifies
// the description config on its creation.
type Option func(*Config)

// WithLength sets the allowed length interval for the description config.
func WithLength(length value.Range) Option {
	return func(c *Config) {
		c.length = length
	}
}
