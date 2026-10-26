package subject

import (
	"fmt"

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

// Validate checks the commit message subject configuration for correctness.
//
// It validates the description and the task configurations in that order
// and returns the error of the first invalid one,
// wrapped with the prefix describing the failed section.
func (c Config) Validate() error {
	err := c.Description().Validate()
	if err != nil {
		return fmt.Errorf("invalid description: %w", err)
	}
	err = c.Task().Validate()
	if err != nil {
		return fmt.Errorf("invalid task: %w", err)
	}
	return nil
}
