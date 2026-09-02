package message

import (
	"github.com/gitamix/lint/commit/message/body"
	"github.com/gitamix/lint/commit/message/subject"
	"github.com/gitamix/lint/commit/message/task"
	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of issues describing
// any validation problems found in the commit message.
//
// It aggregates issues from the subject and body sub-linters.
// When the body is empty and the configuration marks it as mandatory
// for the commit type, a critical issue is returned instead of
// running the body sub-linter.
func (m Message) Issues() []issue.Issue {
	issues := make([]issue.Issue, 0, 5)
	issues = append(
		issues,
		task.
			NewTask(
				m.msg,
				m.cfg.Subject().Task(),
			).
			Issues()...,
	)
	issues = append(
		issues,
		subject.
			NewSubject(
				m.msg.Subject(),
				m.cfg.Subject(),
				m.typcfg,
				m.scpcfg,
			).
			Issues()...,
	)
	typ := m.msg.Subject().Type()
	if m.msg.Body().Empty() {
		if m.cfg.Body().Mandate().For(typ) {
			issues = append(issues, issue.NewCritical(
				"body is required for type '"+typ.String()+"'",
			))
		}
	} else {
		issues = append(issues, body.
			NewBody(
				m.msg.Body(),
				m.cfg.Body(),
			).
			Issues()...,
		)
	}
	return issues
}
