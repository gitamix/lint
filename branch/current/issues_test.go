package current_test

import (
	"context"
	"testing"

	"github.com/gitamix/types/branch"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/branch/current"
	config "github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/ticket"
	"github.com/gitamix/lint/config/ticket/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/errs"
	"github.com/gitamix/lint/internal/test/fake/repo/git"
	"github.com/gitamix/lint/issue"
)

func TestBranch_Issues(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
	}
	type want struct {
		issues []issue.Issue
		err    error
		panic  bool
	}
	tests := []struct {
		name string
		b    *impl.Branch
		args args
		want want
	}{
		{
			name: "ok",
			b: impl.NewBranch(
				git.NewRepository(
					git.WithCurrentBranch(
						func(_ context.Context) (branch.Branch, error) {
							return branch.NewBranch(
								branch.NewName("feature/TASK-123"),
							), nil
						},
					),
				),
				config.NewConfig(
					config.WithName(
						name.NewConfig(
							value.NewString(
								issue.Critical,
								`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
							),
						),
					),
					config.WithTicket(
						ticket.NewConfig(
							ticket.WithID(
								id.NewConfig(
									value.NewString(
										issue.Warning,
										`(TASK|PROJ|BUG)-[0-9]+`,
									),
								),
							),
						),
					),
				),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				issues: []issue.Issue{},
				err:    nil,
			},
		},
		{
			name: "failed to get current branch from git",
			b: impl.NewBranch(
				git.NewRepository(
					git.WithCurrentBranch(
						func(_ context.Context) (branch.Branch, error) {
							return branch.Branch{}, assert.AnError
						},
					),
				),
				config.NewConfig(
					config.WithName(
						name.NewConfig(
							value.NewString(
								issue.Critical,
								`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
							),
						),
					),
				),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				issues: nil,
				err:    errs.ErrGitFailed,
			},
		},
		{
			name: "critical on name not matched pattern",
			b: impl.NewBranch(
				git.NewRepository(
					git.WithCurrentBranch(
						func(_ context.Context) (branch.Branch, error) {
							return branch.NewBranch(
								branch.NewName("release/TASK-123"),
							), nil
						},
					),
				),
				config.NewConfig(
					config.WithName(
						name.NewConfig(
							value.NewString(
								issue.Critical,
								`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
							),
						),
					),
				),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				issues: []issue.Issue{
					issue.NewCritical(`branch name doesn't match the required pattern '^(feature|bugfix|hotfix)/[A-Z]+-\d+'`),
				},
				err: nil,
			},
		},
		{
			name: "criticals on name & ticket not match pattern",
			b: impl.NewBranch(
				git.NewRepository(
					git.WithCurrentBranch(
						func(_ context.Context) (branch.Branch, error) {
							return branch.NewBranch(
								branch.NewName("my-favorite-feature"),
							), nil
						},
					),
				),
				config.NewConfig(
					config.WithName(
						name.NewConfig(
							value.NewString(
								issue.Critical,
								`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
							),
						),
					),
					config.WithTicket(
						ticket.NewConfig(
							ticket.WithID(
								id.NewConfig(
									value.NewString(
										issue.Critical,
										`(TASK|PROJ|BUG)-[0-9]+`,
									),
								),
							),
						),
					),
				),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				issues: []issue.Issue{
					issue.NewCritical(`branch name doesn't match the required pattern '^(feature|bugfix|hotfix)/[A-Z]+-\d+'`),
					issue.NewCritical(`ticket doesn't match the required pattern '(TASK|PROJ|BUG)-[0-9]+'`),
				},
				err: nil,
			},
		},
		{
			name: "warning on empty branch name",
			b: impl.NewBranch(
				git.NewRepository(
					git.WithCurrentBranch(
						func(_ context.Context) (branch.Branch, error) {
							return branch.NewBranch(
								branch.NewName(""),
							), nil
						},
					),
				),
				config.NewConfig(
					config.WithName(
						name.NewConfig(
							value.NewString(
								issue.Warning,
								`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
							),
						),
					),
				),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				issues: []issue.Issue{
					issue.NewWarning(`branch name doesn't match the required pattern '^(feature|bugfix|hotfix)/[A-Z]+-\d+'`),
				},
				err: nil,
			},
		},
		{
			name: "empty pattern",
			b: impl.NewBranch(
				git.NewRepository(
					git.WithCurrentBranch(
						func(_ context.Context) (branch.Branch, error) {
							return branch.NewBranch(
								branch.NewName("feature/TASK-123"),
							), nil
						},
					),
				),
				config.NewConfig(
					config.WithName(
						name.NewConfig(
							value.NewString(
								issue.Warning,
								``,
							),
						),
					),
				),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				issues: []issue.Issue{},
				err:    nil,
			},
		},
		{
			name: "empty ticket id pattern & name matched pattern",
			b: impl.NewBranch(
				git.NewRepository(
					git.WithCurrentBranch(
						func(_ context.Context) (branch.Branch, error) {
							return branch.NewBranch(
								branch.NewName("feature/TASK-123"),
							), nil
						},
					),
				),
				config.NewConfig(
					config.WithName(
						name.NewConfig(
							value.NewString(
								issue.Warning,
								`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
							),
						),
					),
					config.WithTicket(
						ticket.NewConfig(
							ticket.WithID(
								id.NewConfig(
									value.NewString(
										issue.Critical,
										``,
									),
								),
							),
						),
					),
				),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				issues: []issue.Issue{},
				err:    nil,
			},
		},
		{
			name: "empty current branch name and pattern",
			b: impl.NewBranch(
				git.NewRepository(
					git.WithCurrentBranch(
						func(_ context.Context) (branch.Branch, error) {
							return branch.NewBranch(
								branch.NewName(""),
							), nil
						},
					),
				),
				config.NewConfig(
					config.WithName(
						name.NewConfig(
							value.NewString(
								issue.Warning,
								``,
							),
						),
					),
				),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				issues: []issue.Issue{},
				err:    nil,
			},
		},
		{
			name: "empty config",
			b: impl.NewBranch(
				git.NewRepository(
					git.WithCurrentBranch(
						func(_ context.Context) (branch.Branch, error) {
							return branch.NewBranch(
								branch.NewName("feature/TASK-123"),
							), nil
						},
					),
				),
				config.NewConfig(),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				issues: []issue.Issue{},
				err:    nil,
			},
		},
		{
			name: "empty value",
			b:    &impl.Branch{},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				panic: true,
			},
		},
		{
			name: "nil value",
			b:    nil,
			args: args{
				ctx: context.Background(),
			},
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
					_, _ = tt.b.Issues(tt.args.ctx)
				})
				return
			}
			got, gotErr := tt.b.Issues(tt.args.ctx)
			assert.Equal(t, tt.want.issues, got)
			assert.ErrorIs(t, gotErr, tt.want.err)
		})
	}
}
