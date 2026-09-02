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
func (s String) Level() issue.Type {
	return s.lvl
}

// Equal defines whether the exact value equals the provided string.
func (s String) Equal(other string) bool {
	return s.Exact() == other
}

// Empty defines whether the value is empty.
func (s String) Empty() bool {
	return s.v == ""
}

// String returns the exact value.
func (s String) String() string {
	return s.v
}

// WithLevel creates a new instance with provided issue type level.
func (s String) WithLevel(lvl issue.Type) String {
	return NewString(lvl, s.v)
}
