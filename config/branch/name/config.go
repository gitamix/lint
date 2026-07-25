package name

import (
	"github.com/gitamix/lint/config/value"
)

// Config represents configuration rules for branch names.
type Config struct {
	// pattern is a regex pattern that branch names must match.
	pattern value.String
}

// NewConfig creates a new Config istance
// with provided pattern config value.
func NewConfig(pattern value.String) Config {
	return Config{
		pattern: pattern,
	}
}

// Pattern returns pattern config value.
func (c Config) Pattern() value.String {
	return c.pattern
}
