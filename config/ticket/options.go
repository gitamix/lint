package ticket

import (
	"github.com/gitamix/lint/config/ticket/id"
)

// Option represents a functional option
// to configure Config instance on its creation.
type Option func(*Config)

// WithID sets identifier to Config instance on its creation.
func WithID(id id.Config) Option {
	return func(c *Config) {
		c.id = id
	}
}
