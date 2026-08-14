package message_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/commit/message"
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
}
