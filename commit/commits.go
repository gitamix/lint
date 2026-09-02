package commit

import (
	"github.com/gitamix/types/commit"

	config "github.com/gitamix/lint/config/commit"
	"github.com/gitamix/lint/issue"
)

// Commits represents a linter that validates a collection of git commits
// against the configured commit message rules.
type Commits struct {
	// commits is the collection of parsed git commits inspected by the linter.
	commits []commit.Commit

	// cfg is the configuration that defines the rules
	// applied to each commit message during the lint run.
	cfg config.Config
}

// NewCommits creates a new Commits linter
// with the provided git commits and lint configuration.
func NewCommits(
	commits []commit.Commit,
	cfg config.Config,
) Commits {
	return Commits{
		commits: commits,
		cfg:     cfg,
	}
}

// Issues returns a slice of issues describing
// any validation problems found across all the commits.
//
// It aggregates issues from each commit by building a Commit linter
// for every entry and collecting the issues it reports.
func (c Commits) Issues() []issue.Issue {
	issues := make([]issue.Issue, 0, 5*len(c.commits))
	for _, cmt := range c.commits {
		for _, iss := range NewCommit(cmt, c.cfg).Issues() {
			issues = append(
				issues,
				issue.NewIssue(
					iss.Type(),
					cmt.Hash().ShortString()+": "+iss.Message(),
				),
			)
		}
	}
	return issues
}
