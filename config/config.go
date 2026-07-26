package config

import "github.com/gitamix/lint/config/branch"

// Config represents the main configuration for the lint tool.
type Config struct {
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
