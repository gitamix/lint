package message_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/config/commit/message/body"
	"github.com/gitamix/lint/config/commit/message/body/mandate"
	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/commit/message"
	tbody "github.com/gitamix/lint/internal/marshalling/config/commit/message/body"
	tmandate "github.com/gitamix/lint/internal/marshalling/config/commit/message/body/mandate"
	tsubject "github.com/gitamix/lint/internal/marshalling/config/commit/message/subject"
	tdescription "github.com/gitamix/lint/internal/marshalling/config/commit/message/subject/description"
	mvalue "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestMessage_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero config when empty", func(t *testing.T) {
		t.Parallel()
		m := impl.Message{}
		want := message.NewConfig()
		assert.Equal(t, want, m.Config())
	})

	t.Run("converts subject config into message config", func(t *testing.T) {
		t.Parallel()
		m := impl.Message{
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
		}
		want := message.NewConfig(
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
		)
		assert.Equal(t, want, m.Config())
	})

	t.Run("converts body config into message config", func(t *testing.T) {
		t.Parallel()
		m := impl.Message{
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
		}
		want := message.NewConfig(
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
		)
		assert.Equal(t, want, m.Config())
	})
}
