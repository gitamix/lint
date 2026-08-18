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
