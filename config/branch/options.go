package branch

import (
	"github.com/gitamix/lint/config/branch/defaults"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/task"
)

// Option configures a Config instance.
type Option func(*Config)

// WithName sets the branch name configuration.
func WithName(name name.Config) Option {
	return func(c *Config) {
		c.name = name
	}
}

// WithTask sets the task integration configuration.
func WithTask(tkt task.Config) Option {
	return func(c *Config) {
		c.tkt = tkt
	}
}

// WithDefault sets the default branch configuration.
func WithDefault(def defaults.Config) Option {
	return func(c *Config) {
		c.def = def
	}
}
