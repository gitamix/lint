package name

import (
	"fmt"
	"regexp"

	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of issues describing
// any validation problems with the branch name.
//
// It checks whether the branch name matches the required regex pattern
// from the linter configuration. If the name does not match, a single
// issue is returned describing the mismatch.
func (n *Name) Issues() []issue.Issue {
	issues := make([]issue.Issue, 0, 1)
	pattern := n.cfg.Pattern().Exact()
	matched := regexp.
		MustCompile(pattern).
		MatchString(n.name.String())
	if !matched {
		issues = append(issues, issue.NewIssue(
			n.cfg.Pattern().Level(),
			fmt.Sprintf(
				"branch name '%s' doesn't match the required pattern '%s'",
				n.name.String(),
				pattern,
			),
		))
	}
	return issues
}
