package mandate

import "github.com/gitamix/lint/config/value"

// Option is a functional option that modifies
// the mandate config on its creation.
type Option func(*Config)

// WithTypes sets the commit types config for the commit config.
func WithTypes(types value.Strings) Option {
	return func(c *Config) {
		c.types = types
	}
}
