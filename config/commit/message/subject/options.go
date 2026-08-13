package subject

import (
	"github.com/gitamix/lint/config/commit/message/subject/scope"
	"github.com/gitamix/lint/config/ticket"
	"github.com/gitamix/lint/config/value"
)

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
func WithScope(cfg scope.Config) Option {
	return func(c *Config) {
		c.scope = cfg
	}
}

// WithLength sets the allowed length interval for the subject config.
func WithLength(length value.Range) Option {
	return func(c *Config) {
		c.length = length
	}
}

// WithTicket sets the ticket integration configuration.
func WithTicket(cfg ticket.Config) Option {
	return func(c *Config) {
		c.ticket = cfg
	}
}
