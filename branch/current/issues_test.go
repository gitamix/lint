package current_test

import (
	"context"
	"testing"

	"github.com/gitamix/types/branch"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/branch/current"
	config "github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/errs"
	"github.com/gitamix/lint/internal/test/fake/repo/git"
	"github.com/gitamix/lint/issue"
)

func TestLinter_Issues(t *testing.T) {
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
		l    *impl.Linter
		args args
		want want
	}{
		{
			name: "ok",
			l: impl.NewLinter(
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
			l: impl.NewLinter(
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
			l: impl.NewLinter(
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
			name: "warning on empty branch name",
			l: impl.NewLinter(
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
			name: "warning on empty pattern",
			l: impl.NewLinter(
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
			l: impl.NewLinter(
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
			l: impl.NewLinter(
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
			name: "empty linter",
			l:    &impl.Linter{},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				panic: true,
			},
		},
		{
			name: "nil linter",
			l:    nil,
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
					_, _ = tt.l.Issues(tt.args.ctx)
				})
				return
			}
			got, gotErr := tt.l.Issues(tt.args.ctx)
			assert.Equal(t, tt.want.issues, got)
			assert.ErrorIs(t, gotErr, tt.want.err)
		})
	}
}
