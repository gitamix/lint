package subject

import (
	"github.com/gitamix/lint/config/commit/message/subject/scope"
	"github.com/gitamix/lint/config/ticket"
	"github.com/gitamix/lint/config/value"
)

// Config represents the configuration of the commit message subject.
type Config struct {
	// types stores the subject types with their issue type levels.
	types value.Strings

	// scope stores the configuration of the subject scope.
	scope scope.Config

	// length stores the allowed length interval of the subject text.
	//
	// The value is used to validate the number of characters
	// in the commit message subject.
	length value.Range

	// ticket stores the configuration for ticket integration.
	ticket ticket.Config
}

// NewConfig creates a new subject config
// with the provided functional options.
func NewConfig(opts ...Option) Config {
	c := Config{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// Types returns the types for the subject config.
func (c Config) Types() value.Strings {
	return c.types
}

// Scope returns the config of the subject scope.
func (c Config) Scope() scope.Config {
	return c.scope
}

// Length returns the allowed length interval of the subject text.
func (c Config) Length() value.Range {
	return c.length
}

// Ticket returns the configuration for ticket integration.
func (c Config) Ticket() ticket.Config {
	return c.ticket
}
