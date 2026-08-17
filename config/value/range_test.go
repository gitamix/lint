package value_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestNewRange(t *testing.T) {
	t.Parallel()

	t.Run("creates a range with same min and max", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 5, 5).
			String()
		want := "5-5"
		assert.Equal(t, want, got)
	})

	t.Run("creates a range with negative bounds", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, -10, -1).
			String()
		want := "-10--1"
		assert.Equal(t, want, got)
	})

	t.Run("stores invalid bounds without panic in constructor", func(t *testing.T) {
		t.Parallel()
		assert.PanicsWithValue(
			t,
			"min value is greater than max: 10 > 5",
			func() {
				_ = impl.NewRange(issue.Critical, 10, 5).String()
			},
		)
	})
}

func TestRange_Level(t *testing.T) {
	t.Parallel()

	t.Run("returns critical level when critical is configured", func(t *testing.T) {
		t.Parallel()
		r := impl.NewRange(issue.Critical, 1, 100)
		assert.Equal(t, issue.Critical, r.Level())
	})

	t.Run("returns warning level when warning is configured", func(t *testing.T) {
		t.Parallel()
		r := impl.NewRange(issue.Warning, 1, 100)
		assert.Equal(t, issue.Warning, r.Level())
	})

	t.Run("returns info level when info is configured", func(t *testing.T) {
		t.Parallel()
		r := impl.NewRange(issue.Info, 1, 100)
		assert.Equal(t, issue.Info, r.Level())
	})

	t.Run("keeps unspecified level on unset", func(t *testing.T) {
		t.Parallel()
		var r impl.Range
		assert.True(t, r.Level().Unspecified())
	})
}

func TestRange_String(t *testing.T) {
	t.Parallel()

	t.Run("returns canonical form for positive bounds", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 1, 100).
			String()
		want := "1-100"
		assert.Equal(t, want, got)
	})

	t.Run("returns canonical form when bounds are equal", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 42, 42).
			String()
		want := "42-42"
		assert.Equal(t, want, got)
	})

	t.Run("returns canonical form for negative bounds", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, -100, -1).
			String()
		want := "-100--1"
		assert.Equal(t, want, got)
	})

	t.Run("panics when min is greater than max", func(t *testing.T) {
		t.Parallel()
		assert.PanicsWithValue(
			t,
			"min value is greater than max: 10 > 5",
			func() {
				_ = impl.NewRange(issue.Critical, 10, 5).String()
			},
		)
	})
}

func TestRange_Fit(t *testing.T) {
	t.Parallel()

	t.Run("returns true when n equals min", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 5, 10).
			Fit(5)
		assert.True(t, got)
	})

	t.Run("returns true when n equals max", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 5, 10).
			Fit(10)
		assert.True(t, got)
	})

	t.Run("returns true when n is strictly inside", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 5, 10).
			Fit(7)
		assert.True(t, got)
	})

	t.Run("returns false when n is strictly below min", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 5, 10).
			Fit(4)
		assert.False(t, got)
	})

	t.Run("returns false when n is strictly above max", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 5, 10).
			Fit(11)
		assert.False(t, got)
	})

	t.Run("returns true when range is a single point and n matches", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 7, 7).
			Fit(7)
		assert.True(t, got)
	})

	t.Run("returns false when range is a single point and n differs", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 7, 7).
			Fit(8)
		assert.False(t, got)
	})

	t.Run("returns true for negative n inside negative range", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, -10, -1).
			Fit(-5)
		assert.True(t, got)
	})

	t.Run("panics when min is greater than max", func(t *testing.T) {
		t.Parallel()
		assert.PanicsWithValue(
			t,
			"min value is greater than max: 10 > 5",
			func() {
				_ = impl.NewRange(issue.Critical, 10, 5).Fit(7)
			},
		)
	})
}

func TestRange_Below(t *testing.T) {
	t.Parallel()

	t.Run("returns true when n is strictly less than min", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 5, 10).
			Below(4)
		assert.True(t, got)
	})

	t.Run("returns false when n equals min", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 5, 10).
			Below(5)
		assert.False(t, got)
	})

	t.Run("returns false when n is strictly greater than min", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 5, 10).
			Below(6)
		assert.False(t, got)
	})

	t.Run("returns true for negative n below negative min", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, -10, -1).
			Below(-11)
		assert.True(t, got)
	})

	t.Run("panics when min is greater than max", func(t *testing.T) {
		t.Parallel()
		assert.PanicsWithValue(
			t,
			"min value is greater than max: 10 > 5",
			func() {
				_ = impl.NewRange(issue.Critical, 10, 5).Below(1)
			},
		)
	})
}

func TestRange_Above(t *testing.T) {
	t.Parallel()

	t.Run("returns true when n is strictly greater than max", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 5, 10).
			Above(11)
		assert.True(t, got)
	})

	t.Run("returns false when n equals max", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 5, 10).
			Above(10)
		assert.False(t, got)
	})

	t.Run("returns false when n is strictly less than max", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, 5, 10).
			Above(9)
		assert.False(t, got)
	})

	t.Run("returns true for positive n above negative max", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewRange(issue.Critical, -10, -1).
			Above(0)
		assert.True(t, got)
	})

	t.Run("panics when min is greater than max", func(t *testing.T) {
		t.Parallel()
		assert.PanicsWithValue(
			t,
			"min value is greater than max: 10 > 5",
			func() {
				_ = impl.NewRange(issue.Critical, 10, 5).Above(1)
			},
		)
	})
}
