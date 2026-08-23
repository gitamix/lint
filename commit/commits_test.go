package commit_test

import (
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/commit"
	config "github.com/gitamix/lint/config/commit"
	"github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/config/commit/message/body"
	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/commit/types"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestCommits_Issues(t *testing.T) {
	t.Parallel()

	t.Run("pass on defaults with no commits", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		c := impl.Commits{}
		got := c.Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass when all commits are valid", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewCommits(
				[]commit.Commit{
					commit.NewCommit(
						commit.NewHash("abc1234567"),
						commit.ParseMessage(
							[]byte("[WS-1234] feat(ui): add new feature\n\nbody text"),
						),
					),
					commit.NewCommit(
						commit.NewHash("def89abcdef"),
						commit.ParseMessage(
							[]byte("[WS-5678] fix(ui): resolve bug\n\nbody text"),
						),
					),
				},
				config.NewConfig(
					config.WithMessage(
						message.NewConfig(
							message.WithSubject(
								subject.NewConfig(
									subject.WithTask(
										task.NewConfig(
											task.WithID(
												id.NewConfig(
													value.NewString(
														issue.Warning,
														`^\[(WS[A-Z]*-[0-9]+)\]`,
													),
												),
											),
										),
									),
									subject.WithDescription(
										description.NewConfig(
											description.WithLength(
												value.NewRange(issue.Warning, 1, 72),
											),
										),
									),
								),
							),
							message.WithBody(
								body.NewConfig(
									body.WithLength(
										value.NewRange(issue.Warning, 1, 72),
									),
								),
							),
						),
					),
					config.WithTypes(
						types.NewConfig(
							value.NewStrings(
								issue.Warning,
								"feat",
								"fix",
							),
						),
					),
					config.WithScope(
						scope.NewConfig(
							value.NewString(
								issue.Warning,
								"ui",
							),
						),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("aggregates issues from a single invalid commit", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("abc1234: type must be one of [feat,fix]"),
		}
		got := impl.
			NewCommits(
				[]commit.Commit{
					commit.NewCommit(
						commit.NewHash("abc1234567"),
						commit.ParseMessage(
							[]byte("refactor(ui): add new feature\n\nbody text"),
						),
					),
				},
				config.NewConfig(
					config.WithTypes(
						types.NewConfig(
							value.NewStrings(
								issue.Warning,
								"feat",
								"fix",
							),
						),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("aggregates issues from multiple commits in order", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("abc1234: type must be one of [feat,fix]"),
			issue.NewWarning("def89ab: type must be one of [feat,fix]"),
		}
		got := impl.
			NewCommits(
				[]commit.Commit{
					commit.NewCommit(
						commit.NewHash("abc1234567"),
						commit.ParseMessage(
							[]byte("refactor(ui): add new feature\n\nbody text"),
						),
					),
					commit.NewCommit(
						commit.NewHash("def89abcdef"),
						commit.ParseMessage(
							[]byte("chore(ui): bump version\n\nbody text"),
						),
					),
				},
				config.NewConfig(
					config.WithTypes(
						types.NewConfig(
							value.NewStrings(
								issue.Warning,
								"feat",
								"fix",
							),
						),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("skips valid commits and reports only invalid ones", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("def89ab: type must be one of [feat,fix]"),
		}
		got := impl.
			NewCommits(
				[]commit.Commit{
					commit.NewCommit(
						commit.NewHash("abc1234567"),
						commit.ParseMessage(
							[]byte("feat(ui): add new feature\n\nbody text"),
						),
					),
					commit.NewCommit(
						commit.NewHash("def89abcdef"),
						commit.ParseMessage(
							[]byte("refactor(ui): add new feature\n\nbody text"),
						),
					),
				},
				config.NewConfig(
					config.WithTypes(
						types.NewConfig(
							value.NewStrings(
								issue.Warning,
								"feat",
								"fix",
							),
						),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("preserves order across commits with multiple issues each", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("abc1234: type must be one of [feat,fix]"),
			issue.NewWarning("abc1234: scope not found in subject by expression 'core|ui'"),
			issue.NewWarning("def89ab: type must be one of [feat,fix]"),
			issue.NewWarning("def89ab: scope not found in subject by expression 'core|ui'"),
		}
		got := impl.
			NewCommits(
				[]commit.Commit{
					commit.NewCommit(
						commit.NewHash("abc1234567"),
						commit.ParseMessage(
							[]byte("refactor(db): hi\n\nbody text"),
						),
					),
					commit.NewCommit(
						commit.NewHash("def89abcdef"),
						commit.ParseMessage(
							[]byte("chore(api): yo\n\nbody text"),
						),
					),
				},
				config.NewConfig(
					config.WithTypes(
						types.NewConfig(
							value.NewStrings(
								issue.Warning,
								"feat",
								"fix",
							),
						),
					),
					config.WithScope(
						scope.NewConfig(
							value.NewString(
								issue.Warning,
								"core|ui",
							),
						),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})
}
