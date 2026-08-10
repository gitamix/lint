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
		assert.Equal(t, "foo", str.String())
	})

	t.Run("returns empty string when no value", func(t *testing.T) {
		t.Parallel()
		str := impl.NewString(issue.Critical, "")
		assert.Equal(t, "", str.String())
	})

	t.Run("preserves surrounding spaces", func(t *testing.T) {
		t.Parallel()
		str := impl.NewString(issue.Critical, " foo ")
		assert.Equal(t, " foo ", str.String())
	})
}

func TestString_Equal(t *testing.T) {
	t.Parallel()
	type args struct {
		other string
	}
	type want struct {
		ok bool
	}
	tests := []struct {
		name string
		v    impl.String
		args args
		want want
	}{
		{
			name: "same with exact",
			v:    impl.NewString(issue.Warning, "foo"),
			args: args{
				other: "foo",
			},
			want: want{
				ok: true,
			},
		},
		{
			name: "critical level & same with exact but with trim spaces",
			v:    impl.NewString(issue.Critical, " foo "),
			args: args{
				other: "foo",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "warning level & same with exact but with trim spaces",
			v:    impl.NewString(issue.Warning, " foo "),
			args: args{
				other: "foo",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "info level & same with exact but with trim spaces",
			v:    impl.NewString(issue.Info, " foo "),
			args: args{
				other: "foo",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "critical level & not same with exact",
			v:    impl.NewString(issue.Critical, "foo"),
			args: args{
				other: "bar",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "warning level & not same with exact",
			v:    impl.NewString(issue.Warning, "foo"),
			args: args{
				other: "bar",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "info level & not same with exact",
			v:    impl.NewString(issue.Info, "foo"),
			args: args{
				other: "bar",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "critical level & empty provided",
			v:    impl.NewString(issue.Critical, "foo"),
			args: args{
				other: "",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "warning level & empty provided",
			v:    impl.NewString(issue.Warning, "foo"),
			args: args{
				other: "",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "info level & empty provided",
			v:    impl.NewString(issue.Info, "foo"),
			args: args{
				other: "",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "critical level & empty exact",
			v:    impl.NewString(issue.Critical, ""),
			args: args{
				other: "foo",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "warning level & empty exact",
			v:    impl.NewString(issue.Warning, ""),
			args: args{
				other: "foo",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "info level & empty exact",
			v:    impl.NewString(issue.Info, ""),
			args: args{
				other: "foo",
			},
			want: want{
				ok: false,
			},
		},
		{
			name: "critical level & all empty",
			v:    impl.NewString(issue.Critical, ""),
			args: args{
				other: "",
			},
			want: want{
				ok: true,
			},
		},
		{
			name: "warning level & all empty",
			v:    impl.NewString(issue.Warning, ""),
			args: args{
				other: "",
			},
			want: want{
				ok: true,
			},
		},
		{
			name: "info level & all empty",
			v:    impl.NewString(issue.Info, ""),
			args: args{
				other: "",
			},
			want: want{
				ok: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.v.Equal(tt.args.other)
			if tt.want.ok {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
