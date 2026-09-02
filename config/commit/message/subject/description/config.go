package description

import "github.com/gitamix/lint/config/value"

// Config represents the configuration
// of the commit message subject description.
type Config struct {
	// length stores the allowed length interval
	// of the subject description text,
	// excluding the task identifier, type, and scope.
	//
	// The value is used to validate the number of characters
	// in the commit message subject description.
	length value.Range
}

// NewConfig creates a new description config
// with the provided functional options.
func NewConfig(opts ...Option) Config {
	c := Config{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// Length returns the allowed length interval
// of the subject description text,
// excluding the task identifier, type, and scope.
func (c Config) Length() value.Range {
	return c.length
}
