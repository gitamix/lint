package types

import (
	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of issues describing
// any validation problems with the commit type.
//
// It reports an issue when the commit type is not one
// of the configured allowed types, using the level from the configuration.
// When the commit type is empty and allowed types are configured,
// the issue is escalated to critical.
// When both the commit type and the configured list are empty,
// no issues are returned.
func (t Types) Issues() []issue.Issue {
	issues := make([]issue.Issue, 0, 1)
	if !t.cfg.List().Has(t.typ.String()) {
		lvl := t.cfg.List().Level()
		if t.typ.Empty() {
			if t.cfg.List().Empty() {
				return issues
			}
			lvl = issue.Critical
		}
		issues = append(
			issues,
			issue.NewIssue(
				lvl,
				"type must be one of ["+
					t.cfg.List().String()+
					"]",
			),
		)
	}
	return issues
}
