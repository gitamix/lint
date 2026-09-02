package message

import (
	"github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/internal/marshalling/config/commit/message/body"
	"github.com/gitamix/lint/internal/marshalling/config/commit/message/subject"
)

// Message is the transport representation
// of the commit message config.
//
// Message groups the subject and body of a commit message and converts
// them into domain message config consumed by the linter.
type Message struct {
	// Subj stores the transport representation
	// of the commit message subject config.
	Subj subject.Subject `yaml:"subject,omitempty"`

	// Body stores the transport representation
	// of the commit message body config.
	Body body.Body `yaml:"body,omitempty"`
}

// Config converts the transport representation
// into domain message config,
// wiring the subject and body representations into it.
func (c Message) Config() message.Config {
	return message.NewConfig(
		message.WithSubject(
			c.Subj.Config(),
		),
		message.WithBody(
			c.Body.Config(),
		),
	)
}
