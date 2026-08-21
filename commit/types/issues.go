package types

import (
	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of issues describing
// any validation problems with the commit type.
//
// It reports a critical issue when the subject contains no type,
// and an issue with the configured level when the type
// is not one of the allowed commit types.
func (t Types) Issues() []issue.Issue {
	if t.cfg.List().Empty() {
		return []issue.Issue{}
	}
	issues := make([]issue.Issue, 0, 1)
	if t.typ.Empty() && !t.cfg.List().Empty() {
		issues = append(
			issues,
			issue.NewCritical(
				"subject must contain one type of ["+
					t.cfg.List().String()+
					"]",
			),
		)
		return issues
	}
	if !t.cfg.List().Has(t.typ.String()) {
		issues = append(
			issues,
			issue.NewIssue(
				t.cfg.List().Level(),
				"type must be one of ["+
					t.cfg.List().String()+
					"]",
			),
		)
	}
	return issues
}
