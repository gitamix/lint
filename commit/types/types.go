package types

import (
	"github.com/gitamix/types/commit"

	"github.com/gitamix/lint/config/commit/types"
)

// Types represents a linter that validates
// commit type against the configured allowed types.
type Types struct {
	// cfg is the configuration that defines
	// the allowed commit types for the lint run.
	cfg types.Config

	// typ is the commit type parsed
	// from the commit message subject.
	typ commit.Type
}

// NewTypes creates a new Types linter
// with the provided commit type and lint configuration
// that defines the allowed commit types.
func NewTypes(
	typ commit.Type,
	cfg types.Config,
) Types {
	return Types{
		typ: typ,
		cfg: cfg,
	}
}
