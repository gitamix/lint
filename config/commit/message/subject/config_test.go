package subject_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_Length(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.NewConfig().Length()
		want := value.Range{}
		assert.Equal(t, want, got)
	})

	t.Run("with all options", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(10, 72),
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
			Length()
		want := value.NewRange(10, 72)
		assert.Equal(t, want, got)
	})

	t.Run("only with length", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(1, 100),
				),
			).
			Length()
		want := value.NewRange(1, 100)
		assert.Equal(t, want, got)
	})

	t.Run("with zero bounds", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(0, 0),
				),
			).
			Length()
		want := value.NewRange(0, 0)
		assert.Equal(t, want, got)
	})

	t.Run("last WithLength wins", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(1, 10),
				),
				impl.WithLength(
					value.NewRange(10, 72),
				),
			).
			Length()
		want := value.NewRange(10, 72)
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
				impl.WithLength(
					value.NewRange(10, 72),
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

	t.Run("only with length does not set task", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(10, 72),
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
