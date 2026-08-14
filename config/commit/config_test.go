package commit_test

import (
	"testing"

	impl "github.com/gitamix/lint/config/commit"
	"github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/commit/scope"
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

func TestConfig_Scope(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig().
			Scope()
		var want scope.Config
		assert.Equal(t, want, got)
	})

	t.Run("with all options", func(t *testing.T) {
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
									),
								),
							),
						),
					),
				),
				impl.WithScope(
					scope.NewConfig(
						value.NewString(
							issue.Critical,
							`^[A-Za-z _-]+$`,
						),
					),
				),
			).
			Scope()
		want := scope.NewConfig(
			value.NewString(
				issue.Critical,
				`^[A-Za-z _-]+$`,
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("only with message does not set scope", func(t *testing.T) {
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
			Scope()
		var want scope.Config
		assert.Equal(t, want, got)
	})

	t.Run("only with scope", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithScope(
					scope.NewConfig(
						value.NewString(
							issue.Critical,
							`^[A-Za-z _-]+$`,
						),
					),
				),
			).
			Scope()
		want := scope.NewConfig(
			value.NewString(
				issue.Critical,
				`^[A-Za-z _-]+$`,
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("last WithScope wins", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithScope(
					scope.NewConfig(
						value.NewString(
							issue.Critical,
							`^[A-Za-z _-]+$`,
						),
					),
				),
				impl.WithScope(
					scope.NewConfig(
						value.NewString(
							issue.Warning,
							`^[A-Za-z]+`,
						),
					),
				),
			).
			Scope()
		want := scope.NewConfig(
			value.NewString(
				issue.Warning,
				`^[A-Za-z]+`,
			),
		)
		assert.Equal(t, want, got)
	})
}
