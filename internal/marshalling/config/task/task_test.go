package task_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/task"
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
			ID: mvalue.Pattern{
				Issue: mvalue.Issue{
					Level: "info",
				},
				Pattern: "TMS-\\d+",
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

	t.Run("sets critical level for id when pattern is set without level", func(t *testing.T) {
		t.Parallel()
		tk := impl.Task{
			ID: mvalue.Pattern{
				Pattern: "TMS-\\d+",
			},
		}
		assert.Equal(
			t,
			issue.Critical,
			tk.Config().
				ID().
				Pattern().
				Level(),
		)
	})

	t.Run("keeps unknown level for id when it is incorrect", func(t *testing.T) {
		t.Parallel()
		tk := impl.Task{
			ID: mvalue.Pattern{
				Issue: mvalue.Issue{
					Level: "foo",
				},
				Pattern: "TMS-\\d+",
			},
		}
		assert.True(
			t,
			tk.Config().
				ID().
				Pattern().
				Level().
				Unknown(),
		)
	})
}
