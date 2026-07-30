package current

import (
	"context"
	"errors"

	"github.com/gitamix/types/branch"

	lintname "github.com/gitamix/lint/branch/name"
	"github.com/gitamix/lint/errs"
	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of validation issues for the current git branch.
//
// It retrieves the current branch from the git client
// and validates it with the config set in the instance on created.
//
// Returns error if retrieving the current branch fails.
func (b *Branch) Issues(ctx context.Context) ([]issue.Issue, error) {
	br, err := b.git.CurrentBranch(ctx)
	if err != nil {
		return nil, errors.Join(errs.ErrGitFailed, err)
	}
	issues := make([]issue.Issue, 0, 2)
	nameIssues := lintname.
		NewName(
			branch.NewName(br.String()),
			b.cfg.Name(),
		).
		Issues()
	issues = append(issues, nameIssues...)
	return issues, nil
}
