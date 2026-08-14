package commit

import (
	"github.com/gitamix/lint/config/commit/message"
)

// Option configures a commit Config instance.
type Option func(*Config)

// WithMessage sets the commit message config for the commit config.
func WithMessage(msg message.Config) Option {
	return func(c *Config) {
		c.msg = msg
	}
}
