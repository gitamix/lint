package mandate

import (
	"github.com/gitamix/types/commit"

	"github.com/gitamix/lint/config/value"
)

// Config describes the set of commit types
// for which the commit message body is mandatory.
type Config struct {
	// types stores the names of the commit types for which
	// the commit message body is mandatory.
	//
	// The value is used to check whether a given commit type
	// falls under the body mandate.
	types value.Strings
}

// NewConfig creates a new mandate config
// with the provided functional options.
func NewConfig(opts ...Option) Config {
	c := Config{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// For reports whether the commit message body is mandatory
// for the provided commit type.
func (c Config) For(typ commit.Type) bool {
	return c.types.Has(typ.String())
}
