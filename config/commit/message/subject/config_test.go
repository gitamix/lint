package subject_test

import (
	"testing"

	impl "github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
	"github.com/stretchr/testify/assert"
)

func TestConfig_Types(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.NewConfig().Types()
		want := impl.NewConfig().Types()
		assert.Equal(t, want, got)
	})

	t.Run("with types and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Critical,
						"feat",
						"fix",
					),
				),
			).
			Types()
		want := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Critical,
						"feat",
						"fix",
					),
				),
			).
			Types()
		assert.Equal(t, want, got)
	})

	t.Run("with empty types and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(issue.Critical),
				),
			).
			Types()
		want := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(issue.Critical),
				),
			).
			Types()
		assert.Equal(t, want, got)
	})

	t.Run("with empty types and warning level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(issue.Warning),
				),
			).
			Types()
		want := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(issue.Warning),
				),
			).
			Types()
		assert.Equal(t, want, got)
	})

	t.Run("with empty types and info level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(issue.Info),
				),
			).
			Types()
		want := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(issue.Info),
				),
			).
			Types()
		assert.Equal(t, want, got)
	})

	t.Run("with some empty type and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Critical,
						"foo",
						"",
						"bar",
					),
				),
			).
			Types()
		want := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Critical,
						"foo",
						"",
						"bar",
					),
				),
			).
			Types()
		assert.Equal(t, want, got)
	})

	t.Run("with some empty type and warning level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Warning,
						"foo",
						"",
						"bar",
					),
				),
			).
			Types()
		want := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Warning,
						"foo",
						"",
						"bar",
					),
				),
			).
			Types()
		assert.Equal(t, want, got)
	})

	t.Run("with some empty type and info level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Info,
						"foo",
						"",
						"bar",
					),
				),
			).
			Types()
		want := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Info,
						"foo",
						"",
						"bar",
					),
				),
			).
			Types()
		assert.Equal(t, want, got)
	})
}
