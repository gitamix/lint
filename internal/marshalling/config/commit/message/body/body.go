package body

import (
	"github.com/gitamix/lint/config/commit/message/body"
	"github.com/gitamix/lint/internal/marshalling/config/commit/message/body/mandate"
	"github.com/gitamix/lint/internal/marshalling/config/value"
)

// Body is the transport representation
// of the commit message body config.
//
// Body groups the body mandate
// and the allowed body length and converts
// them into the domain body config consumed by the linter.
type Body struct {
	// Mandate stores the transport representation
	// of the body mandate config,
	// i.e. the commit types for which the body is mandatory.
	Mandate mandate.Mandate `yaml:"mandate,omitempty"`

	// Length stores the transport representation
	// of the allowed body length interval.
	Length value.Range `yaml:"length,omitempty"`
}

// Config converts the transport representation into domain body config.
func (c Body) Config() body.Config {
	return body.NewConfig(
		body.WithMandate(
			c.Mandate.Config(),
		),
		body.WithLength(
			c.Length.Config(),
		),
	)
}
