package name_test

import (
	"testing"

	"github.com/gitamix/types/branch"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/branch/name"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestLinter_Issues(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		n := impl.NewName(
			branch.NewName("TASK-1234"),
			name.NewConfig(
				value.NewString(
					issue.Critical,
					`(TASK|PROJ|BUG)-[0-9]+`,
				),
			),
		)
		assert.Equal(
			t,
			[]issue.Issue{},
			n.Issues(),
		)
	})

	t.Run("critical by pattern", func(t *testing.T) {
		t.Parallel()
		n := impl.NewName(
			branch.NewName("FEATURE-1234"),
			name.NewConfig(
				value.NewString(
					issue.Critical,
					`(TASK|PROJ|BUG)-[0-9]+`,
				),
			),
		)
		assert.Equal(
			t,
			[]issue.Issue{
				issue.NewCritical("branch name 'FEATURE-1234' doesn't match the required pattern '(TASK|PROJ|BUG)-[0-9]+'"),
			},
			n.Issues(),
		)
	})

	t.Run("warning by pattern", func(t *testing.T) {
		t.Parallel()
		n := impl.NewName(
			branch.NewName("FEATURE-1234"),
			name.NewConfig(
				value.NewString(
					issue.Warning,
					`(TASK|PROJ|BUG)-[0-9]+`,
				),
			),
		)
		assert.Equal(
			t,
			[]issue.Issue{
				issue.NewWarning("branch name 'FEATURE-1234' doesn't match the required pattern '(TASK|PROJ|BUG)-[0-9]+'"),
			},
			n.Issues(),
		)
	})

	t.Run("empty pattern", func(t *testing.T) {
		t.Parallel()
		n := impl.NewName(
			branch.NewName("FEATURE-1234"),
			name.NewConfig(
				value.NewString(
					issue.Warning,
					``,
				),
			),
		)
		assert.Equal(
			t,
			[]issue.Issue{},
			n.Issues(),
		)
	})

	t.Run("critical empty branch name", func(t *testing.T) {
		t.Parallel()
		n := impl.NewName(
			branch.NewName(""),
			name.NewConfig(
				value.NewString(
					issue.Critical,
					`(TASK|PROJ|BUG)-[0-9]+`,
				),
			),
		)
		assert.Equal(
			t,
			[]issue.Issue{
				issue.NewCritical("branch name '' doesn't match the required pattern '(TASK|PROJ|BUG)-[0-9]+'"),
			},
			n.Issues(),
		)
	})

	t.Run("critical empty branch name & empty pattern", func(t *testing.T) {
		t.Parallel()
		n := impl.NewName(
			branch.NewName(""),
			name.NewConfig(
				value.NewString(
					issue.Critical,
					``,
				),
			),
		)
		assert.Equal(
			t,
			[]issue.Issue{},
			n.Issues(),
		)
	})

	t.Run("empty impl", func(t *testing.T) {
		t.Parallel()
		n := &impl.Name{}
		assert.Equal(
			t,
			[]issue.Issue{},
			n.Issues(),
		)
	})

	t.Run("nil impl panics", func(t *testing.T) {
		t.Parallel()
		var n *impl.Name
		assert.Panics(t, func() {
			_ = n.Issues()
		})
	})
}
