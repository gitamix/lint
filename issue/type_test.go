package issue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/issue"
)

func TestType_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		t    impl.Type
		want string
	}{
		{
			name: "Critical",
			t:    impl.Critical,
			want: "critical",
		},
		{
			name: "Warning",
			t:    impl.Warning,
			want: "warning",
		},
		{
			name: "Info",
			t:    impl.Info,
			want: "info",
		},
		{
			name: "custom 255",
			t:    impl.Type(255),
			want: "",
		},
		{
			name: "custom zero",
			t:    impl.Type(0),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				tt.t.String(),
			)
		})
	}
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

	t.Run("custom 255 type is specified", func(t *testing.T) {
		t.Parallel()
		tt := impl.Type(255)
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
		tt := impl.Type(255)
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

	t.Run("custom 255 type is unknown", func(t *testing.T) {
		t.Parallel()
		tt := impl.Type(255)
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
			impl.Type(0),
			impl.Parse(" critical "),
		)
	})

	t.Run("Critical not parsed", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			impl.Type(0),
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
			impl.Type(0),
			impl.Parse("foo"),
		)
	})
}

func TestParseOr(t *testing.T) {
	t.Parallel()
	type args struct {
		s   string
		def impl.Type
	}
	type want struct {
		typ impl.Type
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "critical parsed",
			args: args{
				s:   "critical",
				def: impl.Info,
			},
			want: want{
				typ: impl.Critical,
			},
		},
		{
			name: "critical not trimmed not parsed",
			args: args{
				s:   " critical ",
				def: impl.Info,
			},
			want: want{
				typ: impl.Info,
			},
		},
		{
			name: "Critical not parsed",
			args: args{
				s:   "Critical",
				def: impl.Info,
			},
			want: want{
				typ: impl.Info,
			},
		},
		{
			name: "CRITICAL not parsed",
			args: args{
				s:   "CRITICAL",
				def: impl.Info,
			},
			want: want{
				typ: impl.Info,
			},
		},
		{
			name: "warning parsed",
			args: args{
				s:   "warning",
				def: impl.Info,
			},
			want: want{
				typ: impl.Warning,
			},
		},
		{
			name: "warning not trimmed not parsed",
			args: args{
				s:   " warning ",
				def: impl.Info,
			},
			want: want{
				typ: impl.Info,
			},
		},
		{
			name: "Warning not parsed",
			args: args{
				s:   "Warning",
				def: impl.Info,
			},
			want: want{
				typ: impl.Info,
			},
		},
		{
			name: "WARNING not parsed",
			args: args{
				s:   "WARNING",
				def: impl.Info,
			},
			want: want{
				typ: impl.Info,
			},
		},
		{
			name: "info parsed",
			args: args{
				s:   "info",
				def: impl.Warning,
			},
			want: want{
				typ: impl.Info,
			},
		},
		{
			name: "info not trimmed not parsed",
			args: args{
				s:   " info ",
				def: impl.Warning,
			},
			want: want{
				typ: impl.Warning,
			},
		},
		{
			name: "Info not parsed",
			args: args{
				s:   "Info",
				def: impl.Warning,
			},
			want: want{
				typ: impl.Warning,
			},
		},
		{
			name: "INFO not parsed",
			args: args{
				s:   "INFO",
				def: impl.Warning,
			},
			want: want{
				typ: impl.Warning,
			},
		},
		{
			name: "empty string",
			args: args{
				s:   "",
				def: impl.Info,
			},
			want: want{
				typ: impl.Info,
			},
		},
		{
			name: "unknown string",
			args: args{
				s:   "foo",
				def: impl.Critical,
			},
			want: want{
				typ: impl.Critical,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.typ,
				impl.ParseOr(
					tt.args.s,
					tt.args.def,
				),
			)
		})
	}
}
