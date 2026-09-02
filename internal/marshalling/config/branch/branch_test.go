package branch_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/branch"
	mvalue "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestBranch_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero config when empty", func(t *testing.T) {
		t.Parallel()
		b := impl.Branch{}
		want := branch.NewConfig()
		assert.Equal(t, want, b.Config())
	})

	t.Run("converts task pattern into task id config", func(t *testing.T) {
		t.Parallel()
		b := impl.Branch{
			Task: mvalue.Pattern{
				Issue: mvalue.Issue{
					Level: "critical",
				},
				Pattern: "PROJ-\\d+",
			},
		}
		want := branch.NewConfig(
			branch.WithTask(
				task.NewConfig(
					task.WithID(
						id.NewConfig(
							value.NewString(issue.Critical, "PROJ-\\d+"),
						),
					),
				),
			),
		)
		assert.Equal(t, want, b.Config())
	})

	t.Run("converts name pattern into name config", func(t *testing.T) {
		t.Parallel()
		b := impl.Branch{
			Name: mvalue.Pattern{
				Issue: mvalue.Issue{
					Level: "info",
				},
				Pattern: "feat/.*",
			},
		}
		want := branch.NewConfig(
			branch.WithName(
				name.NewConfig(
					value.NewString(issue.Info, "feat/.*"),
				),
			),
		)
		assert.Equal(t, want, b.Config())
	})
}
