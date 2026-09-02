package task_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_ID(t *testing.T) {
	t.Parallel()

	t.Run("with id config & correct pattern", func(t *testing.T) {
		t.Parallel()
		want := id.NewConfig(
			value.NewString(
				issue.Warning,
				`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
			),
		)
		assert.Equal(
			t,
			want,
			impl.
				NewConfig(impl.WithID(want)).
				ID(),
		)
	})

	t.Run("with id config & just a word as pattern", func(t *testing.T) {
		t.Parallel()
		want := id.NewConfig(
			value.NewString(
				issue.Critical,
				"foo",
			),
		)
		assert.Equal(
			t,
			want,
			impl.
				NewConfig(impl.WithID(want)).
				ID(),
		)
	})

	t.Run("with id config & empty pattern", func(t *testing.T) {
		t.Parallel()
		want := id.NewConfig(
			value.NewString(
				issue.Info,
				"",
			),
		)
		assert.Equal(
			t,
			want,
			impl.
				NewConfig(impl.WithID(want)).
				ID(),
		)
	})

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			id.Config{},
			impl.NewConfig().ID(),
		)
	})
}
