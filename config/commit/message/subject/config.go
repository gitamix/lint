package subject

import (
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/value"
)

// Config represents the configuration of the commit message subject.
type Config struct {
	// length stores the allowed length interval of the subject text.
	//
	// The value is used to validate the number of characters
	// in the commit message subject.
	length value.Range

	// task stores the configuration for task integration.
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

// Length returns the allowed length interval of the subject text.
func (c Config) Length() value.Range {
	return c.length
}

// Task returns the configuration for task integration.
func (c Config) Task() task.Config {
	return c.task
}
