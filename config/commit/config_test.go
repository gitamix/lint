package commit_test

import (
	"testing"

	impl "github.com/gitamix/lint/config/commit"
	"github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/commit/types"
	"github.com/gitamix/lint/config/ticket"
	"github.com/gitamix/lint/config/ticket/id"
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

	t.Run("with message subject length", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithMessage(
					message.NewConfig(
						message.WithSubject(
							subject.NewConfig(
								subject.WithLength(
									value.NewRange(10, 72),
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
					subject.WithLength(
						value.NewRange(10, 72),
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

	t.Run("with message subject ticket", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithMessage(
					message.NewConfig(
						message.WithSubject(
							subject.NewConfig(
								subject.WithTicket(
									ticket.NewConfig(
										ticket.WithID(
											id.NewConfig(
												value.NewString(
													issue.Critical,
													`^[A-Z]+-\d+$`,
												),
											),
										),
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
					subject.WithTicket(
						ticket.NewConfig(
							ticket.WithID(
								id.NewConfig(
									value.NewString(
										issue.Critical,
										`^[A-Z]+-\d+$`,
									),
								),
							),
						),
					),
				),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with message subject length and ticket", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithMessage(
					message.NewConfig(
						message.WithSubject(
							subject.NewConfig(
								subject.WithLength(
									value.NewRange(1, 100),
								),
								subject.WithTicket(
									ticket.NewConfig(
										ticket.WithID(
											id.NewConfig(
												value.NewString(
													issue.Warning,
													`^[A-Z]+-\d+$`,
												),
											),
										),
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
					subject.WithLength(
						value.NewRange(1, 100),
					),
					subject.WithTicket(
						ticket.NewConfig(
							ticket.WithID(
								id.NewConfig(
									value.NewString(
										issue.Warning,
										`^[A-Z]+-\d+$`,
									),
								),
							),
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
								subject.WithLength(
									value.NewRange(10, 72),
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
								subject.WithLength(
									value.NewRange(1, 100),
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

func TestConfig_Types(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig().
			Types()
		want := types.Config{}
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
								subject.WithLength(
									value.NewRange(10, 72),
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
				impl.WithTypes(
					types.NewConfig(
						value.NewStrings(
							issue.Warning,
							"feat",
							"fix",
						),
					),
				),
			).
			Types()
		want := types.NewConfig(
			value.NewStrings(
				issue.Warning,
				"feat",
				"fix",
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with types and warning level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					types.NewConfig(
						value.NewStrings(
							issue.Warning,
							"feat",
							"fix",
						),
					),
				),
			).
			Types()
		want := types.NewConfig(
			value.NewStrings(
				issue.Warning,
				"feat",
				"fix",
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with types and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					types.NewConfig(
						value.NewStrings(
							issue.Critical,
							"feat",
							"fix",
						),
					),
				),
			).
			Types()
		want := impl.
			NewConfig(
				impl.WithTypes(
					types.NewConfig(
						value.NewStrings(
							issue.Critical,
							"feat",
							"fix",
						),
					),
				),
			).
			Types()
		assert.Equal(t, want, got)
	})

	t.Run("with empty types and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					types.NewConfig(
						value.NewStrings(issue.Critical),
					),
				),
			).
			Types()
		want := types.NewConfig(
			value.NewStrings(issue.Critical),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with empty types and warning level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					types.NewConfig(
						value.NewStrings(issue.Warning),
					),
				),
			).
			Types()
		want := types.NewConfig(
			value.NewStrings(issue.Warning),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with empty types and info level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					types.NewConfig(
						value.NewStrings(issue.Info),
					),
				),
			).
			Types()
		want := types.NewConfig(
			value.NewStrings(issue.Info),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with some empty type and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					types.NewConfig(
						value.NewStrings(
							issue.Critical,
							"foo",
							"",
							"bar",
						),
					),
				),
			).
			Types()
		want := types.NewConfig(
			value.NewStrings(
				issue.Critical,
				"foo",
				"",
				"bar",
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with some empty type and warning level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					types.NewConfig(
						value.NewStrings(
							issue.Warning,
							"foo",
							"",
							"bar",
						),
					),
				),
			).
			Types()
		want := types.NewConfig(
			value.NewStrings(
				issue.Warning,
				"foo",
				"",
				"bar",
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with some empty type and info level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					types.NewConfig(
						value.NewStrings(
							issue.Info,
							"foo",
							"",
							"bar",
						),
					),
				),
			).
			Types()
		want := types.NewConfig(
			value.NewStrings(
				issue.Info,
				"foo",
				"",
				"bar",
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("only with message does not set types", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithMessage(
					message.NewConfig(
						message.WithSubject(
							subject.NewConfig(
								subject.WithLength(
									value.NewRange(10, 72),
								),
							),
						),
					),
				),
			).
			Types()
		want := types.Config{}
		assert.Equal(t, want, got)
	})

	t.Run("only with scope does not set types", func(t *testing.T) {
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
			Types()
		want := types.Config{}
		assert.Equal(t, want, got)
	})

	t.Run("last WithTypes wins", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					types.NewConfig(
						value.NewStrings(
							issue.Critical,
							"feat",
							"fix",
						),
					),
				),
				impl.WithTypes(
					types.NewConfig(
						value.NewStrings(
							issue.Warning,
							"docs",
							"style",
						),
					),
				),
			).
			Types()
		want := types.NewConfig(
			value.NewStrings(
				issue.Warning,
				"docs",
				"style",
			),
		)
		assert.Equal(t, want, got)
	})
}
