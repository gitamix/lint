package defaults

import (
	"github.com/gitamix/lint/config/branch/defaults"
)

// Default is the transport representation
// of the default branch config, describing
// the branch that the linter treats as the main branch of the repository.
type Default struct {
	// Name stores the transport representation
	// of the default branch name, for example "master".
	Name string `yaml:"name,omitempty"`
}

// Config converts the transport representation
// into the domain defaults config, wiring the default branch name into it.
func (d Default) Config() defaults.Config {
	return defaults.NewConfig(
		defaults.WithName(d.Name),
	)
}
