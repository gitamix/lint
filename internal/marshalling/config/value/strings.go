package value

import (
	"github.com/gitamix/lint/config/value"
)

// Strings is the transport representation of a string list config value,
// pairing a list of exact string values with the issue type level
// used to report violations of the list.
type Strings struct {
	// List stores the slice of exact string values
	// in its transport representation.
	List []string `yaml:"list,omitempty"`

	// Issue stores the transport representation
	// of the issue type level associated with the list of values.
	Issue Issue `yaml:"issue,omitempty"`
}

// Empty reports whether the Strings is empty,
// i.e. neither the issue level nor the list of values is set.
func (s Strings) Empty() bool {
	return s.Issue.Empty() && len(s.List) == 0
}

// Config converts the Strings into the domain value.
//
// When empty, the zero strings value is returned.
// Otherwise a new strings value is built
// from the issue level and the list of values.
func (s Strings) Config() value.Strings {
	if s.Empty() {
		return value.Strings{}
	}
	return value.NewStrings(
		s.Issue.Config(),
		s.List...,
	)
}
