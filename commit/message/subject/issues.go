package subject

import (
	"github.com/gitamix/lint/commit/message/subject/description"
	"github.com/gitamix/lint/commit/scope"
	"github.com/gitamix/lint/commit/types"
	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of issues describing
// any validation problems found in the commit message subject.
//
// It aggregates issues from four sub-linters, each built from the
// corresponding part of the subject and its configuration:
// the task identifier, the commit type, the scope, and the description.
func (s Subject) Issues() []issue.Issue {
	issues := make([]issue.Issue, 0, 4)
	issues = append(
		issues,
		types.
			NewTypes(
				s.subj.Type(),
				s.typcfg,
			).
			Issues()...,
	)
	issues = append(
		issues,
		scope.
			NewScope(
				s.subj.Scope(),
				s.scpcfg,
			).
			Issues()...,
	)
	issues = append(
		issues,
		description.
			NewDescription(
				s.subj.Description(),
				s.cfg.Description(),
			).
			Issues()...,
	)
	return issues
}
