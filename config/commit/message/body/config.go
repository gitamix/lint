package body

import (
	"github.com/gitamix/lint/config/commit/message/body/mandate"
	"github.com/gitamix/lint/config/value"
)

// Config represents the configuration of the commit message body.
type Config struct {
	// mandate stores the configuration of the commit types
	// for which the commit message body is mandatory.
	//
	// The value is used to check whether a given commit type
	// falls under the body mandate.
	mandate mandate.Config
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

// Mandate returns the mandate config of the body config.
func (c Config) Mandate() mandate.Config {
	return c.mandate
}
