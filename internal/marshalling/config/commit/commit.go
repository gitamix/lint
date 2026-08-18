package commit

import (
	"github.com/gitamix/lint/config/commit"
	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/commit/types"
	"github.com/gitamix/lint/internal/marshalling/config/commit/message"
	"github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

// Commit is the transport representation of the commit config.
//
// Commit groups the commit message, scope, and accepted commit types and
// converts them into the domain commit.Config consumed by the linter.
type Commit struct {
	// Msg stores the transport representation of the commit message config.
	Msg message.Message `yaml:"message,omitempty"`

	// Scope stores the transport representation of the commit scope config.
	Scope value.Pattern `yaml:"scope,omitempty"`

	// Types stores the transport representation of the accepted commit types.
	Types value.Strings `yaml:"types,omitempty"`
}

// Config converts the transport representation into the domain commit.Config,
// wiring the commit message, scope, and types representations into it.
func (c Commit) Config() commit.Config {
	scp := c.Scope.Config()
	if !c.Scope.Empty() && scp.Level().Unspecified() {
		scp = scp.WithLevel(issue.Warning)
	}
	typs := c.Types.Config()
	if !c.Types.Empty() && typs.Level().Unspecified() {
		typs = typs.WithLevel(issue.Critical)
	}
	return commit.NewConfig(
		commit.WithMessage(
			c.Msg.Config(),
		),
		commit.WithScope(
			scope.NewConfig(scp),
		),
		commit.WithTypes(
			types.NewConfig(typs),
		),
	)
}
