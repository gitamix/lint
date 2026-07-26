package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config"
	"github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/ticket"
	"github.com/gitamix/lint/config/ticket/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_Branch(t *testing.T) {
	t.Parallel()
	type want struct {
		branch branch.Config
		panic  bool
	}
	tests := []struct {
		name string
		c    *impl.Config
		want want
	}{
		{
			name: "with branch name & ticket configs",
			c: impl.NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithName(
							name.NewConfig(
								value.NewString(
									issue.Warning,
									`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
								),
							),
						),
						branch.WithTicket(
							ticket.NewConfig(
								ticket.WithID(
									id.NewConfig(
										value.NewString(
											issue.Critical,
											`[A-Z]+-\d+`,
										),
									),
								),
							),
						),
					),
				),
			),
			want: want{
				branch: branch.NewConfig(
					branch.WithName(
						name.NewConfig(
							value.NewString(
								issue.Warning,
								`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
							),
						),
					),
					branch.WithTicket(
						ticket.NewConfig(
							ticket.WithID(
								id.NewConfig(
									value.NewString(
										issue.Critical,
										`[A-Z]+-\d+`,
									),
								),
							),
						),
					),
				),
			},
		},
		{
			name: "with name config only",
			c: impl.NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithName(
							name.NewConfig(
								value.NewString(
									issue.Critical,
									"foo",
								),
							),
						),
					),
				),
			),
			want: want{
				branch: branch.NewConfig(
					branch.WithName(
						name.NewConfig(
							value.NewString(
								issue.Critical,
								"foo",
							),
						),
					),
				),
			},
		},
		{
			name: "with ticket config only",
			c: impl.NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithTicket(
							ticket.NewConfig(
								ticket.WithID(
									id.NewConfig(
										value.NewString(
											issue.Info,
											"",
										),
									),
								),
							),
						),
					),
				),
			),
			want: want{
				branch: branch.NewConfig(
					branch.WithTicket(
						ticket.NewConfig(
							ticket.WithID(
								id.NewConfig(
									value.NewString(
										issue.Info,
										"",
									),
								),
							),
						),
					),
				),
			},
		},
		{
			name: "without any options",
			c:    impl.NewConfig(),
			want: want{
				branch: branch.NewConfig(),
			},
		},
		{
			name: "default value",
			c:    nil,
			want: want{
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.c.Branch()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.branch,
				tt.c.Branch(),
			)
		})
	}
}
