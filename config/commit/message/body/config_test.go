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

	t.Run("with length", func(t *testing.T) {
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
					value.NewRange(issue.Warning, 20, 255),
				),
			).
			Length()
		want := value.NewRange(issue.Warning, 20, 255)
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

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	t.Run("ok without any option", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.NewConfig().Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok with valid mandate and length", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithMandate(
					mandateImpl.NewConfig(
						mandateImpl.WithTypes(
							value.NewStrings(issue.Warning, "feat", "fix"),
						),
					),
				),
				impl.WithLength(
					value.NewRange(issue.Warning, 1, 100),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("err with unknown mandate issue level", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithMandate(
					mandateImpl.NewConfig(
						mandateImpl.WithTypes(
							value.NewStrings(issue.Type(-1), "feat"),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid mandate: invalid config: unknown issue level",
		)
	})

	t.Run("err with invalid length", func(t *testing.T) {
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
			"invalid length: invalid config: min is greater than max",
		)
	})

	t.Run("err for mandate even if length is invalid", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithMandate(
					mandateImpl.NewConfig(
						mandateImpl.WithTypes(
							value.NewStrings(issue.Type(-1), "feat"),
						),
					),
				),
				impl.WithLength(
					value.NewRange(issue.Warning, -1, 100),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid mandate: invalid config: unknown issue level",
		)
	})
}
