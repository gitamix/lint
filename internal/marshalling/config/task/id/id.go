package id

import (
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

// ID is the transport representation
// of the task (issue) identifier config.
type ID struct {
	// Pattern stores the transport representation
	// of the task identifier pattern.
	Pattern value.Pattern `yaml:"pattern"`
}

// Config converts the transport representation
// into domain id config, wiring the identifier pattern into it.
//
// Sets critical issue level if it is unspecified
// and the pattern is present.
func (i ID) Config() id.Config {
	v := i.Pattern.Config()
	if v.Level().Unspecified() && !i.Pattern.Empty() {
		v = v.WithLevel(issue.Critical)
	}
	return id.NewConfig(v)
}
