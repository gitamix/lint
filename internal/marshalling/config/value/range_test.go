package value_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestRange_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero range value when empty", func(t *testing.T) {
		t.Parallel()
		r := impl.Range{}
		want := value.Range{}
		assert.Equal(t, want, r.Config())
	})

	t.Run("builds range value from issue level and bounds", func(t *testing.T) {
		t.Parallel()
		r := impl.Range{
			Issue: impl.Issue{
				Level: "info",
			},
			Min: 1,
			Max: 5,
		}
		want := value.NewRange(issue.Info, 1, 5)
		assert.Equal(t, want, r.Config())
	})

	t.Run("keeps zero level when issue is not set", func(t *testing.T) {
		t.Parallel()
		r := impl.Range{
			Min: 1,
			Max: 5,
		}
		want := value.NewRange(issue.Type(0), 1, 5)
		assert.Equal(t, want, r.Config())
	})

	t.Run("defaults unknown level to warning", func(t *testing.T) {
		t.Parallel()
		r := impl.Range{
			Issue: impl.Issue{Level: "bogus"},
			Min:   1,
			Max:   5,
		}
		want := value.NewRange(issue.Warning, 1, 5)
		assert.Equal(t, want, r.Config())
	})

	t.Run("builds zero bounds range when only issue is set", func(t *testing.T) {
		t.Parallel()
		r := impl.Range{
			Issue: impl.Issue{
				Level: "critical",
			},
		}
		want := value.NewRange(issue.Critical, 0, 0)
		assert.Equal(t, want, r.Config())
	})
}
