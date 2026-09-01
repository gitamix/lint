package branch_test

import (
	"testing"

	"github.com/gitamix/types/branch"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/branch"
	config "github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestBranch_Issues(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewBranch(
				branch.NewBranch(
					branch.NewName("feature/TASK-123"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("critical on name not matched pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical(
				`branch name 'release/TASK-123' doesn't match the required pattern '^(feature|bugfix|hotfix)/[A-Z]+-\d+'`,
			),
		}
		got := impl.
			NewBranch(
				branch.NewBranch(
					branch.NewName("release/TASK-123"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("criticals on name & task not match pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical(
				`branch name 'my-favorite-feature' doesn't match the required pattern '^(feature|bugfix|hotfix)/[A-Z]+-\d+'`,
			),
			issue.NewCritical(`task doesn't match the required pattern '(TASK|PROJ|BUG)-[0-9]+'`),
		}
		got := impl.
			NewBranch(
				branch.NewBranch(
					branch.NewName("my-favorite-feature"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warning on empty branch name", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning(`branch name '' doesn't match the required pattern '^(feature|bugfix|hotfix)/[A-Z]+-\d+'`),
		}
		got := impl.
			NewBranch(
				branch.NewBranch(
					branch.NewName(""),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("empty pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewBranch(
				branch.NewBranch(
					branch.NewName("feature/TASK-123"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("empty task id pattern & name matched pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewBranch(
				branch.NewBranch(
					branch.NewName("feature/TASK-123"),
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
			).Issues()
		assert.Equal(t, want, got)
	})

	t.Run("empty current branch name and pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		b := impl.NewBranch(
			branch.NewBranch(
				branch.NewName(""),
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
		got := b.Issues()
		assert.Equal(t, want, got)
	})

	t.Run("empty config", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewBranch(
				branch.NewBranch(
					branch.NewName("feature/TASK-123"),
				),
				config.NewConfig(),
			).
			Issues()
		assert.Equal(t, want, got)
	})
}
