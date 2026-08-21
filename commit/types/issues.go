package types

import (
	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of issues describing
// any validation problems with the commit type.
//
// It returns no issues when type validation is not configured,
// that is when the allowed types list is empty and its level is unspecified.
// Otherwise it reports an issue when the commit type is not one
// of the configured allowed types, using the level from the configuration.
// When the commit type is empty and allowed types are configured,
// the issue is escalated to critical.
func (t Types) Issues() []issue.Issue {
	issues := make([]issue.Issue, 0, 1)
	list := t.cfg.List()
	lvl := list.Level()
	if list.Empty() && lvl.Unspecified() {
		return issues
	}
	if !list.Has(t.typ.String()) {
		if t.typ.Empty() {
			if list.Empty() {
				return issues
			}
			lvl = issue.Critical
		}
		issues = append(
			issues,
			issue.NewIssue(
				lvl,
				"type must be one of ["+
					list.String()+
					"]",
			),
		)
	}
	return issues
}
