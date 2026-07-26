package config

import (
	"github.com/gitamix/lint/config/branch"
)

// Option configures a Config instance on its creation.
type Option func(*Config)

// WithBranch sets branch-related configuration.
func WithBranch(br branch.Config) Option {
	return func(c *Config) {
		c.branch = br
	}
}
