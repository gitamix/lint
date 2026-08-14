package task_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_ID(t *testing.T) {
	t.Parallel()
	type want struct {
		id id.Config
	}
	tests := []struct {
		name string
		c    impl.Config
		want want
	}{
		{
			name: "with id config & correct pattern",
			c: impl.NewConfig(
				impl.WithID(
					id.NewConfig(
						value.NewString(
							issue.Warning,
							`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
						),
					),
				),
			),
			want: want{
				id: id.NewConfig(
					value.NewString(
						issue.Warning,
						`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
					),
				),
			},
		},
		{
			name: "with id config & just a word as pattern",
			c: impl.NewConfig(
				impl.WithID(
					id.NewConfig(
						value.NewString(
							issue.Critical,
							"foo",
						),
					),
				),
			),
			want: want{
				id: id.NewConfig(
					value.NewString(
						issue.Critical,
						"foo",
					),
				),
			},
		},
		{
			name: "with id config & empty pattern",
			c: impl.NewConfig(
				impl.WithID(
					id.NewConfig(
						value.NewString(
							issue.Info,
							"",
						),
					),
				),
			),
			want: want{
				id: id.NewConfig(
					value.NewString(
						issue.Info,
						"",
					),
				),
			},
		},
		{
			name: "without any option",
			c:    impl.NewConfig(),
			want: want{
				id: id.Config{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.id,
				tt.c.ID(),
			)
		})
	}
}
