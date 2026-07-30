package current

import (
	"context"

	"github.com/gitamix/types/branch"

	config "github.com/gitamix/lint/config/branch"
)

type (
	// gitClient abstracts the git client used
	// to interact with the version control system.
	gitClient interface {
		// CurrentBranch retrieves the current branch
		// or returns an error if the git command execution fails.
		CurrentBranch(ctx context.Context) (branch.Branch, error)
	}

	// Linter validates the current git branch against configured rules.
	Linter struct {
		// git is the git client used to retrieve the current branch.
		git gitClient

		// cfg is the configuration containing validation rules.
		cfg config.Config
	}
)

// NewLinter creates a new Linter
// with the provided git client used to retrieve the current branch
// and configuration containing validation rules.
func NewLinter(git gitClient, cfg config.Config) *Linter {
	return &Linter{
		git: git,
		cfg: cfg,
	}
}
