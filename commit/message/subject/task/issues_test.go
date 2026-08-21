package task_test

import (
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/commit/message/subject/task"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestTask_Issues(t *testing.T) {
	t.Parallel()

	t.Run("ticket matches WS with extra letters", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewTask(
				commit.ParseSubject("[WSTS-1234] some fix"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("ticket matches WS without extra letters", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewTask(
				commit.ParseSubject("[WS-1234] some fix"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("ticket matches WS with many extra letters", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewTask(
				commit.ParseSubject("[WSPROJ-99] fix"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warning on ticket not matching pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
		}
		got := impl.
			NewTask(
				commit.ParseSubject("add new feature"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("info on ticket not matching pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewInfo("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
		}
		got := impl.
			NewTask(
				commit.ParseSubject("add new feature"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("critical on ticket not matching pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
		}
		got := impl.
			NewTask(
				commit.ParseSubject("add new feature"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("critical on invalid pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("failed to parse task id with expression '[': error parsing regexp: missing closing ]: `[`"),
		}
		got := impl.
			NewTask(
				commit.ParseSubject("[WSTS-1234] some fix"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on defaults", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		tk := impl.Task{}
		got := tk.Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on defaults when subject is set but config is not", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewTask(
				commit.ParseSubject("[WSTS-1234] some fix"),
				task.Config{},
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on empty pattern with level set", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewTask(
				commit.ParseSubject("[WSTS-1234] some fix"),
				task.NewConfig(
					task.WithID(
						id.NewConfig(
							value.NewString(
								issue.Warning,
								"",
							),
						),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on empty subject with valid pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
		}
		got := impl.
			NewTask(
				commit.ParseSubject(""),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on subject with type but no ticket", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
		}
		got := impl.
			NewTask(
				commit.ParseSubject("feat: add new feature"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on prefix missing WS", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
		}
		got := impl.
			NewTask(
				commit.ParseSubject("[ABTS-1234] some fix"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on ticket without trailing digits", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
		}
		got := impl.
			NewTask(
				commit.ParseSubject("[WSTS] some fix"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("critical on lowercased ticket not matching pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("ticket not found in subject by expression '(WS[A-Z]*-[0-9]+)'"),
		}
		got := impl.
			NewTask(
				commit.ParseSubject("[wsts-1234] some fix"),
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
			).
			Issues()
		assert.Equal(t, want, got)
	})
}
