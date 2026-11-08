package description_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
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
					value.NewRange(issue.Warning, 1, 100),
				),
			).
			Length()
		want := value.NewRange(issue.Warning, 1, 100)
		assert.Equal(t, want, got)
	})

	t.Run("with zero bounds", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(issue.Warning, 0, 0),
				),
			).
			Length()
		want := value.NewRange(issue.Warning, 0, 0)
		assert.Equal(t, want, got)
	})

	t.Run("last WithLength wins", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(issue.Warning, 1, 10),
				),
				impl.WithLength(
					value.NewRange(issue.Warning, 10, 72),
				),
			).
			Length()
		want := value.NewRange(issue.Warning, 10, 72)
		assert.Equal(t, want, got)
	})
}

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	t.Run("ok without any option", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.NewConfig().Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok with valid length", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(issue.Warning, 1, 100),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok with equal bounds", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(issue.Critical, 10, 10),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("err with unknown issue level", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(issue.Type(-1), 1, 100),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid config: unknown issue level",
		)
	})

	t.Run("err with negative min", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(issue.Warning, -1, 100),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid config: min is negative",
		)
	})

	t.Run("err with negative max", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(issue.Warning, 0, -1),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid config: max is negative",
		)
	})

	t.Run("err when min is greater than max", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(issue.Warning, 100, 1),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid config: min is greater than max",
		)
	})
}
