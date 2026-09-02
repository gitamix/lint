package value_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestString_String(t *testing.T) {
	t.Parallel()

	t.Run("returns the exact value", func(t *testing.T) {
		t.Parallel()
		str := impl.NewString(issue.Critical, "foo")
		assert.Equal(
			t,
			"foo",
			str.String(),
		)
	})

	t.Run("returns empty string when no value", func(t *testing.T) {
		t.Parallel()
		str := impl.NewString(issue.Critical, "")
		assert.Equal(
			t,
			"",
			str.String(),
		)
	})

	t.Run("preserves surrounding spaces", func(t *testing.T) {
		t.Parallel()
		str := impl.NewString(issue.Critical, " foo ")
		assert.Equal(
			t,
			" foo ",
			str.String(),
		)
	})
}

func TestString_Equal(t *testing.T) {
	t.Parallel()

	t.Run("same with exact", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Warning, "foo")
		assert.True(t, v.Equal("foo"))
	})

	t.Run("critical level & same with exact but with trim spaces", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Critical, " foo ")
		assert.False(t, v.Equal("foo"))
	})

	t.Run("warning level & same with exact but with trim spaces", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Warning, " foo ")
		assert.False(t, v.Equal("foo"))
	})

	t.Run("info level & same with exact but with trim spaces", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Info, " foo ")
		assert.False(t, v.Equal("foo"))
	})

	t.Run("critical level & not same with exact", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Critical, "foo")
		assert.False(t, v.Equal("bar"))
	})

	t.Run("warning level & not same with exact", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Warning, "foo")
		assert.False(t, v.Equal("bar"))
	})

	t.Run("info level & not same with exact", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Info, "foo")
		assert.False(t, v.Equal("bar"))
	})

	t.Run("critical level & empty provided", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Critical, "foo")
		assert.False(t, v.Equal(""))
	})

	t.Run("warning level & empty provided", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Warning, "foo")
		assert.False(t, v.Equal(""))
	})

	t.Run("info level & empty provided", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Info, "foo")
		assert.False(t, v.Equal(""))
	})

	t.Run("critical level & empty exact", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Critical, "")
		assert.False(t, v.Equal("foo"))
	})

	t.Run("warning level & empty exact", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Warning, "")
		assert.False(t, v.Equal("foo"))
	})

	t.Run("info level & empty exact", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Info, "")
		assert.False(t, v.Equal("foo"))
	})

	t.Run("critical level & all empty", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Critical, "")
		assert.True(t, v.Equal(""))
	})

	t.Run("warning level & all empty", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Warning, "")
		assert.True(t, v.Equal(""))
	})

	t.Run("info level & all empty", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Info, "")
		assert.True(t, v.Equal(""))
	})
}

func TestString_Empty(t *testing.T) {
	t.Parallel()

	t.Run("critical level & empty value", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Critical, "")
		assert.True(t, v.Empty())
	})

	t.Run("warning level & empty value", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Warning, "")
		assert.True(t, v.Empty())
	})

	t.Run("info level & empty value", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Info, "")
		assert.True(t, v.Empty())
	})

	t.Run("critical level & non-empty value", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Critical, "foo")
		assert.False(t, v.Empty())
	})

	t.Run("warning level & non-empty value", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Warning, "foo")
		assert.False(t, v.Empty())
	})

	t.Run("info level & non-empty value", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Info, "foo")
		assert.False(t, v.Empty())
	})

	t.Run("critical level & value with surrounding spaces", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Critical, " foo ")
		assert.False(t, v.Empty())
	})

	t.Run("warning level & value with surrounding spaces", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Warning, " foo ")
		assert.False(t, v.Empty())
	})

	t.Run("info level & value with surrounding spaces", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Info, " foo ")
		assert.False(t, v.Empty())
	})

	t.Run("zero level & empty value", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Type(0), "")
		assert.True(t, v.Empty())
	})

	t.Run("zero level & non-empty value", func(t *testing.T) {
		t.Parallel()
		v := impl.NewString(issue.Type(0), "foo")
		assert.False(t, v.Empty())
	})
}

func TestString_WithLevel(t *testing.T) {
	t.Parallel()

	t.Run("changes critical to info level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewString(issue.Critical, "foo").
			WithLevel(issue.Info)
		assert.Equal(
			t,
			issue.Info,
			got.Level(),
		)
	})

	t.Run("changes critical to zero level", func(t *testing.T) {
		t.Parallel()
		assert.True(
			t,
			impl.
				NewString(issue.Critical, "foo").
				WithLevel(issue.Type(0)).
				Level().
				Unspecified(),
		)
	})

	t.Run("does not change level for the previous one", func(t *testing.T) {
		t.Parallel()
		prev := impl.NewString(issue.Critical, "foo")
		_ = prev.WithLevel(issue.Info)
		assert.Equal(
			t,
			issue.Critical,
			prev.Level(),
		)
	})
}
