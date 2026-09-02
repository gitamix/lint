package body_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/commit/message/body"
	"github.com/gitamix/lint/config/commit/message/body/mandate"
	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/commit/message/body"
	tmandate "github.com/gitamix/lint/internal/marshalling/config/commit/message/body/mandate"
	mvalue "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestBody_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero config when empty", func(t *testing.T) {
		t.Parallel()
		b := impl.Body{}
		want := body.NewConfig()
		assert.Equal(t, want, b.Config())
	})

	t.Run("converts mandate config into body config", func(t *testing.T) {
		t.Parallel()
		b := impl.Body{
			Mandate: tmandate.Mandate{
				Types: mvalue.Strings{
					List: []string{"feat", "fix"},
					Issue: mvalue.Issue{
						Level: "critical",
					},
				},
			},
		}
		want := body.NewConfig(
			body.WithMandate(
				mandate.NewConfig(
					mandate.WithTypes(
						value.NewStrings(issue.Critical, "feat", "fix"),
					),
				),
			),
		)
		assert.Equal(t, want, b.Config())
	})

	t.Run("converts length range into body config", func(t *testing.T) {
		t.Parallel()
		b := impl.Body{
			Length: mvalue.Range{
				Issue: mvalue.Issue{
					Level: "warning",
				},
				Min: 1,
				Max: 100,
			},
		}
		want := body.NewConfig(
			body.WithLength(
				value.NewRange(issue.Warning, 1, 100),
			),
		)
		assert.Equal(t, want, b.Config())
	})

	t.Run("keeps unspecified lvl for body length when it is not set", func(t *testing.T) {
		t.Parallel()
		b := impl.Body{
			Length: mvalue.Range{
				Min: 1,
				Max: 100,
			},
		}
		assert.True(
			t,
			b.Config().
				Length().
				Level().
				Unspecified(),
		)
	})

	t.Run("keeps unknown lvl for body length when it is incorrect", func(t *testing.T) {
		t.Parallel()
		b := impl.Body{
			Length: mvalue.Range{
				Issue: mvalue.Issue{
					Level: "foo",
				},
				Min: 1,
				Max: 100,
			},
		}
		assert.True(
			t,
			b.Config().
				Length().
				Level().
				Unknown(),
		)
	})

	t.Run("mandate lvl does not affect length lvl on unspecified", func(t *testing.T) {
		t.Parallel()
		b := impl.Body{
			Mandate: tmandate.Mandate{
				Types: mvalue.Strings{
					List: []string{"feat", "fix"},
				},
			},
			Length: mvalue.Range{
				Issue: mvalue.Issue{
					Level: "warning",
				},
				Min: 1,
				Max: 100,
			},
		}
		assert.Equal(
			t,
			issue.Warning,
			b.Config().
				Length().
				Level(),
		)
	})
}
