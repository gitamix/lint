package current_test

import (
	"context"
	"testing"

	"github.com/gitamix/types/branch"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/branch/current"
	config "github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/errs"
	"github.com/gitamix/lint/internal/test/fake/repo/git"
	"github.com/gitamix/lint/issue"
)

func TestBranch_Issues(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		b := impl.NewBranch(
			git.NewRepository(
				git.WithCurrentBranch(
					func(_ context.Context) (branch.Branch, error) {
						return branch.NewBranch(
							branch.NewName("feature/TASK-123"),
						), nil
					},
				),
			),
			config.NewConfig(
				config.WithName(
					name.NewConfig(
						value.NewString(
							issue.Critical,
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
				config.WithTask(
					task.NewConfig(
						task.WithID(
							id.NewConfig(
								value.NewString(
									issue.Warning,
									`(TASK|PROJ|BUG)-[0-9]+`,
								),
							),
						),
					),
				),
			),
		)
		got, gotErr := b.Issues(context.Background())
		assert.Equal(
			t,
			[]issue.Issue{},
			got,
		)
		assert.NoError(t, gotErr)
	})

	t.Run("failed to get current branch from git", func(t *testing.T) {
		t.Parallel()
		b := impl.NewBranch(
			git.NewRepository(
				git.WithCurrentBranch(
					func(_ context.Context) (branch.Branch, error) {
						return branch.Branch{}, assert.AnError
					},
				),
			),
			config.NewConfig(
				config.WithName(
					name.NewConfig(
						value.NewString(
							issue.Critical,
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
			),
		)
		got, gotErr := b.Issues(context.Background())
		assert.Equal(
			t,
			[]issue.Issue(nil),
			got,
		)
		assert.ErrorIs(t, gotErr, errs.ErrGitFailed)
	})

	t.Run("critical on name not matched pattern", func(t *testing.T) {
		t.Parallel()
		b := impl.NewBranch(
			git.NewRepository(
				git.WithCurrentBranch(
					func(_ context.Context) (branch.Branch, error) {
						return branch.NewBranch(
							branch.NewName("release/TASK-123"),
						), nil
					},
				),
			),
			config.NewConfig(
				config.WithName(
					name.NewConfig(
						value.NewString(
							issue.Critical,
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
			),
		)
		got, gotErr := b.Issues(context.Background())
		assert.Equal(
			t,
			[]issue.Issue{
				issue.NewCritical(
					`branch name 'release/TASK-123' doesn't match the required pattern '^(feature|bugfix|hotfix)/[A-Z]+-\d+'`,
				),
			},
			got,
		)
		assert.NoError(t, gotErr)
	})

	t.Run("criticals on name & task not match pattern", func(t *testing.T) {
		t.Parallel()
		b := impl.NewBranch(
			git.NewRepository(
				git.WithCurrentBranch(
					func(_ context.Context) (branch.Branch, error) {
						return branch.NewBranch(
							branch.NewName("my-favorite-feature"),
						), nil
					},
				),
			),
			config.NewConfig(
				config.WithName(
					name.NewConfig(
						value.NewString(
							issue.Critical,
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
				config.WithTask(
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
			),
		)
		got, gotErr := b.Issues(context.Background())
		assert.Equal(
			t,
			[]issue.Issue{
				issue.NewCritical(
					`branch name 'my-favorite-feature' doesn't match the required pattern '^(feature|bugfix|hotfix)/[A-Z]+-\d+'`,
				),
				issue.NewCritical(`task doesn't match the required pattern '(TASK|PROJ|BUG)-[0-9]+'`),
			},
			got,
		)
		assert.NoError(t, gotErr)
	})

	t.Run("warning on empty branch name", func(t *testing.T) {
		t.Parallel()
		b := impl.NewBranch(
			git.NewRepository(
				git.WithCurrentBranch(
					func(_ context.Context) (branch.Branch, error) {
						return branch.NewBranch(
							branch.NewName(""),
						), nil
					},
				),
			),
			config.NewConfig(
				config.WithName(
					name.NewConfig(
						value.NewString(
							issue.Warning,
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
			),
		)
		got, gotErr := b.Issues(context.Background())
		assert.Equal(
			t,
			[]issue.Issue{
				issue.NewWarning(`branch name '' doesn't match the required pattern '^(feature|bugfix|hotfix)/[A-Z]+-\d+'`),
			},
			got,
		)
		assert.NoError(t, gotErr)
	})

	t.Run("empty pattern", func(t *testing.T) {
		t.Parallel()
		b := impl.NewBranch(
			git.NewRepository(
				git.WithCurrentBranch(
					func(_ context.Context) (branch.Branch, error) {
						return branch.NewBranch(
							branch.NewName("feature/TASK-123"),
						), nil
					},
				),
			),
			config.NewConfig(
				config.WithName(
					name.NewConfig(
						value.NewString(
							issue.Warning,
							``,
						),
					),
				),
			),
		)
		got, gotErr := b.Issues(context.Background())
		assert.Equal(
			t,
			[]issue.Issue{},
			got,
		)
		assert.NoError(t, gotErr)
	})

	t.Run("empty task id pattern & name matched pattern", func(t *testing.T) {
		t.Parallel()
		b := impl.NewBranch(
			git.NewRepository(
				git.WithCurrentBranch(
					func(_ context.Context) (branch.Branch, error) {
						return branch.NewBranch(
							branch.NewName("feature/TASK-123"),
						), nil
					},
				),
			),
			config.NewConfig(
				config.WithName(
					name.NewConfig(
						value.NewString(
							issue.Warning,
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
				config.WithTask(
					task.NewConfig(
						task.WithID(
							id.NewConfig(
								value.NewString(
									issue.Critical,
									``,
								),
							),
						),
					),
				),
			),
		)
		got, gotErr := b.Issues(context.Background())
		assert.Equal(
			t,
			[]issue.Issue{},
			got,
		)
		assert.NoError(t, gotErr)
	})

	t.Run("empty current branch name and pattern", func(t *testing.T) {
		t.Parallel()
		b := impl.NewBranch(
			git.NewRepository(
				git.WithCurrentBranch(
					func(_ context.Context) (branch.Branch, error) {
						return branch.NewBranch(
							branch.NewName(""),
						), nil
					},
				),
			),
			config.NewConfig(
				config.WithName(
					name.NewConfig(
						value.NewString(
							issue.Warning,
							``,
						),
					),
				),
			),
		)
		got, gotErr := b.Issues(context.Background())
		assert.Equal(
			t,
			[]issue.Issue{},
			got,
		)
		assert.NoError(t, gotErr)
	})

	t.Run("empty config", func(t *testing.T) {
		t.Parallel()
		b := impl.NewBranch(
			git.NewRepository(
				git.WithCurrentBranch(
					func(_ context.Context) (branch.Branch, error) {
						return branch.NewBranch(
							branch.NewName("feature/TASK-123"),
						), nil
					},
				),
			),
			config.NewConfig(),
		)
		got, gotErr := b.Issues(context.Background())
		assert.Equal(
			t,
			[]issue.Issue{},
			got,
		)
		assert.NoError(t, gotErr)
	})

	t.Run("empty value panics", func(t *testing.T) {
		t.Parallel()
		b := &impl.Branch{}
		assert.Panics(t, func() {
			_, _ = b.Issues(context.Background())
		})
	})

	t.Run("nil value panics", func(t *testing.T) {
		t.Parallel()
		var b *impl.Branch
		assert.Panics(t, func() {
			_, _ = b.Issues(context.Background())
		})
	})
}
