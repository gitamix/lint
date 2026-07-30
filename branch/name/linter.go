package name

import (
	"github.com/gitamix/types/branch"

	"github.com/gitamix/lint/config/branch/name"
)

// Linter represents a linter to validate branch names.
type Linter struct {
	// cfg is the configuration defines
	// the rules to validate the branch name.
	cfg name.Config

	// name is the branch name to be validated.
	name branch.Name
}

// NewLinter creates a new Linter for validating
// the provided branch name using lint configuration
// definines the rules to validate the branch name.
func NewLinter(
	name branch.Name,
	cfg name.Config,
) *Linter {
	return &Linter{
		cfg:  cfg,
		name: name,
	}
}
