package scope_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_Pattern(t *testing.T) {
	t.Parallel()

	t.Run("pattern with warning issue type", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				value.NewString(
					issue.Warning,
					`^[A-Za-z _-]+$`,
				),
			).
			Pattern()
		want := value.NewString(
			issue.Warning,
			`^[A-Za-z _-]+$`,
		)
		assert.Equal(t, want, got)
	})

	t.Run("empty pattern with critical issue type", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				value.NewString(
					issue.Critical,
					"",
				),
			).
			Pattern()
		want := value.NewString(
			issue.Critical,
			"",
		)
		assert.Equal(t, want, got)
	})

	t.Run("default value", func(t *testing.T) {
		t.Parallel()
		var cfg impl.Config
		var want value.String
		assert.Equal(t, want, cfg.Pattern())
	})
}

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	t.Run("ok with zero pattern value", func(t *testing.T) {
		t.Parallel()
		var str value.String
		gotErr := impl.NewConfig(str).Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok with empty pattern value", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				value.NewString(issue.Critical, ""),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok with unspecfified issue level", func(t *testing.T) {
		t.Parallel()
		var lvl issue.Type
		gotErr := impl.
			NewConfig(
				value.NewString(lvl, "foo"),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok with valid pattern", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				value.NewString(
					issue.Warning,
					`^[a-z]+$`,
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("err with unknown issue level", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				value.NewString(
					issue.Type(-1),
					`^[a-z]+$`,
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
