package scope

import (
	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

// Scope is the transport representation
// of the commit message subject scope config.
type Scope struct {
	// Pattern stores the transport representation of the scope pattern.
	Pattern value.Pattern `yaml:"pattern,omitempty"`
}

// Config converts the transport representation
// into domain scope config, wiring the scope pattern into it.
//
// Sets warning issue level if it is unspecified.
func (c Scope) Config() scope.Config {
	v := c.Pattern.Config()
	if !c.Pattern.Empty() && v.Level().Unspecified() {
		v = v.WithLevel(issue.Warning)
	}
	return scope.NewConfig(v)
}
