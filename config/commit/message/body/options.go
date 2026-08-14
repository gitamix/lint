package body

import (
	"github.com/gitamix/lint/config/commit/message/body/mandate"
	"github.com/gitamix/lint/config/value"
)

// Option is a functional option that modifies
// the body config on its creation.
type Option func(*Config)

// WithLength sets the allowed length interval for the body config.
func WithLength(length value.Range) Option {
	return func(c *Config) {
		c.length = length
	}
}

// WithMandate sets the mandate config for the body config.
func WithMandate(cfg mandate.Config) Option {
	return func(c *Config) {
		c.mandate = cfg
	}
}
