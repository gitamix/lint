package message_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/config/commit/message/body"
	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/ticket"
	"github.com/gitamix/lint/config/ticket/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
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

	t.Run("with subject length", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(
						subject.WithLength(
							value.NewRange(10, 72),
						),
					),
				),
			).
			Subject()
		want := subject.NewConfig(
			subject.WithLength(
				value.NewRange(10, 72),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with subject config without options", func(t *testing.T) {
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

	t.Run("with subject ticket", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
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
			).
			Subject()
		want := subject.NewConfig(
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
		)
		assert.Equal(t, want, got)
	})

	t.Run("with subject length and ticket", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
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
			).
			Subject()
		want := subject.NewConfig(
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
		)
		assert.Equal(t, want, got)
	})

	t.Run("with all options & subject is unaffected by body", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(
						subject.WithLength(
							value.NewRange(10, 72),
						),
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
				impl.WithBody(
					body.NewConfig(
						body.WithLength(
							value.NewRange(20, 255),
						),
					),
				),
			).
			Subject()
		want := subject.NewConfig(
			subject.WithLength(
				value.NewRange(10, 72),
			),
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
		)
		assert.Equal(t, want, got)
	})

	t.Run("default subject stays default when body is set", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(),
				),
				impl.WithBody(
					body.NewConfig(
						body.WithLength(
							value.NewRange(20, 255),
						),
					),
				),
			).
			Subject()
		want := subject.NewConfig()
		assert.Equal(t, want, got)
	})
}

func TestConfig_Body(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig().
			Body()
		want := body.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("with body length", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithBody(
					body.NewConfig(
						body.WithLength(
							value.NewRange(20, 255),
						),
					),
				),
			).
			Body()
		want := body.NewConfig(
			body.WithLength(
				value.NewRange(20, 255),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with body config without options", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithBody(
					body.NewConfig(),
				),
			).
			Body()
		want := body.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("with body zero bounds", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithBody(
					body.NewConfig(
						body.WithLength(
							value.NewRange(0, 0),
						),
					),
				),
			).
			Body()
		want := body.NewConfig(
			body.WithLength(
				value.NewRange(0, 0),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("last WithBody wins", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithBody(
					body.NewConfig(
						body.WithLength(
							value.NewRange(1, 10),
						),
					),
				),
				impl.WithBody(
					body.NewConfig(
						body.WithLength(
							value.NewRange(20, 255),
						),
					),
				),
			).
			Body()
		want := body.NewConfig(
			body.WithLength(
				value.NewRange(20, 255),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("with all options & body is unaffected by subject", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(
						subject.WithLength(
							value.NewRange(10, 72),
						),
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
				impl.WithBody(
					body.NewConfig(
						body.WithLength(
							value.NewRange(20, 255),
						),
					),
				),
			).
			Body()
		want := body.NewConfig(
			body.WithLength(
				value.NewRange(20, 255),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("default body stays default when subject is set", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithSubject(
					subject.NewConfig(
						subject.WithLength(
							value.NewRange(10, 72),
						),
					),
				),
				impl.WithBody(
					body.NewConfig(),
				),
			).
			Body()
		want := body.NewConfig()
		assert.Equal(t, want, got)
	})
}
