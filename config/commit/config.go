package commit

import "github.com/gitamix/lint/config/commit/message"

// Config groups together the configurable parts
// of a commit message inspected during a git lint run.
type Config struct {
	// msg stores the configuration of the commit message.
	msg message.Config
}

// NewConfig creates a new commit Config
// with the provided options applied in order.
func NewConfig(opts ...Option) Config {
	c := Config{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// Message returns the commit message config.
func (c Config) Message() message.Config {
	return c.msg
}
