package value_test

import (
	"testing"

	impl "github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStrings_Exact(t *testing.T) {
	t.Parallel()

	t.Run("returns the configured strings", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewStrings(issue.Critical, "foo", "bar").
			Exact()
		assert.Equal(t, []string{"foo", "bar"}, got)
	})

	t.Run("returns a copy of the underlying slice", func(t *testing.T) {
		t.Parallel()
		str := impl.NewStrings(issue.Critical, "foo", "bar")
		got := str.Exact()
		require.Equal(t, []string{"foo", "bar"}, got)
		got[0] = "bar"
		assert.Equal(t, "foo", str.Exact()[0])
	})

	t.Run("returns empty slice for no values", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewStrings(issue.Critical).
			Exact()
		assert.Equal(t, []string{}, got)
	})

	t.Run("single element", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewStrings(issue.Critical, "foo").
			Exact()
		assert.Equal(t, []string{"foo"}, got)
	})

	t.Run("empty string values are preserved", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewStrings(issue.Critical, "", "").
			Exact()
		assert.Equal(t, []string{"", ""}, got)
	})

	t.Run("order and duplicates are preserved", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewStrings(issue.Critical, "b", "a", "b").
			Exact()
		assert.Equal(t, []string{"b", "a", "b"}, got)
	})

	t.Run("exact matches equal", func(t *testing.T) {
		t.Parallel()
		str := impl.NewStrings(issue.Critical, "foo", "bar")
		assert.True(t, str.Equal(str.Exact()))
	})
}

func TestStrings_Equal(t *testing.T) {
	t.Parallel()

	t.Run("returns true for identical values", func(t *testing.T) {
		t.Parallel()
		str := impl.NewStrings(issue.Critical, "foo", "bar")
		assert.True(t, str.Equal([]string{"foo", "bar"}))
	})

	t.Run("returns false for differing values", func(t *testing.T) {
		t.Parallel()
		str := impl.
			NewStrings(
				issue.Critical,
				"foo",
				"bar",
			)
		assert.False(t, str.Equal([]string{"foo", "baz"}))
	})

	t.Run("returns false for different lengths", func(t *testing.T) {
		t.Parallel()
		str := impl.
			NewStrings(
				issue.Critical,
				"foo",
				"bar",
			)
		assert.False(t, str.Equal([]string{"foo"}))
	})

	t.Run("returns false for empty other when values present", func(t *testing.T) {
		t.Parallel()
		str := impl.
			NewStrings(
				issue.Critical,
				"foo",
			)
		assert.False(t, str.Equal([]string{}))
	})

	t.Run("returns false for empty config compared to empty slice", func(t *testing.T) {
		t.Parallel()
		str := impl.NewStrings(issue.Critical)
		assert.False(t, str.Equal([]string{}))
	})

	t.Run("returns false for different order", func(t *testing.T) {
		t.Parallel()
		str := impl.
			NewStrings(
				issue.Critical,
				"foo",
				"bar",
			)
		assert.False(t, str.Equal([]string{"bar", "foo"}))
	})
}

func TestStrings_Has(t *testing.T) {
	t.Parallel()

	t.Run("returns true when a single value is present", func(t *testing.T) {
		t.Parallel()
		str := impl.
			NewStrings(
				issue.Critical,
				"foo",
				"bar",
			)
		assert.True(t, str.Has("foo"))
	})

	t.Run("returns true when all values are present", func(t *testing.T) {
		t.Parallel()
		str := impl.
			NewStrings(
				issue.Critical,
				"foo",
				"bar",
				"baz",
			)
		assert.True(t, str.Has("foo", "bar"))
	})

	t.Run("returns false when a value is absent", func(t *testing.T) {
		t.Parallel()
		str := impl.
			NewStrings(
				issue.Critical,
				"foo",
				"bar",
			)
		assert.False(t, str.Has("foo", "baz"))
	})

	t.Run("returns false when single value is absent", func(t *testing.T) {
		t.Parallel()
		str := impl.
			NewStrings(
				issue.Critical,
				"foo",
				"bar",
			)
		assert.False(t, str.Has("baz"))
	})

	t.Run("returns false when one of many is absent", func(t *testing.T) {
		t.Parallel()
		str := impl.
			NewStrings(
				issue.Critical,
				"foo",
				"bar",
			)
		assert.False(t, str.Has("foo", "missing"))
	})

	t.Run("requires the first argument to be present", func(t *testing.T) {
		t.Parallel()
		str := impl.
			NewStrings(
				issue.Critical,
				"foo",
				"bar",
			)
		assert.False(t, str.Has("missing"))
	})
}

func TestStrings_Level(t *testing.T) {
	t.Parallel()

	t.Run("returns the configured level", func(t *testing.T) {
		t.Parallel()
		str := impl.NewStrings(issue.Critical)
		assert.Equal(t, issue.Critical, str.Level())
	})

	t.Run("returns warning when warning is configured", func(t *testing.T) {
		t.Parallel()
		str := impl.NewStrings(issue.Warning)
		assert.Equal(t, issue.Warning, str.Level())
	})

	t.Run("returns info level when info is configured", func(t *testing.T) {
		t.Parallel()
		str := impl.
			NewStrings(
				issue.Info,
				"foo",
				"bar",
			)
		assert.Equal(t, issue.Info, str.Level())
	})

	t.Run("returns warning despite configured values", func(t *testing.T) {
		t.Parallel()
		str := impl.
			NewStrings(
				issue.Warning,
				"foo",
				"bar",
				"baz",
			)
		assert.Equal(t, issue.Warning, str.Level())
	})
}
