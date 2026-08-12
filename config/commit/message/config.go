package message

import (
	"github.com/gitamix/lint/config/commit/message/subject"
)

// Config groups together the configurable parts
// of a commit message inspected during a git branch lint run.
type Config struct {
	// subj stores the subject configuration of the commit message.
	subj subject.Config
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
