package scope_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/value"
	impl "github.com/gitamix/lint/internal/marshalling/config/commit/scope"
	mvalue "github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

func TestScope_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero config when empty", func(t *testing.T) {
		t.Parallel()
		s := impl.Scope{}
		want := scope.NewConfig(value.String{})
		assert.Equal(t, want, s.Config())
	})

	t.Run("converts pattern into config", func(t *testing.T) {
		t.Parallel()
		s := impl.Scope{
			Pattern: mvalue.Pattern{
				Issue: mvalue.Issue{
					Level: "critical",
				},
				Pattern: "feat|fix",
			},
		}
		want := scope.NewConfig(
			value.NewString(
				issue.Critical,
				"feat|fix",
			),
		)
		assert.Equal(t, want, s.Config())
	})

	t.Run("sets warning level when it is not set", func(t *testing.T) {
		t.Parallel()
		s := impl.Scope{
			Pattern: mvalue.Pattern{
				Pattern: "feat|fix",
			},
		}
		want := scope.NewConfig(
			value.NewString(
				issue.Warning,
				"feat|fix",
			),
		)
		assert.Equal(t, want, s.Config())
	})
}
