package subject_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/commit/message/subject"
	tdescription "github.com/gitamix/lint/internal/marshalling/config/commit/message/subject/description"
	mtask "github.com/gitamix/lint/internal/marshalling/config/task"
	mvalue "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestSubject_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero config when empty", func(t *testing.T) {
		t.Parallel()
		s := impl.Subject{}
		want := subject.NewConfig()
		assert.Equal(t, want, s.Config())
	})

	t.Run("converts description config into subject config", func(t *testing.T) {
		t.Parallel()
		s := impl.Subject{
			Desc: tdescription.Description{
				Length: mvalue.Range{
					Issue: mvalue.Issue{
						Level: "info",
					},
					Min: 1,
					Max: 72,
				},
			},
		}
		want := subject.NewConfig(
			subject.WithDescription(
				description.NewConfig(
					description.WithLength(
						value.NewRange(issue.Info, 1, 72),
					),
				),
			),
		)
		assert.Equal(t, want, s.Config())
	})

	t.Run("sets warning lvl for desc len when it is not set", func(t *testing.T) {
		t.Parallel()
		s := impl.Subject{
			Desc: tdescription.Description{
				Length: mvalue.Range{
					Min: 1,
					Max: 72,
				},
			},
		}
		want := subject.NewConfig(
			subject.WithDescription(
				description.NewConfig(
					description.WithLength(
						value.NewRange(issue.Warning, 1, 72),
					),
				),
			),
		)
		assert.Equal(t, want, s.Config())
	})

	t.Run("keeps unknown lvl for desc len when it is incorrect", func(t *testing.T) {
		t.Parallel()
		s := impl.Subject{
			Desc: tdescription.Description{
				Length: mvalue.Range{
					Issue: mvalue.Issue{
						Level: "foo",
					},
					Min: 1,
					Max: 72,
				},
			},
		}
		assert.True(
			t,
			s.Config().
				Description().
				Length().
				Level().
				Unknown(),
		)
	})

	t.Run("sets critical lvl for task when it is not set", func(t *testing.T) {
		t.Parallel()
		s := impl.Subject{
			Task: mtask.Task{
				ID: mvalue.Pattern{
					Pattern: `(TASK|PROJ|BUG)-[0-9]+`,
				},
			},
		}
		want := subject.NewConfig(
			subject.WithTask(
				task.NewConfig(
					task.WithID(
						id.NewConfig(
							value.NewString(
								issue.Critical,
								`(TASK|PROJ|BUG)-[0-9]+`,
							),
						),
					),
				),
			),
		)
		assert.Equal(t, want, s.Config())
	})

	t.Run("keeps unknown lvl for task when it is incorrect", func(t *testing.T) {
		t.Parallel()
		s := impl.Subject{
			Task: mtask.Task{
				ID: mvalue.Pattern{
					Issue: mvalue.Issue{
						Level: "foo",
					},
					Pattern: `(TASK|PROJ|BUG)-[0-9]+`,
				},
			},
		}
		assert.True(
			t,
			s.Config().
				Task().
				ID().
				Pattern().
				Level().
				Unknown(),
		)
	})

	t.Run("desc lvl does not affect task lvl on unspecified", func(t *testing.T) {
		t.Parallel()
		s := impl.Subject{
			Desc: tdescription.Description{
				Length: mvalue.Range{
					Issue: mvalue.Issue{
						Level: "warning",
					},
					Min: 1,
					Max: 72,
				},
			},
			Task: mtask.Task{
				ID: mvalue.Pattern{
					Pattern: `(TASK|PROJ|BUG)-[0-9]+`,
				},
			},
		}
		assert.Equal(
			t,
			issue.Critical,
			s.Config().
				Task().
				ID().
				Pattern().
				Level(),
		)
	})

	t.Run("desc lvl does not affect task lvl on incorrect", func(t *testing.T) {
		t.Parallel()
		s := impl.Subject{
			Desc: tdescription.Description{
				Length: mvalue.Range{
					Issue: mvalue.Issue{
						Level: "warning",
					},
					Min: 1,
					Max: 72,
				},
			},
			Task: mtask.Task{
				ID: mvalue.Pattern{
					Issue: mvalue.Issue{
						Level: "foo",
					},
					Pattern: `(TASK|PROJ|BUG)-[0-9]+`,
				},
			},
		}
		assert.True(
			t,
			s.Config().
				Task().
				ID().
				Pattern().
				Level().
				Unknown(),
		)
	})
}
