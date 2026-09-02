package description_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/commit/message/subject/description"
	mvalue "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestDescription_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero config when empty", func(t *testing.T) {
		t.Parallel()
		d := impl.Description{}
		want := description.NewConfig()
		assert.Equal(t, want, d.Config())
	})

	t.Run("converts length range into description config", func(t *testing.T) {
		t.Parallel()
		d := impl.Description{
			Length: mvalue.Range{
				Issue: mvalue.Issue{
					Level: "info",
				},
				Min: 1,
				Max: 72,
			},
		}
		want := description.NewConfig(
			description.WithLength(
				value.NewRange(issue.Info, 1, 72),
			),
		)
		assert.Equal(t, want, d.Config())
	})

	t.Run("sets warning level when it is not set", func(t *testing.T) {
		t.Parallel()
		d := impl.Description{
			Length: mvalue.Range{
				Min: 1,
				Max: 72,
			},
		}
		want := description.NewConfig(
			description.WithLength(
				value.NewRange(issue.Warning, 1, 72),
			),
		)
		assert.Equal(t, want, d.Config())
	})

	t.Run("keeps unknown level for length when it is incorrect", func(t *testing.T) {
		t.Parallel()
		d := impl.Description{
			Length: mvalue.Range{
				Issue: mvalue.Issue{
					Level: "foo",
				},
				Min: 1,
				Max: 72,
			},
		}
		assert.True(
			t,
			d.Config().
				Length().
				Level().
				Unknown(),
		)
	})
}
