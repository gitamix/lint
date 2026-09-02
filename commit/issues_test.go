package commit_test

import (
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/commit"
	config "github.com/gitamix/lint/config/commit"
	"github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/config/commit/message/body"
	"github.com/gitamix/lint/config/commit/message/body/mandate"
	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/commit/types"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestCommit_Issues(t *testing.T) {
	t.Parallel()

	t.Run("pass on defaults", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		c := impl.Commit{}
		got := c.Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on defaults when commit is set but config is not", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewCommit(
				commit.NewCommit(
					commit.NewHash("abc1234567"),
					commit.ParseMessage(
						[]byte("[WS-1234] feat(ui): add new feature\n\nbody text"),
					),
				),
				config.Config{},
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass when message and full config are valid", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewCommit(
				commit.NewCommit(
					commit.NewHash("abc1234567"),
					commit.ParseMessage(
						[]byte("[WS-1234] feat(ui): add new feature\n\nbody text"),
					),
				),
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

	t.Run("aggregates subject issues through config wiring", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
			issue.NewWarning("type must be one of [feat,fix]"),
			issue.NewWarning("scope not found in subject by expression 'core|ui'"),
			issue.NewWarning("subject description length is not in range [10-72]"),
		}
		got := impl.
			NewCommit(
				commit.NewCommit(
					commit.NewHash("abc1234567"),
					commit.ParseMessage(
						[]byte("refactor(db): hi\n\nbody text"),
					),
				),
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
														"(WS[A-Z]*-[0-9]+)",
													),
												),
											),
										),
									),
									subject.WithDescription(
										description.NewConfig(
											description.WithLength(
												value.NewRange(issue.Warning, 10, 72),
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
								"core|ui",
							),
						),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("aggregates body issue through config wiring", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("body is not in range [10-72]"),
		}
		got := impl.
			NewCommit(
				commit.NewCommit(
					commit.NewHash("abc1234567"),
					commit.ParseMessage(
						[]byte("[WS-1234] feat(ui): add new feature\n\nhi"),
					),
				),
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
										value.NewRange(issue.Warning, 10, 72),
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

	t.Run("critical when body is empty and mandated for type", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("body is required for type 'feat'"),
		}
		got := impl.
			NewCommit(
				commit.NewCommit(
					commit.NewHash("abc1234567"),
					commit.ParseMessage(
						[]byte("[WS-1234] feat(ui): add new feature"),
					),
				),
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
									body.WithMandate(
										mandate.NewConfig(
											mandate.WithTypes(
												value.NewStrings(
													issue.Warning,
													"feat",
												),
											),
										),
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

	t.Run("aggregates issues with mixed levels through config wiring", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
			issue.NewWarning("type must be one of [feat,fix]"),
			issue.NewInfo("scope not found in subject by expression 'core|ui'"),
			issue.NewCritical("subject description length is not in range [10-72]"),
			issue.NewCritical("body is not in range [10-72]"),
		}
		got := impl.
			NewCommit(
				commit.NewCommit(
					commit.NewHash("abc1234567"),
					commit.ParseMessage(
						[]byte("refactor(db): hi\n\nhi"),
					),
				),
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
														issue.Critical,
														"(WS[A-Z]*-[0-9]+)",
													),
												),
											),
										),
									),
									subject.WithDescription(
										description.NewConfig(
											description.WithLength(
												value.NewRange(issue.Critical, 10, 72),
											),
										),
									),
								),
							),
							message.WithBody(
								body.NewConfig(
									body.WithLength(
										value.NewRange(issue.Critical, 10, 72),
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
								issue.Info,
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
