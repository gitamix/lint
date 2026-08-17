package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/commit/types"
	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/commit/types"
	mvalue "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestTypes_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero config when empty", func(t *testing.T) {
		t.Parallel()
		tt := impl.Types{}
		want := types.NewConfig(value.Strings{})
		assert.Equal(t, want, tt.Config())
	})

	t.Run("converts types into types config", func(t *testing.T) {
		t.Parallel()
		tt := impl.Types{
			Types: mvalue.Strings{
				List: []string{"feat", "fix"},
				Issue: mvalue.Issue{
					Level: "info",
				},
			},
		}
		want := types.NewConfig(
			value.NewStrings(
				issue.Info,
				"feat",
				"fix",
			),
		)
		assert.Equal(t, want, tt.Config())
	})

	t.Run("keeps unspecified level when issue is not set", func(t *testing.T) {
		t.Parallel()
		tt := impl.Types{
			Types: mvalue.Strings{
				List: []string{"feat", "fix"},
			},
		}
		assert.True(
			t,
			tt.Config().
				Types().
				Level().
				Unspecified(),
		)
	})
}
