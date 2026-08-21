package scope

import (
	"regexp"

	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of issues describing
// any validation problems with the commit scope.
//
// It returns no issues when the configured pattern is empty.
// Otherwise it matches the scope against the regex pattern
// from the linter configuration.
// If the pattern fails to compile, a critical issue is returned describing
// the compilation error. If the scope does not match the pattern,
// an issue is returned with the level from the configuration.
func (s Scope) Issues() []issue.Issue {
	if s.cfg.Pattern().Empty() {
		return []issue.Issue{}
	}
	issues := make([]issue.Issue, 0, 1)
	exp := s.cfg.
		Pattern().
		Exact()
	ok, err := regexp.MatchString(
		exp,
		s.scp.String(),
	)
	if err != nil {
		issues = append(
			issues,
			issue.NewCritical(
				"failed to compile scope expression '"+
					exp+
					"': "+
					err.Error(),
			),
		)
		return issues
	}
	if !ok {
		issues = append(
			issues,
			issue.NewIssue(
				s.cfg.
					Pattern().
					Level(),
				"scope not found in subject by expression '"+exp+"'",
			),
		)
	}
	return issues
}
