package id_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/task/id"
	mvalue "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestID_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero config when empty", func(t *testing.T) {
		t.Parallel()
		i := impl.ID{}
		want := id.NewConfig(value.String{})
		assert.Equal(t, want, i.Config())
	})

	t.Run("converts pattern into id config", func(t *testing.T) {
		t.Parallel()
		i := impl.ID{
			Pattern: mvalue.Pattern{
				Issue: mvalue.Issue{
					Level: "info",
				},
				Pattern: "TMS-\\d+",
			},
		}
		want := id.NewConfig(
			value.NewString(issue.Info, "TMS-\\d+"),
		)
		assert.Equal(t, want, i.Config())
	})

	t.Run("sets critical level when pattern is set without level", func(t *testing.T) {
		t.Parallel()
		i := impl.ID{
			Pattern: mvalue.Pattern{
				Pattern: "TMS-\\d+",
			},
		}
		want := id.NewConfig(
			value.NewString(issue.Critical, "TMS-\\d+"),
		)
		assert.Equal(t, want, i.Config())
	})

	t.Run("keeps unknown level when it is incorrect", func(t *testing.T) {
		t.Parallel()
		i := impl.ID{
			Pattern: mvalue.Pattern{
				Issue: mvalue.Issue{
					Level: "foo",
				},
				Pattern: "TMS-\\d+",
			},
		}
		assert.True(
			t,
			i.Config().
				Pattern().
				Level().
				Unknown(),
		)
	})
}
