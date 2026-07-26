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
	type want struct {
		value value.String
	}
	tests := []struct {
		name string
		c    impl.Config
		want want
	}{
		{
			name: "regexp",
			c: impl.NewConfig(
				value.NewString(
					issue.Warning,
					`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
				),
			),
			want: want{
				value: value.NewString(
					issue.Warning,
					`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
				),
			},
		},
		{
			name: "critical level & just a word",
			c: impl.NewConfig(
				value.NewString(
					issue.Critical,
					"foo",
				),
			),
			want: want{
				value: value.NewString(
					issue.Critical,
					"foo",
				),
			},
		},
		{
			name: "warning level & just a word",
			c: impl.NewConfig(
				value.NewString(
					issue.Warning,
					"foo",
				),
			),
			want: want{
				value: value.NewString(
					issue.Warning,
					"foo",
				),
			},
		},
		{
			name: "info level & just a word",
			c: impl.NewConfig(
				value.NewString(
					issue.Info,
					"foo",
				),
			),
			want: want{
				value: value.NewString(
					issue.Info,
					"foo",
				),
			},
		},
		{
			name: "critical level & empty string",
			c: impl.NewConfig(
				value.NewString(
					issue.Critical,
					"",
				),
			),
			want: want{
				value: value.NewString(
					issue.Critical,
					"",
				),
			},
		},
		{
			name: "warning level & empty string",
			c: impl.NewConfig(
				value.NewString(
					issue.Warning,
					"",
				),
			),
			want: want{
				value: value.NewString(
					issue.Warning,
					"",
				),
			},
		},
		{
			name: "info level & empty string",
			c: impl.NewConfig(
				value.NewString(
					issue.Info,
					"",
				),
			),
			want: want{
				value: value.NewString(
					issue.Info,
					"",
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.c.Pattern(),
			)
		})
	}
}

func TestConfig_Level(t *testing.T) {
	t.Parallel()
	type want struct {
		level issue.Type
	}
	tests := []struct {
		name string
		c    impl.Config
		want want
	}{
		{
			name: "critical level",
			c: impl.NewConfig(
				value.NewString(
					issue.Critical,
					"foo",
				),
			),
			want: want{
				level: issue.Critical,
			},
		},
		{
			name: "warning level",
			c: impl.NewConfig(
				value.NewString(
					issue.Warning,
					"foo",
				),
			),
			want: want{
				level: issue.Warning,
			},
		},
		{
			name: "info level",
			c: impl.NewConfig(
				value.NewString(
					issue.Info,
					"foo",
				),
			),
			want: want{
				level: issue.Info,
			},
		},
		{
			name: "zero level defaults to warning",
			c: impl.NewConfig(
				value.NewString(
					0,
					"foo",
				),
			),
			want: want{
				level: issue.Warning,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.level,
				tt.c.Pattern().Level(),
			)
		})
	}
}
