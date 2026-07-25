package ticket

import (
	"github.com/gitamix/lint/config/ticket/id"
)

// Config represents configuration for ticket (issue).
type Config struct {
	// id is a configuration for ticket identifier.
	id id.Config
}

// NewConfig creates a new Config instance
// with provided options.
func NewConfig(opts ...Option) Config {
	c := Config{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// ID returns configuration for ticket identifier.
func (c Config) ID() id.Config {
	return c.id
}
