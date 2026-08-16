package value_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestPattern_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero string value when empty", func(t *testing.T) {
		t.Parallel()
		p := impl.Pattern{}
		want := value.String{}
		assert.Equal(t, want, p.Config())
	})

	t.Run("builds string value from issue level and pattern", func(t *testing.T) {
		t.Parallel()
		p := impl.Pattern{
			Issue:   impl.Issue{
				Level: "critical",
			},
			Pattern: "feat",
		}
		want := value.NewString(issue.Critical, "feat")
		assert.Equal(t, want, p.Config())
	})

	t.Run("keeps zero level when issue is not set", func(t *testing.T) {
		t.Parallel()
		p := impl.Pattern{
			Pattern: "feat",
		}
		want := value.NewString(issue.Type(0), "feat")
		assert.Equal(t, want, p.Config())
	})

	t.Run("defaults unknown level to warning", func(t *testing.T) {
		t.Parallel()
		p := impl.Pattern{
			Issue:   impl.Issue{Level: "bogus"},
			Pattern: "feat",
		}
		want := value.NewString(issue.Warning, "feat")
		assert.Equal(t, want, p.Config())
	})

	t.Run("keeps empty pattern when only issue is set", func(t *testing.T) {
		t.Parallel()
		p := impl.Pattern{
			Issue: impl.Issue{
				Level: "info",
			},
		}
		want := value.NewString(issue.Info, "")
		assert.Equal(t, want, p.Config())
	})
}
