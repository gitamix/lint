package value_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestPattern_String(t *testing.T) {
	t.Parallel()

	t.Run("returns the regexp expression", func(t *testing.T) {
		t.Parallel()
		pattern := impl.NewPattern(issue.Critical, regexp.MustCompile(`^foo$`))
		assert.Equal(t, "^foo$", pattern.String())
	})

	t.Run("panics when expression not set", func(t *testing.T) {
		t.Parallel()
		assert.Panics(t, func() {
			_ = impl.
				NewPattern(issue.Critical, nil).
				String()
		})
	})
}

func TestPattern_Level(t *testing.T) {
	t.Parallel()

	t.Run("returns the provided level", func(t *testing.T) {
		t.Parallel()
		pattern := impl.NewPattern(issue.Info, regexp.MustCompile(`foo`))
		assert.Equal(t, issue.Info, pattern.Level())
	})

	t.Run("keeps unknown level when it is incorrect", func(t *testing.T) {
		t.Parallel()
		var pattern impl.Pattern
		assert.True(t, pattern.Level().Unknown())
	})
}

func TestPattern_Match(t *testing.T) {
	t.Parallel()

	t.Run("matches the expression", func(t *testing.T) {
		t.Parallel()
		pattern := impl.NewPattern(issue.Warning, regexp.MustCompile(`^foo$`))
		assert.True(t, pattern.Match("foo"))
	})

	t.Run("matches a substring expression", func(t *testing.T) {
		t.Parallel()
		pattern := impl.NewPattern(issue.Warning, regexp.MustCompile(`foo`))
		assert.True(t, pattern.Match("barfoobaz"))
	})

	t.Run("does not match the expression", func(t *testing.T) {
		t.Parallel()
		pattern := impl.NewPattern(issue.Warning, regexp.MustCompile(`^foo$`))
		assert.False(t, pattern.Match("bar"))
	})

	t.Run("matches an expanded expression", func(t *testing.T) {
		t.Parallel()
		pattern := impl.NewPattern(issue.Info, regexp.MustCompile(`^[a-z]+\d+$`))
		assert.True(t, pattern.Match("abc123"))
	})

	t.Run("does not match an expanded expression", func(t *testing.T) {
		t.Parallel()
		assert.False(
			t,
			impl.
				NewPattern(
					issue.Info,
					regexp.MustCompile(`^[a-z]+\d+$`),
				).
				Match("123abc"),
		)
	})

	t.Run("does not match when expression not set", func(t *testing.T) {
		t.Parallel()
		assert.Panics(t, func() {
			_ = impl.
				NewPattern(issue.Warning, nil).
				Match("foo")
		})
	})
}

func TestPattern_WithLevel(t *testing.T) {
	t.Parallel()

	t.Run("changes critical to info level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewPattern(
				issue.Critical,
				regexp.MustCompile(`foo`),
			).
			WithLevel(issue.Info)
		want := impl.NewPattern(
			issue.Info,
			regexp.MustCompile(`foo`),
		)
		assert.Equal(t, want, got)
	})

	t.Run("changes critical to zero level", func(t *testing.T) {
		t.Parallel()
		assert.True(
			t,
			impl.
				NewPattern(
					issue.Critical,
					regexp.MustCompile(`foo`),
				).
				WithLevel(issue.Type(0)).
				Level().
				Unspecified(),
		)
	})

	t.Run("does not change level for the previous one", func(t *testing.T) {
		t.Parallel()
		prev := impl.NewPattern(
			issue.Critical,
			regexp.MustCompile(`foo`),
		)
		_ = prev.WithLevel(issue.Info)
		assert.Equal(t, issue.Critical, prev.Level())
	})
}
