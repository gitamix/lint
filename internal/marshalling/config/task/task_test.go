package task_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/task"
	mid "github.com/gitamix/lint/internal/marshalling/config/task/id"
	mvalue "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestTask_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero config when empty", func(t *testing.T) {
		t.Parallel()
		tk := impl.Task{}
		want := task.NewConfig()
		assert.Equal(t, want, tk.Config())
	})

	t.Run("converts id pattern into task config", func(t *testing.T) {
		t.Parallel()
		tk := impl.Task{
			ID: mid.ID{
				Pattern: mvalue.Pattern{
					Issue: mvalue.Issue{
						Level: "info",
					},
					Pattern: "TMS-\\d+",
				},
			},
		}
		want := task.NewConfig(
			task.WithID(
				id.NewConfig(
					value.NewString(issue.Info, "TMS-\\d+"),
				),
			),
		)
		assert.Equal(t, want, tk.Config())
	})

	t.Run("sets critical level when pattern is set without level", func(t *testing.T) {
		t.Parallel()
		tk := impl.Task{
			ID: mid.ID{
				Pattern: mvalue.Pattern{
					Pattern: "TMS-\\d+",
				},
			},
		}
		want := task.NewConfig(
			task.WithID(
				id.NewConfig(
					value.NewString(issue.Critical, "TMS-\\d+"),
				),
			),
		)
		assert.Equal(t, want, tk.Config())
	})
}
