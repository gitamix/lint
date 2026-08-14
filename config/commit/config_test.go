package commit_test

import (
	"testing"

	impl "github.com/gitamix/lint/config/commit"
	"github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
	"github.com/stretchr/testify/assert"
)

func TestConfig_Message(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig().
			Message()
		want := message.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("with message subject types", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithMessage(
					message.NewConfig(
						message.WithSubject(
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
					),
				),
			).
			Message()
		want := message.NewConfig(
			message.WithSubject(
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
		)
		assert.Equal(t, want, got)
	})

	t.Run("with message config without subject", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithMessage(
					message.NewConfig(),
				),
			).
			Message()
		want := message.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("with message subject types and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithMessage(
					message.NewConfig(
						message.WithSubject(
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
					),
				),
			).
			Message()
		want := message.NewConfig(
			message.WithSubject(
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
		)
		assert.Equal(t, want, got)
	})

	t.Run("with message subject types and info level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithMessage(
					message.NewConfig(
						message.WithSubject(
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
					),
				),
			).
			Message()
		want := message.NewConfig(
			message.WithSubject(
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
		)
		assert.Equal(t, want, got)
	})

	t.Run("with message subject empty types and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithMessage(
					message.NewConfig(
						message.WithSubject(
							subject.NewConfig(
								subject.WithTypes(
									value.NewStrings(issue.Critical),
								),
							),
						),
					),
				),
			).
			Message()
		want := message.NewConfig(
			message.WithSubject(
				subject.NewConfig(
					subject.WithTypes(
						value.NewStrings(issue.Critical),
					),
				),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with message subject empty types and warning level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithMessage(
					message.NewConfig(
						message.WithSubject(
							subject.NewConfig(
								subject.WithTypes(
									value.NewStrings(issue.Warning),
								),
							),
						),
					),
				),
			).
			Message()
		want := message.NewConfig(
			message.WithSubject(
				subject.NewConfig(
					subject.WithTypes(
						value.NewStrings(issue.Warning),
					),
				),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with message subject empty types and info level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithMessage(
					message.NewConfig(
						message.WithSubject(
							subject.NewConfig(
								subject.WithTypes(
									value.NewStrings(issue.Info),
								),
							),
						),
					),
				),
			).
			Message()
		want := message.NewConfig(
			message.WithSubject(
				subject.NewConfig(
					subject.WithTypes(
						value.NewStrings(issue.Info),
					),
				),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with message subject some empty types and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithMessage(
					message.NewConfig(
						message.WithSubject(
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
					),
				),
			).
			Message()
		want := message.NewConfig(
			message.WithSubject(
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
		)
		assert.Equal(t, want, got)
	})

	t.Run("with message subject single type", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithMessage(
					message.NewConfig(
						message.WithSubject(
							subject.NewConfig(
								subject.WithTypes(
									value.NewStrings(
										issue.Warning,
										"feat",
									),
								),
							),
						),
					),
				),
			).
			Message()
		want := message.NewConfig(
			message.WithSubject(
				subject.NewConfig(
					subject.WithTypes(
						value.NewStrings(
							issue.Warning,
							"feat",
						),
					),
				),
			),
		)
		assert.Equal(t, want, got)
	})
}
