package name_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_Pattern(t *testing.T) {
	t.Parallel()

	t.Run("regexp", func(t *testing.T) {
		t.Parallel()
		want := value.NewString(
			issue.Warning,
			`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
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

func TestConfig_Level(t *testing.T) {
	t.Parallel()

	t.Run("critical level", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			issue.Critical,
			impl.
				NewConfig(
					value.NewString(
						issue.Critical,
						"foo",
					),
				).
				Pattern().
				Level(),
		)
	})

	t.Run("warning level", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			issue.Warning,
			impl.
				NewConfig(
					value.NewString(
						issue.Warning,
						"foo",
					),
				).
				Pattern().
				Level(),
		)
	})

	t.Run("info level", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			issue.Info,
			impl.
				NewConfig(
					value.NewString(
						issue.Info,
						"foo",
					),
				).
				Pattern().
				Level(),
		)
	})

	t.Run("keeps zero level unspecified", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			issue.Type(0),
			impl.
				NewConfig(
					value.NewString(
						0,
						"foo",
					),
				).
				Pattern().
				Level(),
		)
	})
}
