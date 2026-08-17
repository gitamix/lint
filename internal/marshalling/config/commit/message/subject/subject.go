package subject

import (
	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/internal/marshalling/config/commit/message/subject/description"
	"github.com/gitamix/lint/internal/marshalling/config/task"
)

// Subject is the transport representation
// of the commit message subject config.
//
// Subject groups the subject description
// and the task integration pattern
// and converts them into the domain subject config
// consumed by the linter.
type Subject struct {
	// Desc stores the transport representation
	// of the subject description config.
	Desc description.Description `yaml:"description,omitempty"`

	// Task stores the transport representation
	// of the task integration pattern.
	Task task.Task `yaml:"task,omitempty"`
}

// Config converts the transport representation
// into the domain subject config.
func (s Subject) Config() subject.Config {
	return subject.NewConfig(
		subject.WithTask(
			s.Task.Config(),
		),
		subject.WithDescription(
			s.Desc.Config(),
		),
	)
}
