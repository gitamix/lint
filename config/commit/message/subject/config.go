package subject

import (
	"github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/task"
)

// Config represents the configuration
// of the commit message subject.
type Config struct {
	// desc stores the configuration
	// of the subject description text.
	//
	// The value is used to validate
	// the commit message subject description.
	desc description.Config

	// task stores the configuration
	// for task integration.
	task task.Config
}

// NewConfig creates a new subject config
// with the provided functional options.
func NewConfig(opts ...Option) Config {
	c := Config{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// Description returns the configuration
// of the subject description text.
func (c Config) Description() description.Config {
	return c.desc
}

// Task returns the configuration for task integration.
func (c Config) Task() task.Config {
	return c.task
}
