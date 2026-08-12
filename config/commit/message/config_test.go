package message_test

import (
	"testing"

	impl "github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
	"github.com/stretchr/testify/assert"
)

func TestConfig_Subject(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig().
			Subject()
		want := subject.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("with subject types", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(
						subject.WithTypes(
							value.NewStrings(
								issue.Warning,
								"feat",
								"fix",
								"refactor",
							),
						),
					),
				),
			).
			Subject()
		want := subject.NewConfig(
			subject.WithTypes(
				value.NewStrings(
					issue.Warning,
					"feat",
					"fix",
					"refactor",
				),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with subject config without types", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(),
				),
			).
			Subject()
		want := subject.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("with subject types and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(
						subject.WithTypes(
							value.NewStrings(
								issue.Critical,
								"feat",
								"fix",
							),
						),
					),
				),
			).
			Subject()
		want := subject.NewConfig(
			subject.WithTypes(
				value.NewStrings(
					issue.Critical,
					"feat",
					"fix",
				),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with subject types and info level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(
						subject.WithTypes(
							value.NewStrings(
								issue.Info,
								"docs",
								"style",
							),
						),
					),
				),
			).
			Subject()
		want := subject.NewConfig(
			subject.WithTypes(
				value.NewStrings(
					issue.Info,
					"docs",
					"style",
				),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with subject empty types and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(
						subject.WithTypes(
							value.NewStrings(issue.Critical),
						),
					),
				),
			).
			Subject()
		want := subject.NewConfig(
			subject.WithTypes(
				value.NewStrings(issue.Critical),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with subject empty types and warning level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(
						subject.WithTypes(
							value.NewStrings(issue.Warning),
						),
					),
				),
			).
			Subject()
		want := subject.NewConfig(
			subject.WithTypes(
				value.NewStrings(issue.Warning),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with subject empty types and info level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(
						subject.WithTypes(
							value.NewStrings(issue.Info),
						),
					),
				),
			).
			Subject()
		want := subject.NewConfig(
			subject.WithTypes(
				value.NewStrings(issue.Info),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with subject some empty types and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(
						subject.WithTypes(
							value.NewStrings(
								issue.Critical,
								"feat",
								"",
								"fix",
							),
						),
					),
				),
			).
			Subject()
		want := subject.NewConfig(
			subject.WithTypes(
				value.NewStrings(
					issue.Critical,
					"feat",
					"",
					"fix",
				),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with subject single type", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(
						subject.WithTypes(
							value.NewStrings(
								issue.Warning,
								"feat",
							),
						),
					),
				),
			).
			Subject()
		want := subject.NewConfig(
			subject.WithTypes(
				value.NewStrings(
					issue.Warning,
					"feat",
				),
			),
		)
		assert.Equal(t, want, got)
	})
}
