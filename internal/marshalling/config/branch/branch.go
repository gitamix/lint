package branch

import (
	"github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/internal/marshalling/config/value"
)

// Branch is the transport representation of the commit branch config.
//
// Branch pairs the task and branch-name patterns and converts them into
// the domain branch.Config consumed by the linter.
type Branch struct {
	// Task stores the transport representation of the task identifier pattern.
	Task value.Pattern `yaml:"task,omitempty"`

	// Name stores the transport representation of the branch name pattern.
	Name value.Pattern `yaml:"name,omitempty"`
}

// Config converts the Branch into the domain branch.Config, wiring the
// task and name patterns into their respective nested configs.
func (b Branch) Config() branch.Config {
	return branch.NewConfig(
		branch.WithTask(
			task.NewConfig(
				task.WithID(
					id.NewConfig(
						b.Task.Config(),
					),
				),
			),
		),
		branch.WithName(
			name.NewConfig(
				b.Name.Config(),
			),
		),
	)
}
