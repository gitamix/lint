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

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	t.Run("ok if id pattern is valid & info issue", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithID(
					id.NewConfig(
						value.NewString(
							issue.Info,
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("err if id pattern is valid but type is unknown", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithID(
					id.NewConfig(
						value.NewString(
							issue.Type(-1),
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
			).
			Validate()
		assert.EqualError(t, gotErr, "invalid config: unknown issue level")
	})

	t.Run("ok if id pattern is valid but type is zero as default", func(t *testing.T) {
		t.Parallel()
		var typ issue.Type
		gotErr := impl.
			NewConfig(
				impl.WithID(
					id.NewConfig(
						value.NewString(
							typ,
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok if id pattern is valid but not trimmed & info issue", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithID(
					id.NewConfig(
						value.NewString(
							issue.Info,
							` ^(feature|bugfix|hotfix)/[A-Z]+-\d+ `,
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok if id pattern is just a string & warning issue", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithID(
					id.NewConfig(
						value.NewString(
							issue.Warning,
							`foo`,
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok if id pattern is empty & warning issue", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithID(
					id.NewConfig(
						value.NewString(
							issue.Warning,
							"",
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok if id pattern is empty & critical issue", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithID(
					id.NewConfig(
						value.NewString(
							issue.Critical,
							"",
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok on default value", func(t *testing.T) {
		t.Parallel()
		var conf impl.Config
		gotErr := conf.Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok on default id config value", func(t *testing.T) {
		t.Parallel()
		var id id.Config
		gotErr := impl.
			NewConfig(
				impl.WithID(id),
			).
			Validate()
		assert.NoError(t, gotErr)
	})
}
