package description

import (
	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of issues describing
// any validation problems with the commit message subject description.
//
// It returns no issues when the configured length interval is empty.
// Otherwise it reports an issue when the number of characters
// in the subject description is not within the configured interval,
// using the level from the configuration.
func (d Description) Issues() []issue.Issue {
	length := d.cfg.Length()
	if length.Empty() {
		return []issue.Issue{}
	}
	issues := make([]issue.Issue, 0, 1)
	if !length.Fit(len(d.desc.String())) {
		issues = append(issues, issue.NewIssue(
			length.Level(),
			"subject description length is not in range ["+length.String()+"]",
		))
	}
	return issues
}
