package body_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/commit/message/body"
	mandateImpl "github.com/gitamix/lint/config/commit/message/body/mandate"
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

	t.Run("with all options", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(20, 255),
				),
			).
			Length()
		want := value.NewRange(20, 255)
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
					value.NewRange(20, 255),
				),
			).
			Length()
		want := value.NewRange(20, 255)
		assert.Equal(t, want, got)
	})
}

func TestConfig_Mandate(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.NewConfig().Mandate()
		want := mandateImpl.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("with mandate option", func(t *testing.T) {
		t.Parallel()
		want := mandateImpl.
			NewConfig(
				mandateImpl.WithTypes(
					value.NewStrings(
						issue.Warning,
						"fix",
						"perf",
					),
				),
			)
		got := impl.
			NewConfig(
				impl.WithMandate(want),
			).
			Mandate()
		assert.Equal(t, want, got)
	})

	t.Run("last WithMandate wins", func(t *testing.T) {
		t.Parallel()
		first := mandateImpl.
			NewConfig(
				mandateImpl.WithTypes(
					value.NewStrings(
						issue.Warning,
						"fix",
					),
				),
			)
		last := mandateImpl.
			NewConfig(
				mandateImpl.WithTypes(
					value.NewStrings(
						issue.Warning,
						"perf",
					),
				),
			)
		got := impl.
			NewConfig(
				impl.WithMandate(first),
				impl.WithMandate(last),
			).
			Mandate()
		assert.Equal(t, last, got)
	})
}
