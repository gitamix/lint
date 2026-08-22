package body

import (
	"unicode/utf8"

	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of issues describing
// any validation problems with the commit message body.
//
// It returns no issues when the configured length interval is empty.
// Otherwise it counts the runes in the body and checks whether the count
// falls within the configured length interval.
// When the count is out of range, an issue is returned with the level
// from the configuration and a message describing the allowed interval.
func (b Body) Issues() []issue.Issue {
	length := b.cfg.Length()
	if length.Empty() {
		return []issue.Issue{}
	}
	issues := make([]issue.Issue, 0, 1)
	if !length.Fit(utf8.RuneCount(b.body)) {
		issues = append(issues, issue.NewIssue(
			length.Level(),
			"body is not in range ["+length.String()+"]",
		))
	}
	return issues
}
