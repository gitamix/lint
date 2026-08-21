package description_test

import (
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/commit/message/subject/description"
	"github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestDescription_Issues(t *testing.T) {
	t.Parallel()

	t.Run("pass on defaults", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		d := impl.Description{}
		got := d.Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on defaults when description is set but config is not", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewDescription(
				commit.ParseSubject("feat: add new feature").Description(),
				description.Config{},
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass when length fits inside range", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewDescription(
				commit.ParseSubject("feat: add new feature").Description(),
				description.NewConfig(
					description.WithLength(
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
			NewDescription(
				commit.ParseSubject("feat: add new feature").Description(),
				description.NewConfig(
					description.WithLength(
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
			NewDescription(
				commit.ParseSubject("feat: add new feature").Description(),
				description.NewConfig(
					description.WithLength(
						value.NewRange(issue.Warning, 1, 15),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warning when description is too short", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("subject description length is not in range [10-72]"),
		}
		got := impl.
			NewDescription(
				commit.ParseSubject("hello").Description(),
				description.NewConfig(
					description.WithLength(
						value.NewRange(issue.Warning, 10, 72),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warning when description is too long", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("subject description length is not in range [1-10]"),
		}
		got := impl.
			NewDescription(
				commit.ParseSubject("feat: add new feature").Description(),
				description.NewConfig(
					description.WithLength(
						value.NewRange(issue.Warning, 1, 10),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("info when description is too short", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewInfo("subject description length is not in range [10-72]"),
		}
		got := impl.
			NewDescription(
				commit.ParseSubject("hello").Description(),
				description.NewConfig(
					description.WithLength(
						value.NewRange(issue.Info, 10, 72),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("critical when description is too long", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("subject description length is not in range [1-10]"),
		}
		got := impl.
			NewDescription(
				commit.ParseSubject("feat: add new feature").Description(),
				description.NewConfig(
					description.WithLength(
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
			NewDescription(
				commit.ParseSubject("feat: add new feature").Description(),
				description.NewConfig(
					description.WithLength(
						value.NewRange(issue.Warning, 0, 0),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on empty description with valid range", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("subject description length is not in range [1-72]"),
		}
		got := impl.
			NewDescription(
				commit.ParseSubject("").Description(),
				description.NewConfig(
					description.WithLength(
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
			NewDescription(
				commit.ParseSubject("feat: add new feature").Description(),
				description.NewConfig(
					description.WithLength(
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
			issue.NewWarning("subject description length is not in range [10-10]"),
		}
		got := impl.
			NewDescription(
				commit.ParseSubject("feat: add new feature").Description(),
				description.NewConfig(
					description.WithLength(
						value.NewRange(issue.Warning, 10, 10),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass when description excludes type and scope prefix", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewDescription(
				commit.ParseSubject("feat(ui): hello").Description(),
				description.NewConfig(
					description.WithLength(
						value.NewRange(issue.Warning, 1, 5),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn when description excludes type and scope prefix", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("subject description length is not in range [1-3]"),
		}
		got := impl.
			NewDescription(
				commit.ParseSubject("feat(ui): hello").Description(),
				description.NewConfig(
					description.WithLength(
						value.NewRange(issue.Warning, 1, 3),
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})
}
