package value_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestInteger_String(t *testing.T) {
	t.Parallel()

	t.Run("returns the exact value", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Critical, 42).
			String()
		want := "42"
		assert.Equal(t, want, got)
	})

	t.Run("returns zero string when value is zero", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Critical, 0).
			String()
		want := "0"
		assert.Equal(t, want, got)
	})

	t.Run("returns negative number with minus sign", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Critical, -7).
			String()
		want := "-7"
		assert.Equal(t, want, got)
	})

	t.Run("returns large number without formatting", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Critical, 1000000).
			String()
		want := "1000000"
		assert.Equal(t, want, got)
	})
}

func TestInteger_Equal(t *testing.T) {
	t.Parallel()

	t.Run("returns true when equal", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Warning, 42).
			Equal(42)
		assert.True(t, got)
	})

	t.Run("returns false when not equal", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Warning, 42).
			Equal(7)
		assert.False(t, got)
	})

	t.Run("critical level & returns true when equal", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Critical, 42).
			Equal(42)
		assert.True(t, got)
	})

	t.Run("critical level & returns false when not equal", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Critical, 42).
			Equal(7)
		assert.False(t, got)
	})

	t.Run("warning level & returns true when equal", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Warning, 42).
			Equal(42)
		assert.True(t, got)
	})

	t.Run("warning level & returns false when not equal", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Warning, 42).
			Equal(7)
		assert.False(t, got)
	})

	t.Run("info level & returns true when equal", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Info, 42).
			Equal(42)
		assert.True(t, got)
	})

	t.Run("info level & returns false when not equal", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Info, 42).
			Equal(7)
		assert.False(t, got)
	})

	t.Run("critical level & returns false when zero provided and exact nonzero", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Critical, 42).
			Equal(0)
		assert.False(t, got)
	})

	t.Run("warning level & returns false when zero provided and exact nonzero", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Warning, 42).
			Equal(0)
		assert.False(t, got)
	})

	t.Run("info level & returns false when zero provided and exact nonzero", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Info, 42).
			Equal(0)
		assert.False(t, got)
	})

	t.Run("critical level & returns false when exact zero and other nonzero", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Critical, 0).
			Equal(42)
		assert.False(t, got)
	})

	t.Run("warning level & returns false when exact zero and other nonzero", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Warning, 0).
			Equal(42)
		assert.False(t, got)
	})

	t.Run("info level & returns false when exact zero and other nonzero", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Info, 0).
			Equal(42)
		assert.False(t, got)
	})

	t.Run("critical level & returns true when both zero", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Critical, 0).
			Equal(0)
		assert.True(t, got)
	})

	t.Run("warning level & returns true when both zero", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Warning, 0).
			Equal(0)
		assert.True(t, got)
	})

	t.Run("info level & returns true when both zero", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Info, 0).
			Equal(0)
		assert.True(t, got)
	})

	t.Run("critical level & returns true when both negative and equal", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Critical, -7).
			Equal(-7)
		assert.True(t, got)
	})

	t.Run("warning level & returns false when signs differ", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Warning, -7).
			Equal(7)
		assert.False(t, got)
	})
}

func TestInteger_WithLevel(t *testing.T) {
	t.Parallel()

	t.Run("changes critical to info level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewInteger(issue.Critical, 1).
			WithLevel(issue.Info)
		assert.Equal(t, issue.Info, got.Level())
	})

	t.Run("changes critical to zero level", func(t *testing.T) {
		t.Parallel()
		assert.True(
			t,
			impl.
				NewInteger(issue.Critical, 1).
				WithLevel(issue.Type(0)).
				Level().
				Unspecified(),
		)
	})

	t.Run("does not change level for the previous one", func(t *testing.T) {
		t.Parallel()
		prev := impl.NewInteger(issue.Critical, 1)
		_ = prev.WithLevel(issue.Info)
		assert.Equal(t, issue.Critical, prev.Level())
	})
}
