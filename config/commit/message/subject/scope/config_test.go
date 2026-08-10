package scope_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/commit/message/subject/scope"
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
