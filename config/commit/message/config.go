package message

import (
	"fmt"

	"github.com/gitamix/lint/config/commit/message/body"
	"github.com/gitamix/lint/config/commit/message/subject"
)

// Config groups together the configurable parts
// of a commit message inspected during a git branch lint run.
type Config struct {
	// subj stores the subject configuration of the commit message.
	subj subject.Config

	// body stores the body configuration of the commit message.
	body body.Config
}

// NewConfig creates a new message Config
// with the provided options applied in order.
func NewConfig(opts ...Option) Config {
	c := Config{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// Subject returns the subject config.
func (c Config) Subject() subject.Config {
	return c.subj
}

// Body returns the body config.
func (c Config) Body() body.Config {
	return c.body
}

// Validate checks the commit message configuration for correctness.
//
// It validates the subject and the body configurations
// and returns the error of the first invalid one.
func (c Config) Validate() error {
	err := c.Subject().Validate()
	if err != nil {
		return fmt.Errorf("invalid subject: %w", err)
	}
	err = c.Body().Validate()
	if err != nil {
		return fmt.Errorf("invalid body: %w", err)
	}
	return nil
}
