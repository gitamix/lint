package value

import "github.com/gitamix/lint/issue"

// String represents config string value
// that can be configured with issue type level to alert on linting.
type String struct {
	// v is exact value.
	v string

	// lvl is an issue type level to alert on linting.
	lvl issue.Type
}

// NewString creates a new config string value
// with issue type level and exact value.
func NewString(lvl issue.Type, v string) String {
	return String{
		v:   v,
		lvl: lvl,
	}
}

// Exact returns exact value.
func (s String) Exact() string {
	return s.v
}

// Level returns issue type level to alert on linting.
//
// Returns Warning if not set.
func (s String) Level() issue.Type {
	if s.lvl == 0 {
		return issue.Warning
	}
	return s.lvl
}

// Equal defines whether the exact value equals the provided string.
func (s String) Equal(other string) bool {
	return s.Exact() == other
}
