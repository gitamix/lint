package value

import "github.com/gitamix/lint/config/value"

// Pattern is the transport representation of a pattern config value,
// pairing an exact string pattern with the issue type level
// used to report violations of the pattern.
type Pattern struct {
	// Issue stores the transport representation
	// of the issue type level associated with the pattern value.
	Issue Issue `yaml:"issue,omitempty"`

	// Pattern stores the exact string pattern
	// in its transport representation.
	Pattern string `yaml:"pattern,omitempty"`
}

// Empty reports whether the Pattern is empty,
// i.e. neither the issue level nor the pattern string is set.
func (p Pattern) Empty() bool {
	return p.Issue.Empty() && p.Pattern == ""
}

// Config converts the Pattern into the domain value.
//
// When empty, the zero string value is returned.
// Otherwise a new string value is built
// from the issue level and the pattern.
func (p Pattern) Config() value.String {
	if p.Empty() {
		return value.String{}
	}
	return value.NewString(
		p.Issue.Config(),
		p.Pattern,
	)
}
