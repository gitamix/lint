package subject_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/ticket"
	"github.com/gitamix/lint/config/ticket/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

func TestConfig_Types(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.NewConfig().Types()
		want := value.Strings{}
		assert.Equal(t, want, got)
	})

	t.Run("with all options and warning level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				impl.WithLength(
					value.NewRange(10, 72),
				),
				impl.WithTicket(
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
			).
			Types()
		want := value.NewStrings(
			issue.Warning,
			"feat",
			"fix",
		)
		assert.Equal(t, want, got)
	})

	t.Run("with types and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Critical,
						"feat",
						"fix",
					),
				),
			).
			Types()
		want := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Critical,
						"feat",
						"fix",
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
					value.NewStrings(issue.Critical),
				),
			).
			Types()
		want := value.NewStrings(issue.Critical)
		assert.Equal(t, want, got)
	})

	t.Run("with empty types and warning level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(issue.Warning),
				),
			).
			Types()
		want := value.NewStrings(issue.Warning)
		assert.Equal(t, want, got)
	})

	t.Run("with empty types and info level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(issue.Info),
				),
			).
			Types()
		want := value.NewStrings(issue.Info)
		assert.Equal(t, want, got)
	})

	t.Run("with some empty type and critical level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Critical,
						"foo",
						"",
						"bar",
					),
				),
			).
			Types()
		want := value.NewStrings(
			issue.Critical,
			"foo",
			"",
			"bar",
		)
		assert.Equal(t, want, got)
	})

	t.Run("with some empty type and warning level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Warning,
						"foo",
						"",
						"bar",
					),
				),
			).
			Types()
		want := value.NewStrings(
			issue.Warning,
			"foo",
			"",
			"bar",
		)
		assert.Equal(t, want, got)
	})

	t.Run("with some empty type and info level", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Info,
						"foo",
						"",
						"bar",
					),
				),
			).
			Types()
		want := value.NewStrings(
			issue.Info,
			"foo",
			"",
			"bar",
		)
		assert.Equal(t, want, got)
	})
}

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
				impl.WithTypes(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				impl.WithLength(
					value.NewRange(10, 72),
				),
				impl.WithTicket(
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
			).
			Length()
		want := value.NewRange(10, 72)
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

	t.Run("only with types does not set length", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Critical,
						"feat",
						"fix",
					),
				),
			).
			Length()
		want := value.Range{}
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

func TestConfig_Ticket(t *testing.T) {
	t.Parallel()

	t.Run("without any option", func(t *testing.T) {
		t.Parallel()
		got := impl.NewConfig().Ticket()
		var want ticket.Config
		assert.Equal(t, want, got)
	})

	t.Run("with all options", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Warning,
						"feat",
						"fix",
					),
				),
				impl.WithLength(
					value.NewRange(10, 72),
				),
				impl.WithTicket(
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
			).
			Ticket()
		want := ticket.NewConfig(
			ticket.WithID(
				id.NewConfig(
					value.NewString(
						issue.Warning,
						`^[A-Z]+-\d+$`,
					),
				),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("only with ticket", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTicket(
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
			).
			Ticket()
		want := ticket.NewConfig(
			ticket.WithID(
				id.NewConfig(
					value.NewString(
						issue.Critical,
						`^[A-Z]+-\d+$`,
					),
				),
			),
		)
		assert.Equal(t, want, got)
	})

	t.Run("only with types does not set ticket", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTypes(
					value.NewStrings(
						issue.Critical,
						"feat",
						"fix",
					),
				),
			).
			Ticket()
		var want ticket.Config
		assert.Equal(t, want, got)
	})

	t.Run("only with length does not set ticket", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithLength(
					value.NewRange(10, 72),
				),
			).
			Ticket()
		var want ticket.Config
		assert.Equal(t, want, got)
	})

	t.Run("with empty ticket config", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTicket(
					ticket.NewConfig(),
				),
			).
			Ticket()
		want := ticket.NewConfig()
		assert.Equal(t, want, got)
	})

	t.Run("last WithTicket wins", func(t *testing.T) {
		t.Parallel()
		got := impl.
			NewConfig(
				impl.WithTicket(
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
				impl.WithTicket(
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
			).
			Ticket()
		want := ticket.NewConfig(
			ticket.WithID(
				id.NewConfig(
					value.NewString(
						issue.Warning,
						`^[A-Z]+-\d+$`,
					),
				),
			),
		)
		assert.Equal(t, want, got)
	})
}
