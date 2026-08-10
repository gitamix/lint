package value

import (
	"regexp"

	"github.com/gitamix/lint/issue"
)

// Pattern represents config regexp pattern value
// that can be configured with issue type level to alert on linting.
type Pattern struct {
	// exp is the compiled regexp expression.
	exp *regexp.Regexp

	// lvl is an issue type level to alert on linting.
	lvl issue.Type
}

// NewPattern creates a new config regexp pattern value
// with issue type level and compiled regexp expression.
func NewPattern(lvl issue.Type, exp *regexp.Regexp) Pattern {
	return Pattern{
		exp: exp,
		lvl: lvl,
	}
}

// Level returns issue type level to alert on linting.
//
// Returns Warning if not set.
func (p Pattern) Level() issue.Type {
	if p.lvl == 0 {
		return issue.Warning
	}
	return p.lvl
}

// String returns the compiled regexp expression as a string.
func (p Pattern) String() string {
	return p.exp.String()
}

// Match defines whether the compiled regexp expression
// matches the provided string.
func (p Pattern) Match(v string) bool {
	return p.exp.MatchString(v)
}
