package task

import (
	"regexp"

	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of issues describing
// any validation problems with the task identifier
// in the commit message subject.
//
// It returns no issues when the configured task identifier
// pattern is empty. Otherwise it compiles the pattern into a
// regular expression and extracts the ticket from the subject.
// If the pattern fails to compile, a critical issue is returned
// describing the compilation error.
// If no ticket matching the pattern is found in the subject
// an issue is returned with the level from the configuration.
func (t Task) Issues() []issue.Issue {
	tid := t.cfg.ID().Pattern()
	if tid.Empty() {
		return []issue.Issue{}
	}
	issues := make([]issue.Issue, 0, 1)
	texp := tid.Exact()
	id, err := regexp.Compile(texp)
	if err != nil {
		issues = append(
			issues,
			issue.NewCritical(
				"failed to parse task id with expression '"+
					texp+
					"': "+
					err.Error(),
			),
		)
		return issues
	}
	if t.subj.Ticket(id).Empty() {
		issues = append(
			issues,
			issue.NewIssue(
				tid.Level(),
				"ticket not found in subject by expression '"+texp+"'",
			),
		)
	}
	return issues
}
