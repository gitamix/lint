package commit

import (
	"github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/value"
)

// Option configures a commit Config instance.
type Option func(*Config)

// WithMessage sets the commit message config for the commit config.
func WithMessage(msg message.Config) Option {
	return func(c *Config) {
		c.msg = msg
	}
}

// WithScope sets the commit message scope config for the commit config.
func WithScope(cfg scope.Config) Option {
	return func(c *Config) {
		c.scope = cfg
	}
}

// WithTypes sets the commit types for the commit config.
func WithTypes(types value.Strings) Option {
	return func(c *Config) {
		c.types = types
	}
}
