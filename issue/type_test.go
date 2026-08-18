package issue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/issue"
)

func TestType_String(t *testing.T) {
	t.Parallel()

	t.Run("Critical", func(t *testing.T) {
		t.Parallel()
		tt := impl.Critical
		assert.Equal(
			t,
			"critical",
			tt.String(),
		)
	})

	t.Run("Warning", func(t *testing.T) {
		t.Parallel()
		tt := impl.Warning
		assert.Equal(
			t,
			"warning",
			tt.String(),
		)
	})

	t.Run("Info", func(t *testing.T) {
		t.Parallel()
		tt := impl.Info
		assert.Equal(
			t,
			"info",
			tt.String(),
		)
	})

	t.Run("custom 127 returns empty string", func(t *testing.T) {
		t.Parallel()
		tt := impl.Type(127)
		assert.Equal(
			t,
			"",
			tt.String(),
		)
	})

	t.Run("custom zero returns empty string", func(t *testing.T) {
		t.Parallel()
		tt := impl.Type(0)
		assert.Equal(
			t,
			"",
			tt.String(),
		)
	})
}

func TestType_Unspecified(t *testing.T) {
	t.Parallel()

	t.Run("zero type is unspecified", func(t *testing.T) {
		t.Parallel()
		tt := impl.Type(0)
		assert.True(t, tt.Unspecified())
	})

	t.Run("critical type is specified", func(t *testing.T) {
		t.Parallel()
		tt := impl.Critical
		assert.False(t, tt.Unspecified())
	})

	t.Run("warning type is specified", func(t *testing.T) {
		t.Parallel()
		tt := impl.Warning
		assert.False(t, tt.Unspecified())
	})

	t.Run("info type is specified", func(t *testing.T) {
		t.Parallel()
		tt := impl.Info
		assert.False(t, tt.Unspecified())
	})

	t.Run("custom 127 type is specified", func(t *testing.T) {
		t.Parallel()
		tt := impl.Type(127)
		assert.False(t, tt.Unspecified())
	})
}

func TestType_In(t *testing.T) {
	t.Parallel()

	t.Run("type exists in list", func(t *testing.T) {
		t.Parallel()
		tt := impl.Warning
		assert.True(
			t,
			tt.In(
				impl.Critical,
				impl.Warning,
				impl.Info,
			),
		)
	})

	t.Run("type does not exist in list", func(t *testing.T) {
		t.Parallel()
		tt := impl.Type(127)
		assert.False(
			t,
			tt.In(
				impl.Critical,
				impl.Warning,
				impl.Info,
			),
		)
	})

	t.Run("empty list does not contain type", func(t *testing.T) {
		t.Parallel()
		tt := impl.Critical
		assert.False(t, tt.In())
	})
}

func TestType_Unknown(t *testing.T) {
	t.Parallel()

	t.Run("zero type is unknown", func(t *testing.T) {
		t.Parallel()
		tt := impl.Type(0)
		assert.True(t, tt.Unknown())
	})

	t.Run("critical type is known", func(t *testing.T) {
		t.Parallel()
		tt := impl.Critical
		assert.False(t, tt.Unknown())
	})

	t.Run("warning type is known", func(t *testing.T) {
		t.Parallel()
		tt := impl.Warning
		assert.False(t, tt.Unknown())
	})

	t.Run("info type is known", func(t *testing.T) {
		t.Parallel()
		tt := impl.Info
		assert.False(t, tt.Unknown())
	})

	t.Run("custom 127 type is unknown", func(t *testing.T) {
		t.Parallel()
		tt := impl.Type(127)
		assert.True(t, tt.Unknown())
	})
}

func TestParse(t *testing.T) {
	t.Parallel()

	t.Run("critical parsed", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			impl.Critical,
			impl.Parse("critical"),
		)
	})

	t.Run("warning parsed", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			impl.Warning,
			impl.Parse("warning"),
		)
	})

	t.Run("info parsed", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			impl.Info,
			impl.Parse("info"),
		)
	})

	t.Run("not trimmed critical not parsed", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			impl.Type(-1),
			impl.Parse(" critical "),
		)
	})

	t.Run("Critical not parsed", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			impl.Type(-1),
			impl.Parse("Critical"),
		)
	})

	t.Run("empty string not parsed", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			impl.Type(0),
			impl.Parse(""),
		)
	})

	t.Run("unknown string not parsed", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			impl.Type(-1),
			impl.Parse("foo"),
		)
	})
}
