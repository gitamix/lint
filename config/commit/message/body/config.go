package body

import "github.com/gitamix/lint/config/value"

// Config represents the configuration of the commit message body.
type Config struct {
	// length stores the allowed length interval of the body text.
	//
	// The value is used to validate the number of characters
	// in the commit message body.
	length value.Range
}

// NewConfig creates a new body config
// with the provided functional options.
func NewConfig(opts ...Option) Config {
	c := Config{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// Length returns the allowed length interval of the body text.
func (c Config) Length() value.Range {
	return c.length
}
