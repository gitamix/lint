package value

import (
	"github.com/gitamix/lint/config/value"
)

// Range is the transport representation of a range config value,
// an inclusive [min, max] interval paired with the issue type level
// used to report violations of the interval.
type Range struct {
	// Issue stores the transport representation
	// of the issue type level associated with the range value.
	Issue Issue `yaml:"issue,omitempty"`

	// Min stores the lower bound of the interval
	// in its transport representation.
	Min int `yaml:"min,omitempty"`

	// Max stores the upper bound of the interval
	// in its transport representation.
	Max int `yaml:"max,omitempty"`
}

// Empty reports whether the Range is empty,
// i.e. the issue level and both bounds are unset.
func (r Range) Empty() bool {
	return r.Issue.Empty() && r.Min == 0 && r.Max == 0
}

// Config converts the Range into the domain value.
//
// When empty, the zero range value is returned.
// Otherwise a new range value is built
// from the issue level and the bounds.
func (r Range) Config() value.Range {
	if r.Empty() {
		return value.Range{}
	}
	return value.NewRange(
		r.Issue.Config(),
		r.Min,
		r.Max,
	)
}
