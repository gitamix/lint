package subject

import (
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/value"
)

// Option is a functional option that modifies
// the subject config on its creation.
type Option func(*Config)

// WithLength sets the allowed length interval for the subject config.
func WithLength(length value.Range) Option {
	return func(c *Config) {
		c.length = length
	}
}

// WithTask sets the task integration configuration.
func WithTask(cfg task.Config) Option {
	return func(c *Config) {
		c.task = cfg
	}
}
