package mandate_test

import (
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/commit/message/body/mandate"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_For(t *testing.T) {
	t.Parallel()

	t.Run("true if exists by type", func(t *testing.T) {
		t.Parallel()
		assert.True(
			t,
			impl.
				NewConfig(
					impl.WithTypes(
						value.NewStrings(
							issue.Warning,
							"fix",
							"perf",
						),
					),
				).
				For(commit.NewType("fix")),
		)
	})

	t.Run("true if exists by type as string", func(t *testing.T) {
		t.Parallel()
		assert.True(
			t,
			impl.
				NewConfig(
					impl.WithTypes(
						value.NewStrings(
							issue.Warning,
							"fix",
							"perf",
						),
					),
				).
				For("fix"),
		)
	})

	t.Run("true if exists by second type as it declared", func(t *testing.T) {
		t.Parallel()
		assert.True(
			t,
			impl.
				NewConfig(
					impl.WithTypes(
						value.NewStrings(
							issue.Warning,
							"fix",
							"perf",
						),
					),
				).
				For(commit.NewType("perf")),
		)
	})

	t.Run("false if not exists by type", func(t *testing.T) {
		t.Parallel()
		assert.False(
			t,
			impl.
				NewConfig(
					impl.WithTypes(
						value.NewStrings(
							issue.Warning,
							"fix",
							"perf",
						),
					),
				).
				For(commit.NewType("refactor")),
		)
	})

	t.Run("false if config is empty", func(t *testing.T) {
		t.Parallel()
		assert.False(
			t,
			impl.
				NewConfig().
				For(commit.NewType("refactor")),
		)
	})

	t.Run("false if empty type provided", func(t *testing.T) {
		t.Parallel()
		assert.False(
			t,
			impl.
				NewConfig(
					impl.WithTypes(
						value.NewStrings(
							issue.Warning,
							"fix",
							"perf",
						),
					),
				).
				For(commit.NewType("")),
		)
	})

	t.Run("false if empty string provided", func(t *testing.T) {
		t.Parallel()
		assert.False(
			t,
			impl.
				NewConfig(
					impl.WithTypes(
						value.NewStrings(
							issue.Warning,
							"fix",
							"perf",
						),
					),
				).
				For(""),
		)
	})

	t.Run("false on defaults", func(t *testing.T) {
		t.Parallel()
		var cfg impl.Config
		var typ commit.Type
		assert.False(t, cfg.For(typ))
	})
}
