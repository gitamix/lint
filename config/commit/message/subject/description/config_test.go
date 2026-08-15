package description_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/value"
)

func TestConfig_Length(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.NewConfig().Length()
		want := value.Range{}
		assert.Equal(t, want, got)
	})

	t.Run("only with length", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(1, 100),
				),
			).
			Length()
		want := value.NewRange(1, 100)
		assert.Equal(t, want, got)
	})

	t.Run("with zero bounds", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(0, 0),
				),
			).
			Length()
		want := value.NewRange(0, 0)
		assert.Equal(t, want, got)
	})

	t.Run("last WithLength wins", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(1, 10),
				),
				impl.WithLength(
					value.NewRange(10, 72),
				),
			).
			Length()
		want := value.NewRange(10, 72)
		assert.Equal(t, want, got)
	})
}
