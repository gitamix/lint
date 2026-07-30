package name_test

import (
	"testing"

	"github.com/gitamix/types/branch"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/branch/name"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestLinter_Issues(t *testing.T) {
	t.Parallel()
	type want struct {
		issues []issue.Issue
		panic  bool
	}
	tests := []struct {
		name string
		n    *impl.Name
		want want
	}{
		{
			name: "ok",
			n: impl.NewName(
				branch.NewName("TASK-1234"),
				name.NewConfig(
					value.NewString(
						issue.Critical,
						`(TASK|PROJ|BUG)-[0-9]+`,
					),
				),
			),
			want: want{
				issues: []issue.Issue{},
			},
		},
		{
			name: "critical by pattern",
			n: impl.NewName(
				branch.NewName("FEATURE-1234"),
				name.NewConfig(
					value.NewString(
						issue.Critical,
						`(TASK|PROJ|BUG)-[0-9]+`,
					),
				),
			),
			want: want{
				issues: []issue.Issue{
					issue.NewCritical("branch name doesn't match the required pattern '(TASK|PROJ|BUG)-[0-9]+'"),
				},
			},
		},
		{
			name: "warning by pattern",
			n: impl.NewName(
				branch.NewName("FEATURE-1234"),
				name.NewConfig(
					value.NewString(
						issue.Warning,
						`(TASK|PROJ|BUG)-[0-9]+`,
					),
				),
			),
			want: want{
				issues: []issue.Issue{
					issue.NewWarning("branch name doesn't match the required pattern '(TASK|PROJ|BUG)-[0-9]+'"),
				},
			},
		},
		{
			name: "empty pattern",
			n: impl.NewName(
				branch.NewName("FEATURE-1234"),
				name.NewConfig(
					value.NewString(
						issue.Warning,
						``,
					),
				),
			),
			want: want{
				issues: []issue.Issue{},
			},
		},
		{
			name: "critical empty branch name",
			n: impl.NewName(
				branch.NewName(""),
				name.NewConfig(
					value.NewString(
						issue.Critical,
						`(TASK|PROJ|BUG)-[0-9]+`,
					),
				),
			),
			want: want{
				issues: []issue.Issue{
					issue.NewCritical("branch name doesn't match the required pattern '(TASK|PROJ|BUG)-[0-9]+'"),
				},
			},
		},
		{
			name: "critical empty branch name & empty pattern",
			n: impl.NewName(
				branch.NewName(""),
				name.NewConfig(
					value.NewString(
						issue.Critical,
						``,
					),
				),
			),
			want: want{
				issues: []issue.Issue{},
			},
		},
		{
			name: "empty impl",
			n:    &impl.Name{},
			want: want{
				issues: []issue.Issue{},
			},
		},
		{
			name: "nil impl",
			n:    nil,
			want: want{
				issues: nil,
				panic:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.n.Issues()
				})
				return
			}
			assert.Equal(
				t,
				tt.want.issues,
				tt.n.Issues(),
			)
		})
	}
}
