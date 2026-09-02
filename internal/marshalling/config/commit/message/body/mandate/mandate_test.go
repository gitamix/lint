package mandate_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/commit/message/body/mandate"
	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/commit/message/body/mandate"
	mvalue "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestMandate_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero config when empty", func(t *testing.T) {
		t.Parallel()
		m := impl.Mandate{}
		want := mandate.NewConfig()
		assert.Equal(t, want, m.Config())
	})

	t.Run("converts commit types into mandate config", func(t *testing.T) {
		t.Parallel()
		m := impl.Mandate{
			Types: mvalue.Strings{
				List: []string{"feat", "fix"},
				Issue: mvalue.Issue{
					Level: "critical",
				},
			},
		}
		want := mandate.NewConfig(
			mandate.WithTypes(
				value.NewStrings(issue.Critical, "feat", "fix"),
			),
		)
		assert.Equal(t, want, m.Config())
	})
}
