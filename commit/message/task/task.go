package task

import (
	"github.com/gitamix/types/commit"

	"github.com/gitamix/lint/config/task"
)

// Task represents a linter that validates the task identifier
// in a commit message against the configured pattern.
type Task struct {
	// msg is the commit message inspected for a task identifier.
	//
	// The ticket is extracted from the subject line of the message,
	// read from the raw bytes preserved by ParseMessage, so a leading
	// ticket prefix that ParseSubject strips is still detected.
	//
	// The message must carry raw bytes for ticket extraction to work:
	// a Message built with NewMessage without the WithRaw option has no
	// raw bytes, so Ticket always returns empty and Issues reports a
	// false "ticket not found". Prefer ParseMessage, which records the
	// raw bytes, when constructing the message passed to NewTask.
	msg commit.Message

	// cfg is the configuration that defines
	// the pattern used to validate the task identifier.
	cfg task.Config
}

// NewTask creates a new Task linter
// with the provided commit message
// and lint configuration that defines the pattern
// used to validate the task identifier.
func NewTask(
	msg commit.Message,
	cfg task.Config,
) Task {
	return Task{
		msg: msg,
		cfg: cfg,
	}
}
