package config

import (
	"github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/commit"
)

// Option configures a Config instance on its creation.
type Option func(*Config)

// WithBranch sets branch-related configuration.
func WithBranch(br branch.Config) Option {
	return func(c *Config) {
		c.branch = br
	}
}

// WithCommit sets commit-related configuration.
func WithCommit(cfg commit.Config) Option {
	return func(c *Config) {
		c.commit = cfg
	}
}
