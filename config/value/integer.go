package value

import (
	"strconv"

	"github.com/gitamix/lint/issue"
)

// Integer represents config integer value
// that can be configured with issue type level to alert on linting.
type Integer struct {
	// v is the exact integer value.
	v int

	// lvl is an issue type level to alert on linting.
	lvl issue.Type
}

// NewInteger creates a new config integer value
// with issue type level and exact value.
func NewInteger(lvl issue.Type, v int) Integer {
	return Integer{
		v:   v,
		lvl: lvl,
	}
}

// Exact returns the exact integer value.
func (i Integer) Exact() int {
	return i.v
}

// Level returns issue type level to alert on linting.
//
// Returns Warning if not set.
func (i Integer) Level() issue.Type {
	if i.lvl == 0 {
		return issue.Warning
	}
	return i.lvl
}

// Equal defines whether the exact integer value
// equals the provided integer.
func (i Integer) Equal(other int) bool {
	return i.Exact() == other
}

// String returns the exact integer value as a string.
func (i Integer) String() string {
	return strconv.Itoa(i.v)
}
