package config

import (
	"github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/commit"
)

// Config represents the main configuration for the lint tool.
type Config struct {
	// commit stores the configuration for commit-related rules.
	commit commit.Config

	// branch stores the configuration for branch-related rules.
	branch branch.Config
}

// NewConfig creates a new Config instance
// with the provided configurations.
func NewConfig(opts ...Option) *Config {
	c := &Config{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Branch returns the configuration for branch-related rules.
func (c Config) Branch() branch.Config {
	return c.branch
}

// Commit returns the configuration for commit-related rules.
func (c Config) Commit() commit.Config {
	return c.commit
}
