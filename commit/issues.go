package commit

import (
	"github.com/gitamix/lint/commit/message"
	"github.com/gitamix/lint/issue"
)

// Issues returns a slice of issues describing
// any validation problems found in the commit message.
//
// It delegates to the message sub-linter, which aggregates issues
// from the task, subject, and body sub-linters.
func (c Commit) Issues() []issue.Issue {
	return message.
		NewMessage(
			c.commit.Message(),
			c.cfg.Message(),
			c.cfg.Types(),
			c.cfg.Scope(),
		).Issues()
}
