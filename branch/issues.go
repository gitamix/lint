package branch

import (
	"regexp"

	"github.com/gitamix/types/branch"
	task "github.com/gitamix/types/ticket"

	lintname "github.com/gitamix/lint/branch/name"
	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of validation issues for the git branch.
//
// It validates the branch name against the configured name pattern
// and the task identifier embedded in the branch name
// against the configured task identifier pattern.
func (b *Branch) Issues() []issue.Issue {
	issues := make([]issue.Issue, 0, 2)
	nameIssues := lintname.
		NewName(
			branch.NewName(b.br.String()),
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
			b.br.String(),
			regexp.MustCompile(v),
		)
		if tkt.Empty() {
			issues = append(issues, issue.NewIssue(
				taskPattern.Level(),
				"task doesn't match the required pattern '"+v+"'",
			))
		}
	}
	return issues
}
