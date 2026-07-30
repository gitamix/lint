package name

import (
	"regexp"

	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of issues describing
// any validation problems with the branch name.
//
// It checks whether the branch name matches the required regex pattern
// from the linter configuration. If the name does not match, a single
// issue is returned describing the mismatch.
func (l *Linter) Issues() []issue.Issue {
	issues := make([]issue.Issue, 0, 1)
	pattern := l.cfg.Pattern().Exact()
	matched := regexp.
		MustCompile(pattern).
		MatchString(l.name.String())
	if !matched {
		issues = append(issues, issue.NewIssue(
			l.cfg.Pattern().Level(),
			"branch name doesn't match the required pattern '"+pattern+"'",
		))
	}
	return issues
}
