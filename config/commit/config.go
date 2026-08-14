package commit

import (
	"github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/commit/types"
)

// Config groups together the configurable parts
// of a commit message inspected during a git lint run.
type Config struct {
	// msg stores the configuration of the commit message.
	msg message.Config

	// scope stores the configuration of the commit message scope.
	scope scope.Config

	// types stores the commit types config.
	types types.Config
}

// NewConfig creates a new commit Config
// with the provided options applied in order.
func NewConfig(opts ...Option) Config {
	c := Config{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// Message returns the commit message config.
func (c Config) Message() message.Config {
	return c.msg
}

// Scope returns the commit message scope config.
func (c Config) Scope() scope.Config {
	return c.scope
}

// Types returns the commit types config.
func (c Config) Types() types.Config {
	return c.types
}
