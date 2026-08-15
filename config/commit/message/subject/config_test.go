package subject_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_Description(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.NewConfig().Description()
		want := description.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("with description", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithDescription(
					description.NewConfig(
						description.WithLength(
							value.NewRange(10, 72),
						),
					),
				),
			).
			Description()
		want := description.NewConfig(
			description.WithLength(
				value.NewRange(10, 72),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("last WithDescription wins", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithDescription(
					description.NewConfig(
						description.WithLength(
							value.NewRange(1, 10),
						),
					),
				),
				impl.WithDescription(
					description.NewConfig(
						description.WithLength(
							value.NewRange(10, 72),
						),
					),
				),
			).
			Description()
		want := description.NewConfig(
			description.WithLength(
				value.NewRange(10, 72),
			),
		)
		assert.Equal(t, want, got)
	})
}

func TestConfig_Task(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.NewConfig().Task()
		var want task.Config
		assert.Equal(t, want, got)
	})

	t.Run("with all options", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithDescription(
					description.NewConfig(
						description.WithLength(
							value.NewRange(10, 72),
						),
					),
				),
				impl.WithTask(
					task.NewConfig(
						task.WithID(
							id.NewConfig(
								value.NewString(
									issue.Warning,
									`^[A-Z]+-\d+$`,
								),
							),
						),
					),
				),
			).
			Task()
		want := task.NewConfig(
			task.WithID(
				id.NewConfig(
					value.NewString(
						issue.Warning,
						`^[A-Z]+-\d+$`,
					),
				),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("only with task", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTask(
					task.NewConfig(
						task.WithID(
							id.NewConfig(
								value.NewString(
									issue.Critical,
									`^[A-Z]+-\d+$`,
								),
							),
						),
					),
				),
			).
			Task()
		want := task.NewConfig(
			task.WithID(
				id.NewConfig(
					value.NewString(
						issue.Critical,
						`^[A-Z]+-\d+$`,
					),
				),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("only with description does not set task", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithDescription(
					description.NewConfig(
						description.WithLength(
							value.NewRange(10, 72),
						),
					),
				),
			).
			Task()
		var want task.Config
		assert.Equal(t, want, got)
	})

	t.Run("with empty task config", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTask(
					task.NewConfig(),
				),
			).
			Task()
		want := task.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("last WithTask wins", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTask(
					task.NewConfig(
						task.WithID(
							id.NewConfig(
								value.NewString(
									issue.Critical,
									`^[A-Z]+-\d+$`,
								),
							),
						),
					),
				),
				impl.WithTask(
					task.NewConfig(
						task.WithID(
							id.NewConfig(
								value.NewString(
									issue.Warning,
									`^[A-Z]+-\d+$`,
								),
							),
						),
					),
				),
			).
			Task()
		want := task.NewConfig(
			task.WithID(
				id.NewConfig(
					value.NewString(
						issue.Warning,
						`^[A-Z]+-\d+$`,
					),
				),
			),
		)
		assert.Equal(t, want, got)
	})
}
