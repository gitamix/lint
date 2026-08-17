package mandate

import (
	"github.com/gitamix/lint/config/commit/message/body/mandate"
	"github.com/gitamix/lint/internal/marshalling/config/value"
)

// Mandate is the transport representation
// of the body mandate config, describing
// the commit types for which the commit message body is mandatory.
type Mandate struct {
	// Types stores the transport representation
	// of the commit types for which the commit message body is mandatory.
	Types value.Strings `yaml:"types,omitempty"`
}

// Config converts the transport representation
// into the domain mandate config, wiring the commit types into it.
func (c Mandate) Config() mandate.Config {
	return mandate.NewConfig(
		mandate.WithTypes(
			c.Types.Config(),
		),
	)
}
