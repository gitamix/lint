package description

import (
	"github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/internal/marshalling/config/value"
	"github.com/gitamix/lint/issue"
)

// Description is the transport representation
// of the commit message subject description config.
type Description struct {
	// Length stores the transport representation
	// of the allowed length interval of the subject description text.
	Length value.Range `yaml:"length,omitempty"`
}

// Config converts the transport representation into the domain
// description.Config, wiring the length interval into it.
func (c Description) Config() description.Config {
	v := c.Length.Config()
	if !c.Length.Empty() && v.Level().Unspecified() {
		v = v.WithLevel(issue.Warning)
	}
	return description.NewConfig(
		description.WithLength(v),
	)
}
