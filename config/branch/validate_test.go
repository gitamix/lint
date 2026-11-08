package branch_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/defaults"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/errs"
	"github.com/gitamix/lint/issue"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	t.Run("err for empty config", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.NewConfig().Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid default branch: empty default branch name",
		)
	})

	t.Run("err for empty default config", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid default branch: empty default branch name",
		)
	})

	t.Run("err for blank default branch name", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("   "),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid default branch: empty default branch name",
		)
	})

	t.Run("err for empty default with valid name", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithName(
					name.NewConfig(
						value.NewString(
							issue.Warning,
							"",
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid default branch: empty default branch name",
		)
	})

	t.Run("err for default error even if name is invalid", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithName(
					name.NewConfig(
						value.NewString(
							issue.Type(-1),
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid default branch: empty default branch name",
		)
	})

	t.Run("err for last empty default overriding valid one", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("master"),
					),
				),
				impl.WithDefault(
					defaults.NewConfig(),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid default branch: empty default branch name",
		)
	})

	t.Run("ok for default name only", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("master"),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok for valid default & name", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("master"),
					),
				),
				impl.WithName(
					name.NewConfig(
						value.NewString(
							issue.Warning,
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok for valid default & task", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("master"),
					),
				),
				impl.WithTask(
					task.NewConfig(
						task.WithID(
							id.NewConfig(
								value.NewString(
									issue.Info,
									`[A-Z]+-\d+`,
								),
							),
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok for all valid", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("master"),
					),
				),
				impl.WithName(
					name.NewConfig(
						value.NewString(
							issue.Critical,
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
				impl.WithTask(
					task.NewConfig(
						task.WithID(
							id.NewConfig(
								value.NewString(
									issue.Info,
									`[A-Z]+-\d+`,
								),
							),
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok for zero level name & task", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("master"),
					),
				),
				impl.WithName(
					name.NewConfig(
						value.NewString(
							issue.Type(0),
							"",
						),
					),
				),
				impl.WithTask(
					task.NewConfig(
						task.WithID(
							id.NewConfig(
								value.NewString(
									issue.Type(0),
									"",
								),
							),
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok for last valid name overriding invalid one", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("master"),
					),
				),
				impl.WithName(
					name.NewConfig(
						value.NewString(
							issue.Type(-1),
							"",
						),
					),
				),
				impl.WithName(
					name.NewConfig(
						value.NewString(
							issue.Warning,
							"",
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("err for unknown name level with valid default", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("master"),
					),
				),
				impl.WithName(
					name.NewConfig(
						value.NewString(
							issue.Type(-1),
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid name: invalid config: unknown issue level",
		)
	})

	t.Run("err for unknown name level even if task is invalid", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("master"),
					),
				),
				impl.WithName(
					name.NewConfig(
						value.NewString(
							issue.Type(-1),
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
				impl.WithTask(
					task.NewConfig(
						task.WithID(
							id.NewConfig(
								value.NewString(
									issue.Type(4),
									`[A-Z]+-\d+`,
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid name: invalid config: unknown issue level",
		)
	})

	t.Run("err for unknown task level with valid default & name", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("master"),
					),
				),
				impl.WithName(
					name.NewConfig(
						value.NewString(
							issue.Warning,
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
				impl.WithTask(
					task.NewConfig(
						task.WithID(
							id.NewConfig(
								value.NewString(
									issue.Type(-1),
									`[A-Z]+-\d+`,
								),
							),
						),
					),
				),
			).
			Validate()
		assert.ErrorIs(t, gotErr, errs.ErrInvalidConfig)
	})

	t.Run("err for unknown task level with empty default", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithTask(
					task.NewConfig(
						task.WithID(
							id.NewConfig(
								value.NewString(
									issue.Type(-1),
									`[A-Z]+-\d+`,
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid default branch: empty default branch name",
		)
	})
}
