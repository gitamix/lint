package id_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_Pattern(t *testing.T) {
	t.Parallel()

	t.Run("regexp", func(t *testing.T) {
		t.Parallel()
		want := value.NewString(
			issue.Warning,
			`^[A-Z]+-\d+$`,
		)
		assert.Equal(
			t,
			want,
			impl.NewConfig(want).Pattern(),
		)
	})

	t.Run("critical level & just a word", func(t *testing.T) {
		t.Parallel()
		want := value.NewString(
			issue.Critical,
			"foo",
		)
		assert.Equal(
			t,
			want,
			impl.NewConfig(want).Pattern(),
		)
	})

	t.Run("warning level & just a word", func(t *testing.T) {
		t.Parallel()
		want := value.NewString(
			issue.Warning,
			"foo",
		)
		assert.Equal(
			t,
			want,
			impl.NewConfig(want).Pattern(),
		)
	})

	t.Run("info level & just a word", func(t *testing.T) {
		t.Parallel()
		want := value.NewString(
			issue.Info,
			"foo",
		)
		assert.Equal(
			t,
			want,
			impl.NewConfig(want).Pattern(),
		)
	})

	t.Run("critical level & empty string", func(t *testing.T) {
		t.Parallel()
		want := value.NewString(
			issue.Critical,
			"",
		)
		assert.Equal(
			t,
			want,
			impl.NewConfig(want).Pattern(),
		)
	})

	t.Run("warning level & empty string", func(t *testing.T) {
		t.Parallel()
		want := value.NewString(
			issue.Warning,
			"",
		)
		assert.Equal(
			t,
			want,
			impl.NewConfig(want).Pattern(),
		)
	})

	t.Run("info level & empty string", func(t *testing.T) {
		t.Parallel()
		want := value.NewString(
			issue.Info,
			"",
		)
		assert.Equal(
			t,
			want,
			impl.NewConfig(want).Pattern(),
		)
	})
}

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	t.Run("ok if pattern is valid & info issue", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				value.NewString(
					issue.Info,
					`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("err if pattern is valid but type is unknown", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				value.NewString(
					issue.Type(-1),
					`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
				),
			).
			Validate()
		assert.EqualError(t, gotErr, "invalid config: unknown issue level")
	})

	t.Run("ok if pattern is valid but type is zero as default", func(t *testing.T) {
		t.Parallel()
		var typ issue.Type
		gotErr := impl.
			NewConfig(
				value.NewString(
					typ,
					`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok if pattern is valid but not trimmed & info issue", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				value.NewString(
					issue.Info,
					` ^(feature|bugfix|hotfix)/[A-Z]+-\d+ `,
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok if pattern is just a string & warning issue", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				value.NewString(
					issue.Warning,
					`foo`,
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok if pattern is empty & warning issue", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				value.NewString(
					issue.Warning,
					"",
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok if pattern is empty & critical issue", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				value.NewString(
					issue.Critical,
					"",
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})
}
