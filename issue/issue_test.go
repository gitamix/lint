package issue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/issue"
)

func TestIssue_Message(t *testing.T) {
	t.Parallel()

	t.Run("message with critical constructor", func(t *testing.T) {
		t.Parallel()
		i := impl.NewCritical("test message")
		assert.Equal(
			t,
			"test message",
			i.Message(),
		)
	})

	t.Run("message with warning constructor", func(t *testing.T) {
		t.Parallel()
		i := impl.NewWarning("test message")
		assert.Equal(
			t,
			"test message",
			i.Message(),
		)
	})

	t.Run("message with info constructor", func(t *testing.T) {
		t.Parallel()
		i := impl.NewInfo("test message")
		assert.Equal(
			t,
			"test message",
			i.Message(),
		)
	})
}

func TestIssue_Type(t *testing.T) {
	t.Parallel()

	t.Run("critical issue with its constructor", func(t *testing.T) {
		t.Parallel()
		i := impl.NewCritical("test message")
		assert.Equal(
			t,
			impl.Critical,
			i.Type(),
		)
	})

	t.Run("warning issue with its constructor", func(t *testing.T) {
		t.Parallel()
		i := impl.NewWarning("test message")
		assert.Equal(
			t,
			impl.Warning,
			i.Type(),
		)
	})

	t.Run("info issue with its constructor", func(t *testing.T) {
		t.Parallel()
		i := impl.NewInfo("test message")
		assert.Equal(
			t,
			impl.Info,
			i.Type(),
		)
	})

	t.Run("info issue with issue constructor", func(t *testing.T) {
		t.Parallel()
		i := impl.NewIssue(impl.Info, "test message")
		assert.Equal(
			t,
			impl.Info,
			i.Type(),
		)
	})
}
