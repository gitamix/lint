package current

import (
	"context"
	"errors"
	"regexp"

	"github.com/gitamix/types/branch"
	task "github.com/gitamix/types/ticket"

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
	taskPattern := b.cfg.
		Task().
		ID().
		Pattern()
	if v := taskPattern.Exact(); v != "" {
		tkt := task.ParseTicket(
			br.String(),
			regexp.MustCompile(v),
		)
		if tkt.Empty() {
			issues = append(issues, issue.NewIssue(
				taskPattern.Level(),
				"task doesn't match the required pattern '"+v+"'",
			))
		}
	}
	return issues, nil
}
