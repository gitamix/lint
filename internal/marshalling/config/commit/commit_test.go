package commit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/commit"
	"github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/config/commit/message/body"
	"github.com/gitamix/lint/config/commit/message/body/mandate"
	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/commit/types"
	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/commit"
	tmessage "github.com/gitamix/lint/internal/marshalling/config/commit/message"
	tbody "github.com/gitamix/lint/internal/marshalling/config/commit/message/body"
	tmandate "github.com/gitamix/lint/internal/marshalling/config/commit/message/body/mandate"
	tsubject "github.com/gitamix/lint/internal/marshalling/config/commit/message/subject"
	tdescription "github.com/gitamix/lint/internal/marshalling/config/commit/message/subject/description"
	mvalue "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestCommit_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero config when empty", func(t *testing.T) {
		t.Parallel()
		c := impl.Commit{}
		want := commit.NewConfig()
		assert.Equal(t, want, c.Config())
	})

	t.Run("converts message config into commit config", func(t *testing.T) {
		t.Parallel()
		c := impl.Commit{
			Msg: tmessage.Message{
				Subj: tsubject.Subject{
					Desc: tdescription.Description{
						Length: mvalue.Range{
							Issue: mvalue.Issue{
								Level: "info",
							},
							Min: 1,
							Max: 72,
						},
					},
				},
				Body: tbody.Body{
					Mandate: tmandate.Mandate{
						Types: mvalue.Strings{
							List: []string{"feat", "fix"},
							Issue: mvalue.Issue{
								Level: "critical",
							},
						},
					},
					Length: mvalue.Range{
						Issue: mvalue.Issue{
							Level: "warning",
						},
						Min: 1,
						Max: 100,
					},
				},
			},
		}
		want := commit.NewConfig(
			commit.WithMessage(
				message.NewConfig(
					message.WithSubject(
						subject.NewConfig(
							subject.WithDescription(
								description.NewConfig(
									description.WithLength(
										value.NewRange(issue.Info, 1, 72),
									),
								),
							),
						),
					),
					message.WithBody(
						body.NewConfig(
							body.WithMandate(
								mandate.NewConfig(
									mandate.WithTypes(
										value.NewStrings(issue.Critical, "feat", "fix"),
									),
								),
							),
							body.WithLength(
								value.NewRange(issue.Warning, 1, 100),
							),
						),
					),
				),
			),
		)
		assert.Equal(t, want, c.Config())
	})

	t.Run("converts scope config into commit config", func(t *testing.T) {
		t.Parallel()
		c := impl.Commit{
			Scope: mvalue.Pattern{
				Issue: mvalue.Issue{
					Level: "critical",
				},
				Pattern: "feat|fix",
			},
		}
		want := commit.NewConfig(
			commit.WithScope(
				scope.NewConfig(
					value.NewString(issue.Critical, "feat|fix"),
				),
			),
		)
		assert.Equal(t, want, c.Config())
	})

	t.Run("converts types config into commit config", func(t *testing.T) {
		t.Parallel()
		c := impl.Commit{
			Types: mvalue.Strings{
				List: []string{"feat", "fix"},
				Issue: mvalue.Issue{
					Level: "info",
				},
			},
		}
		want := commit.NewConfig(
			commit.WithTypes(
				types.NewConfig(
					value.NewStrings(issue.Info, "feat", "fix"),
				),
			),
		)
		assert.Equal(t, want, c.Config())
	})

	t.Run("sets warning level for scope when it is not set", func(t *testing.T) {
		t.Parallel()
		c := impl.Commit{
			Scope: mvalue.Pattern{
				Pattern: "feat|fix",
			},
		}
		assert.Equal(
			t,
			issue.Warning,
			c.Config().
				Scope().
				Pattern().
				Level(),
		)
	})

	t.Run("keeps unknown level for scope when it is incorrect", func(t *testing.T) {
		t.Parallel()
		c := impl.Commit{
			Scope: mvalue.Pattern{
				Issue: mvalue.Issue{
					Level: "foo",
				},
				Pattern: "feat|fix",
			},
		}
		assert.True(
			t,
			c.Config().
				Scope().
				Pattern().
				Level().
				Unknown(),
		)
	})

	t.Run("sets critical level for types when issue is not set", func(t *testing.T) {
		t.Parallel()
		c := impl.Commit{
			Types: mvalue.Strings{
				List: []string{"feat", "fix"},
			},
		}
		assert.Equal(
			t,
			issue.Critical,
			c.Config().
				Types().
				List().
				Level(),
		)
	})

	t.Run("keeps unknown level for types when it is incorrect", func(t *testing.T) {
		t.Parallel()
		c := impl.Commit{
			Types: mvalue.Strings{
				Issue: mvalue.Issue{
					Level: "foo",
				},
				List: []string{"feat", "fix"},
			},
		}
		assert.True(
			t,
			c.Config().
				Types().
				List().
				Level().
				Unknown(),
		)
	})
}
