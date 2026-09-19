package branch_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/defaults"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_Default(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig().
			Default()
		want := defaults.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("with default branch name", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("master"),
					),
				),
			).
			Default()
		want := defaults.NewConfig(
			defaults.WithName("master"),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with empty default branch name", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName(""),
					),
				),
			).
			Default()
		want := defaults.NewConfig(
			defaults.WithName(""),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with empty default config", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(),
				),
			).
			Default()
		want := defaults.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("only with name does not set default", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithName(
					name.NewConfig(
						value.NewString(
							issue.Warning,
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
			).
			Default()
		want := defaults.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("only with task does not set default", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTask(
					task.NewConfig(),
				),
			).
			Default()
		want := defaults.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("last WithDefault wins", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("master"),
					),
				),
				impl.WithDefault(
					defaults.NewConfig(
						defaults.WithName("main"),
					),
				),
			).
			Default()
		want := defaults.NewConfig(
			defaults.WithName("main"),
		)
		assert.Equal(t, want, got)
	})
}
