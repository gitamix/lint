package scope

import "github.com/gitamix/lint/config/value"

// Config represents the configuration of the commit message subject scope.
type Config struct {
	// pattern stores the pattern of the scope with its issue type level.
	pattern value.String
}

// NewConfig creates a new scope config
// with the provided pattern value.
func NewConfig(pattern value.String) Config {
	return Config{
		pattern: pattern,
	}
}

// Pattern returns the pattern of the commit message subject scope.
func (c Config) Pattern() value.String {
	return c.pattern
}
