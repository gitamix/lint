package message

import (
	"github.com/gitamix/lint/config/commit/message/body"
	"github.com/gitamix/lint/config/commit/message/subject"
)

// Option is a functional option that modifies
// the message config on its creation.
type Option func(*Config)

// WithSubject sets the subject config for the message config.
func WithSubject(subj subject.Config) Option {
	return func(c *Config) {
		c.subj = subj
	}
}

// WithBody sets the body config for the message config.
func WithBody(b body.Config) Option {
	return func(c *Config) {
		c.body = b
	}
}
