//go:build integration
// +build integration

package current_test

import (
	"context"
	"testing"
	"time"

	"github.com/gitamix/git/client/git"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	lintbranch "github.com/gitamix/lint/branch/current"
	"github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/ticket"
	"github.com/gitamix/lint/config/ticket/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/internal/test/container"
	"github.com/gitamix/lint/internal/test/fixture/container/shared"
	shfx "github.com/gitamix/lint/internal/test/fixture/shell"
	"github.com/gitamix/lint/issue"
)

// IssuesSuite groups integration tests for branch current validation.
type IssuesSuite struct {
	suite.Suite
	// git is the client to interact with git repositories.
	git *git.Client
}

// SetupSuite initializes the shared container fixture and creates the git client.
func (s *IssuesSuite) SetupSuite() {
	s.git = git.NewClient(
		shfx.NewShell(
			shared.
				ContainerFixture().
				Container(),
			container.RepoDir,
		),
	)
}

// TestBranch_Issues runs all branch-current validation tests.
func TestBranch_Issues(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(IssuesSuite))
}

func (s *IssuesSuite) TestOk() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()
	currbr := lintbranch.NewBranch(
		s.git,
		branch.NewConfig(
			branch.WithTicket(
				ticket.NewConfig(
					ticket.WithID(
						id.NewConfig(
							value.NewString(
								issue.Info,
								`(TASK|PROJ|BUG)-[0-9]+`,
							),
						),
					),
				),
			),
			branch.WithName(
				name.NewConfig(
					value.NewString(
						issue.Warning,
						`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
					),
				),
			),
		),
	)
	got, gotErr := currbr.Issues(ctx)
	require.NoError(s.T(), gotErr)
	assert.Equal(s.T(), []issue.Issue{}, got)
}

func (s *IssuesSuite) TestNotMatchedWithPattern() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()
	currbr := lintbranch.NewBranch(
		s.git,
		branch.NewConfig(
			branch.WithTicket(
				ticket.NewConfig(
					ticket.WithID(
						id.NewConfig(
							value.NewString(
								issue.Info,
								`([0-9]+)-\w`,
							),
						),
					),
				),
			),
			branch.WithName(
				name.NewConfig(
					value.NewString(
						issue.Warning,
						`^(feature|release)/\w+`,
					),
				),
			),
		),
	)
	got, gotErr := currbr.Issues(ctx)
	require.NoError(s.T(), gotErr)
	assert.Equal(s.T(),
		[]issue.Issue{
			issue.NewWarning("branch name 'bugfix/BUG-456' doesn't match the required pattern '^(feature|release)/\\w+'"),
			issue.NewInfo("ticket doesn't match the required pattern '([0-9]+)-\\w'"),
		},
		got,
	)
}
