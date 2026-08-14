package task

import (
	"github.com/gitamix/lint/config/task/id"
)

// Config represents configuration for task (issue).
type Config struct {
	// id is a configuration for task identifier.
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

// ID returns configuration for task identifier.
func (c Config) ID() id.Config {
	return c.id
}
