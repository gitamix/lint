package branch

import (
	"github.com/gitamix/types/branch"

	config "github.com/gitamix/lint/config/branch"
)

// Branch represents a linter that validates
// a git branch against the configured rules.
type Branch struct {
	// br is the git branch to be validated.
	br branch.Branch

	// cfg is the configuration defining
	// the rules to validate the branch.
	cfg config.Config
}

// NewBranch creates a new Branch linter for the provided git branch,
// using the configuration defining the rules to validate it.
func NewBranch(br branch.Branch, cfg config.Config) *Branch {
	return &Branch{
		br:  br,
		cfg: cfg,
	}
}
