package scope_test

import (
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/commit/scope"
	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestScope_Issues(t *testing.T) {
	t.Parallel()

	t.Run("scope matches pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewScope(
				commit.NewScope("core"),
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

	t.Run("warning on scope not matching pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("scope not found in subject by expression \\core|ui\\"),
		}
		got := impl.
			NewScope(
				commit.NewScope("db"),
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

	t.Run("info on scope not matching pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewInfo("scope not found in subject by expression \\core|ui\\"),
		}
		got := impl.
			NewScope(
				commit.NewScope("db"),
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

	t.Run("critical on scope not matching pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("scope not found in subject by expression \\core|ui\\"),
		}
		got := impl.
			NewScope(
				commit.NewScope("db"),
				scope.NewConfig(
					value.NewString(
						issue.Critical,
						"core|ui",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("critical on invalid pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("failed to compile scope expression \\[\\: error parsing regexp: missing closing ]: `[`"),
		}
		got := impl.
			NewScope(
				commit.NewScope("core"),
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

	t.Run("pass on defaults", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		scp := impl.Scope{}
		got := scp.Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on defaults when scope is set but config is not", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewScope(
				commit.NewScope("core"),
				scope.Config{},
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on empty pattern with level set", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewScope(
				commit.NewScope("core"),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on empty scope with valid pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("scope not found in subject by expression \\core|ui\\"),
		}
		got := impl.
			NewScope(
				commit.NewScope(""),
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

	t.Run("alphanumeric pattern matches letters and allowed special chars", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewScope(
				commit.NewScope("my_module-ui"),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"^[A-Za-z _-]+$",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on digits not matching alphanumeric pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("scope not found in subject by expression \\^[A-Za-z _-]+$\\"),
		}
		got := impl.
			NewScope(
				commit.NewScope("db123"),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"^[A-Za-z _-]+$",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on non-ascii not matching alphanumeric pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("scope not found in subject by expression \\^[A-Za-z _-]+$\\"),
		}
		got := impl.
			NewScope(
				commit.NewScope("café"),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"^[A-Za-z _-]+$",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on empty scope not matching alphanumeric pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("scope not found in subject by expression \\^[A-Za-z _-]+$\\"),
		}
		got := impl.
			NewScope(
				commit.NewScope(""),
				scope.NewConfig(
					value.NewString(
						issue.Warning,
						"^[A-Za-z _-]+$",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("info on digits not matching alphanumeric pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewInfo("scope not found in subject by expression \\^[A-Za-z _-]+$\\"),
		}
		got := impl.
			NewScope(
				commit.NewScope("db123"),
				scope.NewConfig(
					value.NewString(
						issue.Info,
						"^[A-Za-z _-]+$",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("critical on non-ascii not matching alphanumeric pattern", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("scope not found in subject by expression \\^[A-Za-z _-]+$\\"),
		}
		got := impl.
			NewScope(
				commit.NewScope("café"),
				scope.NewConfig(
					value.NewString(
						issue.Critical,
						"^[A-Za-z _-]+$",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})
}
