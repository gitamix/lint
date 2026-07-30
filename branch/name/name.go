package name

import (
	"github.com/gitamix/types/branch"

	"github.com/gitamix/lint/config/branch/name"
)

// Name represents a linter to validate branch names.
type Name struct {
	// cfg is the configuration defines
	// the rules to validate the branch name.
	cfg name.Config

	// name is the branch name to be validated.
	name branch.Name
}

// NewName creates a new Name for validating
// the provided branch name using lint configuration
// definines the rules to validate the branch name.
func NewName(
	name branch.Name,
	cfg name.Config,
) *Name {
	return &Name{
		cfg:  cfg,
		name: name,
	}
}
