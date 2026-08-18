package commit

import (
	"github.com/gitamix/lint/config/commit"
	"github.com/gitamix/lint/internal/marshalling/config/commit/message"
	"github.com/gitamix/lint/internal/marshalling/config/commit/scope"
	"github.com/gitamix/lint/internal/marshalling/config/commit/types"
)

// Commit is the transport representation of the commit config.
//
// Commit groups the commit message, scope, and accepted commit types and
// converts them into the domain commit.Config consumed by the linter.
type Commit struct {
	// Msg stores the transport representation of the commit message config.
	Msg message.Message `yaml:"message,omitempty"`

	// Scope stores the transport representation of the commit scope config.
	Scope scope.Scope `yaml:"scope,omitempty"`

	// Types stores the transport representation of the accepted commit types.
	Types types.Types `yaml:"types,omitempty"`
}

// Config converts the transport representation into the domain commit.Config,
// wiring the commit message, scope, and types representations into it.
func (c Commit) Config() commit.Config {
	return commit.NewConfig(
		commit.WithMessage(
			c.Msg.Config(),
		),
		commit.WithScope(
			c.Scope.Config(),
		),
		commit.WithTypes(
			c.Types.Config(),
		),
	)
}
