package branch

import (
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/ticket"
)

// Option configures a Config instance.
type Option func(*Config)

// WithName sets the branch name configuration.
func WithName(name name.Config) Option {
	return func(c *Config) {
		c.name = name
	}
}

// WithTicket sets the ticket integration configuration.
func WithTicket(tkt ticket.Config) Option {
	return func(c *Config) {
		c.tkt = tkt
	}
}
