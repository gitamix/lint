package value_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestIssue_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero type when empty", func(t *testing.T) {
		t.Parallel()
		i := impl.Issue{}
		assert.True(t, i.Config().Unspecified())
	})

	t.Run("parses critical level", func(t *testing.T) {
		t.Parallel()
		i := impl.Issue{
			Level: "critical",
		}
		assert.Equal(t, issue.Critical, i.Config())
	})

	t.Run("parses warning level", func(t *testing.T) {
		t.Parallel()
		i := impl.Issue{
			Level: "warning",
		}
		assert.Equal(t, issue.Warning, i.Config())
	})

	t.Run("parses info level", func(t *testing.T) {
		t.Parallel()
		i := impl.Issue{
			Level: "info",
		}
		assert.Equal(t, issue.Info, i.Config())
	})

	t.Run("keeps unknown level when it is incorrect", func(t *testing.T) {
		t.Parallel()
		i := impl.Issue{
			Level: "foo",
		}
		assert.True(t, i.Config().Unknown())
	})
}
