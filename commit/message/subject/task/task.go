package task

import (
	"github.com/gitamix/types/commit"

	"github.com/gitamix/lint/config/task"
)

// Task represents a linter that validates the task identifier
// in a commit message subject against the configured pattern.
type Task struct {
	// subj is the commit message subject
	// inspected for a task identifier.
	subj commit.Subject

	// cfg is the configuration that defines
	// the pattern used to validate the task identifier.
	cfg task.Config
}

// NewTask creates a new Task linter
// with the provided commit message subject
// and lint configuration that defines the pattern
// used to validate the task identifier.
func NewTask(
	subj commit.Subject,
	cfg task.Config,
) Task {
	return Task{
		subj: subj,
		cfg:  cfg,
	}
}
