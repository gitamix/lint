package scope

import (
	"github.com/gitamix/types/commit"

	"github.com/gitamix/lint/config/commit/scope"
)

// Scope represents a linter that validates
// commit scope against the configured pattern.
type Scope struct {
	// cfg is the configuration that defines
	// the pattern used to validate the commit scope.
	cfg scope.Config

	// scp is the parsed commit scope.
	scp commit.Scope
}

// NewScope creates a new Scope linter
// with the provided commit scope and lint configuration
// that defines the pattern used to validate the commit scope.
func NewScope(
	scp commit.Scope,
	cfg scope.Config,
) Scope {
	return Scope{
		scp: scp,
		cfg: cfg,
	}
}
