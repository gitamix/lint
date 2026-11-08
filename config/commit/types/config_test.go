package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/commit/types"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_List(t *testing.T) {
	t.Parallel()

	t.Run("types with warning issue type", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				value.NewStrings(
					issue.Warning,
					"feat",
					"fix",
				),
			).
			List()
		want := value.NewStrings(
			issue.Warning,
			"feat",
			"fix",
		)
		assert.Equal(t, want, got)
	})

	t.Run("empty types with critical issue type", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				value.NewStrings(
					issue.Critical,
				),
			).
			List()
		want := value.NewStrings(
			issue.Critical,
		)
		assert.Equal(t, want, got)
	})

	t.Run("default value", func(t *testing.T) {
		t.Parallel()
		var cfg impl.Config
		var want value.Strings
		assert.Equal(t, want, cfg.List())
	})
}

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	t.Run("ok with zero types value", func(t *testing.T) {
		t.Parallel()
		var strs value.Strings
		gotErr := impl.NewConfig(strs).Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok with empty strings", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				value.NewStrings(issue.Critical),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok with unspecified issue level", func(t *testing.T) {
		t.Parallel()
		var lvl issue.Type
		gotErr := impl.
			NewConfig(
				value.NewStrings(
					lvl,
					"feat",
					"fix",
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok with valid types", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				value.NewStrings(
					issue.Warning,
					"feat",
					"fix",
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("err with unknown issue level", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				value.NewStrings(
					issue.Type(-1),
					"feat",
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid config: unknown issue level",
		)
	})
}
