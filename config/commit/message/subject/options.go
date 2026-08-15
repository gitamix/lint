package subject

import (
	"github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/task"
)

// Option is a functional option that modifies
// the subject config on its creation.
type Option func(*Config)

// WithDescription sets the description configuration for the subject config.
func WithDescription(desc description.Config) Option {
	return func(c *Config) {
		c.desc = desc
	}
}

// WithTask sets the task integration configuration.
func WithTask(cfg task.Config) Option {
	return func(c *Config) {
		c.task = cfg
	}
}
