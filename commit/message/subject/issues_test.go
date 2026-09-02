package subject_test

import (
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/commit/message/subject"
	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/commit/types"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestSubject_Issues(t *testing.T) {
	t.Parallel()

	t.Run("pass on defaults", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		s := impl.Subject{}
		got := s.Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on defaults when subject is set but configs are not", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewSubject(
				commit.ParseSubject("feat(ui): add new feature"),
				subject.Config{},
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
			NewSubject(
				commit.ParseSubject("feat(ui): WS-1234 add new feature"),
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

	t.Run("aggregates warnings from all four sub-linters in order", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("type must be one of [feat,fix]"),
			issue.NewWarning("scope not found in subject by expression 'core|ui'"),
			issue.NewWarning("subject description length is not in range [10-72]"),
		}
		got := impl.
			NewSubject(
				commit.ParseSubject("refactor(db): hi"),
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

	t.Run("aggregates info from all four sub-linters in order", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewInfo("type must be one of [feat,fix]"),
			issue.NewInfo("scope not found in subject by expression 'core|ui'"),
			issue.NewInfo("subject description length is not in range [10-72]"),
		}
		got := impl.
			NewSubject(
				commit.ParseSubject("refactor(db): hi"),
				subject.NewConfig(
					subject.WithTask(
						task.NewConfig(
							task.WithID(
								id.NewConfig(
									value.NewString(
										issue.Info,
										"(WS[A-Z]*-[0-9]+)",
									),
								),
							),
						),
					),
					subject.WithDescription(
						description.NewConfig(
							description.WithLength(
								value.NewRange(issue.Info, 10, 72),
							),
						),
					),
				),
				types.NewConfig(
					value.NewStrings(
						issue.Info,
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

	t.Run("aggregates critical from all four sub-linters in order", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("type must be one of [feat,fix]"),
			issue.NewCritical("failed to compile scope expression '[': error parsing regexp: missing closing ]: `[`"),
			issue.NewCritical("subject description length is not in range [10-72]"),
		}
		got := impl.
			NewSubject(
				commit.ParseSubject("hi"),
				subject.NewConfig(
					subject.WithTask(
						task.NewConfig(
							task.WithID(
								id.NewConfig(
									value.NewString(
										issue.Warning,
										"[",
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
						"[",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("aggregates issues with mixed levels from all four sub-linters", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("type must be one of [feat,fix]"),
			issue.NewInfo("scope not found in subject by expression 'core|ui'"),
			issue.NewCritical("subject description length is not in range [10-72]"),
		}
		got := impl.
			NewSubject(
				commit.ParseSubject("refactor(db): hi"),
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

	t.Run("task not found but not issued", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewSubject(
				commit.ParseSubject("feat(ui): add new feature"),
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

	t.Run("types warning only", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("type must be one of [feat,fix]"),
		}
		got := impl.
			NewSubject(
				commit.ParseSubject("refactor(ui): WS-1234 add new feature"),
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

	t.Run("scope warning only", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("scope not found in subject by expression 'core|ui'"),
		}
		got := impl.
			NewSubject(
				commit.ParseSubject("feat(db): WS-1234 add new feature"),
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

	t.Run("description warning only", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("subject description length is not in range [11-72]"),
		}
		got := impl.
			NewSubject(
				commit.ParseSubject("feat(ui): WS-1234 hi"),
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
								value.NewRange(issue.Warning, 11, 72),
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

	t.Run("critical on empty type with configured types", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("type must be one of [feat,fix]"),
		}
		got := impl.
			NewSubject(
				commit.ParseSubject("add new feature"),
				subject.NewConfig(
					subject.WithDescription(
						description.NewConfig(
							description.WithLength(
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
				scope.Config{},
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("critical on invalid scope pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("failed to compile scope expression '[': error parsing regexp: missing closing ]: `[`"),
		}
		got := impl.
			NewSubject(
				commit.ParseSubject("feat(ui): add new feature"),
				subject.Config{},
				types.Config{},
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"[",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})
}
