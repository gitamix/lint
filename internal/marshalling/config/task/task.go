package task

import (
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

// Task is the transport representation
// of the task (issue) config.
//
// Task groups the task identifier config and converts it
// into the domain task config consumed by the linter.
type Task struct {
	// ID stores the transport representation
	// of the task (issue) identifier config.
	ID value.Pattern `yaml:"id,omitempty"`
}

// Config converts the Task into the domain task config,
// wiring the task identifier representation into it.
func (t Task) Config() task.Config {
	idv := t.ID.Config()
	if idv.Level().Unspecified() && !t.ID.Empty() {
		idv = idv.WithLevel(issue.Critical)
	}
	return task.NewConfig(
		task.WithID(
			id.NewConfig(idv),
		),
	)
}
