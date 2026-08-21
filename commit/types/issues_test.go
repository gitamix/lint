package types_test

import (
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/commit/types"
	"github.com/gitamix/lint/config/commit/types"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestTypes_Issues(t *testing.T) {
	t.Parallel()

	t.Run("exists in list", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		got := impl.
			NewTypes(
				commit.NewType("feat"),
				types.NewConfig(
					value.NewStrings(
						issue.Critical,
						"feat",
						"fix",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("crit on empty type but warning provided", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewCritical("type must be one of [feat,fix]"),
		}
		got := impl.
			NewTypes(
				commit.NewType(""),
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("warn on type is not in list", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("type must be one of [feat,fix]"),
		}
		got := impl.
			NewTypes(
				commit.NewType("refactor"),
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("info on type is not in list", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewInfo("type must be one of [feat,fix]"),
		}
		got := impl.
			NewTypes(
				commit.NewType("refactor"),
				types.NewConfig(
					value.NewStrings(
						issue.Info,
						"feat",
						"fix",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})

	t.Run("pass on defaults", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{}
		typ := impl.Types{}
		got := typ.Issues()
		assert.Equal(t, want, got)
	})

	t.Run("incorrect case & warn on empty string in config", func(t *testing.T) {
		t.Parallel()
		want := []issue.Issue{
			issue.NewWarning("type must be one of []"),
		}
		got := impl.
			NewTypes(
				commit.NewType("feat"),
				types.NewConfig(
					value.NewStrings(
						issue.Warning,
						"",
					),
				),
			).
			Issues()
		assert.Equal(t, want, got)
	})
}
