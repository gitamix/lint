package types

import (
	"github.com/gitamix/lint/config/commit/types"
	"github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

// Types is the transport representation of the accepted commit types config.
type Types struct {
	// Types stores the transport representation of the accepted commit types.
	Types value.Strings `yaml:"types,omitempty"`
}

// Config converts the transport representation into the domain types.Config,
// wiring the accepted commit types into it.
func (c Types) Config() types.Config {
	v := c.Types.Config()
	if !c.Types.Empty() && v.Level().Unspecified() {
		v = v.WithLevel(issue.Critical)
	}
	return types.NewConfig(v)
}
