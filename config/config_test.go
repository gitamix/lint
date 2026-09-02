package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config"
	"github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/commit"
	"github.com/gitamix/lint/config/commit/types"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_Branch(t *testing.T) {
	t.Parallel()

	t.Run("with branch name & task configs", func(t *testing.T) {
		t.Parallel()
		want := branch.NewConfig(
			branch.WithName(
				name.NewConfig(
					value.NewString(
						issue.Warning,
						`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
					),
				),
			),
			branch.WithTask(
				task.NewConfig(
					task.WithID(
						id.NewConfig(
							value.NewString(
								issue.Critical,
								`[A-Z]+-\d+`,
							),
						),
					),
				),
			),
		)
		assert.Equal(
			t,
			want,
			impl.
				NewConfig(impl.WithBranch(want)).
				Branch(),
		)
	})

	t.Run("with name config only", func(t *testing.T) {
		t.Parallel()
		want := branch.NewConfig(
			branch.WithName(
				name.NewConfig(
					value.NewString(
						issue.Critical,
						"foo",
					),
				),
			),
		)
		assert.Equal(
			t,
			want,
			impl.
				NewConfig(impl.WithBranch(want)).
				Branch(),
		)
	})

	t.Run("with task config only", func(t *testing.T) {
		t.Parallel()
		want := branch.NewConfig(
			branch.WithTask(
				task.NewConfig(
					task.WithID(
						id.NewConfig(
							value.NewString(
								issue.Info,
								"",
							),
						),
					),
				),
			),
		)
		assert.Equal(
			t,
			want,
			impl.
				NewConfig(impl.WithBranch(want)).
				Branch(),
		)
	})

	t.Run("without any options", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			branch.NewConfig(),
			impl.NewConfig().Branch(),
		)
	})

	t.Run("default value panics", func(t *testing.T) {
		t.Parallel()
		var c *impl.Config
		assert.Panics(t, func() {
			_ = c.Branch()
		})
	})
}

func TestConfig_Commit(t *testing.T) {
	t.Parallel()

	t.Run("returns filled out commit config as it defined", func(t *testing.T) {
		t.Parallel()
		cfg := commit.NewConfig(
			commit.WithTypes(
				types.NewConfig(
					value.NewStrings(
						issue.Critical,
						"feat",
						"refactor",
						"fix",
						"chore",
						"docs",
						"test",
					),
				),
			),
		)
		assert.Equal(
			t,
			cfg,
			impl.
				NewConfig(impl.WithCommit(cfg)).
				Commit(),
		)
	})

	t.Run("returns empty commit config as it defined", func(t *testing.T) {
		t.Parallel()
		cfg := commit.NewConfig()
		assert.Equal(
			t,
			cfg,
			impl.
				NewConfig(impl.WithCommit(cfg)).
				Commit(),
		)
	})

	t.Run("panics on nil value", func(t *testing.T) {
		t.Parallel()
		var cfg *impl.Config
		assert.Panics(t, func() {
			_ = cfg.Commit()
		})
	})
}
