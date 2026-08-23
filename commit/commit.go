package commit

import (
	"github.com/gitamix/types/commit"

	config "github.com/gitamix/lint/config/commit"
)

// Commit represents a linter that validates a single git commit
// against the configured commit message rules.
type Commit struct {
	// commit is the parsed git commit inspected by the linter.
	commit commit.Commit

	// cfg is the configuration that defines the rules
	// applied to the commit message during the lint run.
	cfg config.Config
}

// NewCommit creates a new Commit linter
// with the provided git commit and lint configuration.
func NewCommit(
	c commit.Commit,
	cfg config.Config,
) Commit {
	return Commit{
		commit: c,
		cfg:    cfg,
	}
}
