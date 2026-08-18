package value_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestStrings_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero strings value when empty", func(t *testing.T) {
		t.Parallel()
		s := impl.Strings{}
		want := value.Strings{}
		assert.Equal(t, want, s.Config())
	})

	t.Run("builds strings value from issue level and list", func(t *testing.T) {
		t.Parallel()
		s := impl.Strings{
			List: []string{"feat", "fix"},
			Issue: impl.Issue{
				Level: "critical",
			},
		}
		want := value.NewStrings(issue.Critical, "feat", "fix")
		assert.Equal(t, want, s.Config())
	})

	t.Run("keeps unspecified level when issue is not set", func(t *testing.T) {
		t.Parallel()
		s := impl.Strings{
			List: []string{"feat"},
		}
		assert.True(t, s.Config().Level().Unspecified())
	})

	t.Run("keeps unknown level when it is incorrect", func(t *testing.T) {
		t.Parallel()
		s := impl.Strings{
			List: []string{"feat"},
			Issue: impl.Issue{
				Level: "foo",
			},
		}
		assert.True(
			t, s.
				Config().
				Level().
				Unknown(),
		)
	})

	t.Run("builds empty list when only issue is set", func(t *testing.T) {
		t.Parallel()
		s := impl.Strings{
			Issue: impl.Issue{
				Level: "info",
			},
		}
		want := value.NewStrings(issue.Info)
		assert.Equal(t, want, s.Config())
	})
}
