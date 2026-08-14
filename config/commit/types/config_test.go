package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/commit/types"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_Types(t *testing.T) {
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
			Types()
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
			Types()
		want := value.NewStrings(
			issue.Critical,
		)
		assert.Equal(t, want, got)
	})

	t.Run("default value", func(t *testing.T) {
		t.Parallel()
		var cfg impl.Config
		var want value.Strings
		assert.Equal(t, want, cfg.Types())
	})
}
