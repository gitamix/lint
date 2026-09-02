package branch

import (
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/task"
)

// Config represents configuration for commit branch behavior.
type Config struct {
	// name stores the configuration for branch naming rules.
	name name.Config

	// tkt stores the configuration for task integration.
	tkt task.Config
}

// NewConfig creates a new Config instance
// with provided functional options to customize the configuration.
func NewConfig(opts ...Option) Config {
	c := Config{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// Name returns the configuration for branch naming rules.
func (c Config) Name() name.Config {
	return c.name
}

// Task returns the configuration for task integration.
func (c Config) Task() task.Config {
	return c.tkt
}
