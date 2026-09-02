package body_test

import (
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/commit/message/body"
	"github.com/gitamix/lint/config/commit/message/body"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestBody_Issues(t *testing.T) {
	t.Parallel()

	t.Run("pass on defaults", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		b := impl.Body{}
		got := b.Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on defaults when body is set but config is not", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewBody(
				commit.NewBody([]byte("add new feature")),
				body.Config{},
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass when length fits inside range", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewBody(
				commit.NewBody([]byte("add new feature")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Warning, 1, 72),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass when length equals min boundary", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewBody(
				commit.NewBody([]byte("add new feature")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Warning, 15, 72),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass when length equals max boundary", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewBody(
				commit.NewBody([]byte("add new feature")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Warning, 1, 15),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warning when body is too short", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("body is not in range [10-72]"),
		}
		got := impl.
			NewBody(
				commit.NewBody([]byte("hello")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Warning, 10, 72),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warning when body is too long", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("body is not in range [1-10]"),
		}
		got := impl.
			NewBody(
				commit.NewBody([]byte("add new feature")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Warning, 1, 10),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("info when body is too short", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewInfo("body is not in range [10-72]"),
		}
		got := impl.
			NewBody(
				commit.NewBody([]byte("hello")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Info, 10, 72),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("critical when body is too long", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("body is not in range [1-10]"),
		}
		got := impl.
			NewBody(
				commit.NewBody([]byte("add new feature")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Critical, 1, 10),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on empty range with level set", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewBody(
				commit.NewBody([]byte("add new feature")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Warning, 0, 0),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on empty body with valid range", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("body is not in range [1-72]"),
		}
		got := impl.
			NewBody(
				commit.NewBody([]byte{}),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Warning, 1, 72),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on single-point range where length matches", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewBody(
				commit.NewBody([]byte("add new feature")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Warning, 15, 15),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on single-point range where length differs", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("body is not in range [10-10]"),
		}
		got := impl.
			NewBody(
				commit.NewBody([]byte("add new feature")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Warning, 10, 10),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on multi-line body matching range by rune count", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewBody(
				commit.NewBody([]byte("line1\nline2")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Warning, 11, 11),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on multi-line body too long by rune count", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("body is not in range [1-10]"),
		}
		got := impl.
			NewBody(
				commit.NewBody([]byte("line1\nline2")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Warning, 1, 10),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on cyrillic body matching range by rune count", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewBody(
				commit.NewBody([]byte("привет")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Warning, 6, 6),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on cyrillic body too long by rune count", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("body is not in range [1-5]"),
		}
		got := impl.
			NewBody(
				commit.NewBody([]byte("привет")),
				body.NewConfig(
					body.WithLength(
						value.NewRange(issue.Warning, 1, 5),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})
}
