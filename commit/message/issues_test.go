package message_test

import (
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/commit/message"
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

func TestMessage_Issues(t *testing.T) {
	t.Parallel()

	t.Run("pass on defaults", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		m := impl.Message{}
		got := m.Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on defaults when message is set but configs are not", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewMessage(
				commit.ParseMessage(
					[]byte("[WS-1234] feat(ui): add new feature\n\nbody text"),
				),
				message.Config{},
				types.Config{},
				scope.Config{},
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass when all sub-linters are valid", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewMessage(
				commit.ParseMessage(
					[]byte("[WS-1234] feat(ui): add new feature\n\nbody text"),
				),
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
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"ui",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("aggregates subject issues when body is valid", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
			issue.NewWarning("type must be one of [feat,fix]"),
			issue.NewWarning("scope not found in subject by expression 'core|ui'"),
			issue.NewWarning("subject description length is not in range [10-72]"),
		}
		got := impl.
			NewMessage(
				commit.ParseMessage(
					[]byte("refactor(db): hi\n\nbody text"),
				),
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
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"core|ui",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("aggregates body issue when subject is valid", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("body is not in range [10-72]"),
		}
		got := impl.
			NewMessage(
				commit.ParseMessage(
					[]byte("[WS-1234] feat(ui): add new feature\n\nhi"),
				),
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
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"ui",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("aggregates subject and body issues in order", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
			issue.NewWarning("type must be one of [feat,fix]"),
			issue.NewWarning("scope not found in subject by expression 'core|ui'"),
			issue.NewWarning("subject description length is not in range [10-72]"),
			issue.NewWarning("body is not in range [10-72]"),
		}
		got := impl.
			NewMessage(
				commit.ParseMessage(
					[]byte("refactor(db): hi\n\nhi"),
				),
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
								value.NewRange(issue.Warning, 10, 72),
							),
						),
					),
				),
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"core|ui",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("aggregates issues with mixed levels from subject and body", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
			issue.NewWarning("type must be one of [feat,fix]"),
			issue.NewInfo("scope not found in subject by expression 'core|ui'"),
			issue.NewCritical("subject description length is not in range [10-72]"),
			issue.NewCritical("body is not in range [10-72]"),
		}
		got := impl.
			NewMessage(
				commit.ParseMessage(
					[]byte("refactor(db): hi\n\nhi"),
				),
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
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				scope.NewConfig(
					value.NewString(
						issue.Info,
						"core|ui",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("body issue reflects body config level info", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewInfo("body is not in range [10-72]"),
		}
		got := impl.
			NewMessage(
				commit.ParseMessage(
					[]byte("[WS-1234] feat(ui): add new feature\n\nhi"),
				),
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
								value.NewRange(issue.Info, 10, 72),
							),
						),
					),
				),
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"ui",
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
			NewMessage(
				commit.ParseMessage(
					[]byte("[WS-1234] feat(ui): add new feature"),
				),
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
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"ui",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("critical message uses actual commit type", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("body is required for type 'fix'"),
		}
		got := impl.
			NewMessage(
				commit.ParseMessage(
					[]byte("fix(ui): WS-1234 resolve bug"),
				),
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
											"fix",
										),
									),
								),
							),
						),
					),
				),
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"ui",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("no critical when body is empty and type not in mandate", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewMessage(
				commit.ParseMessage(
					[]byte("[WS-1234] feat(ui): add new feature"),
				),
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
											"fix",
										),
									),
								),
							),
						),
					),
				),
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"ui",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("whitespace-only body treated as empty with mandate", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("body is required for type 'feat'"),
		}
		got := impl.
			NewMessage(
				commit.ParseMessage(
					[]byte("[WS-1234] feat(ui): add new feature\n\n   "),
				),
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
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"ui",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("subject issue and mandated body critical preserve order", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
			issue.NewCritical("body is required for type 'refactor'"),
		}
		got := impl.
			NewMessage(
				commit.ParseMessage(
					[]byte("refactor(ui): add new feature"),
				),
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
											"refactor",
										),
									),
								),
							),
						),
					),
				),
				types.Config{},
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"ui",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("ticket after type fails when id pattern requires ticket at start", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning(`ticket not found in subject by expression '^\[(WS[A-Z]*-[0-9]+)\]'`),
		}
		got := impl.
			NewMessage(
				commit.ParseMessage(
					[]byte("feat(ui): WS-1234 add new feature"),
				),
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
				),
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"ui",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("ticket at start passes when id pattern requires ticket at start", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewMessage(
				commit.ParseMessage(
					[]byte("[WSTS-1234] feat(ui): add new feature"),
				),
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
				),
				types.Config{},
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"ui",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})
}
