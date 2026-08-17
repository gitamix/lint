package task

import (
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/internal/marshalling/config/task/id"
)

// Task is the transport representation
// of the task (issue) config.
//
// Task groups the task identifier config and converts it
// into the domain task config consumed by the linter.
type Task struct {
	// ID stores the transport representation
	// of the task (issue) identifier config.
	ID id.ID `yaml:"id"`
}

// Config converts the Task into the domain task config,
// wiring the task identifier representation into it.
func (t Task) Config() task.Config {
	return task.NewConfig(
		task.WithID(
			t.ID.Config(),
		),
	)
}
