package value

import (
	"reflect"
	"slices"
	"strings"

	"github.com/gitamix/lint/issue"
)

// Strings represents a config value that is a slice of strings
// that can be configured with an issue type level to alert on linting.
type Strings struct {
	// vv stores the slice of exact values.
	vv []string

	// lvl is an issue type level to alert on linting.
	lvl issue.Type
}

// NewStrings creates a new config value from a slice of string values
// with the given issue type level.
func NewStrings(lvl issue.Type, vv ...string) Strings {
	return Strings{
		lvl: lvl,
		vv:  vv,
	}
}

// String returns a comma-separated string
// representation of the slice of exact values.
func (s Strings) String() string {
	return strings.Join(s.vv, ",")
}

// Exact returns a copy of the slice of exact values.
func (s Strings) Exact() []string {
	dst := make([]string, len(s.vv))
	copy(dst, s.vv)
	return dst
}

// WithLevel creates a new instance with provided issue type level.
func (s Strings) WithLevel(lvl issue.Type) Strings {
	return NewStrings(lvl, s.vv...)
}

// Level returns the issue type level to alert on linting.
//
// Returns Warning if not set.
func (s Strings) Level() issue.Type {
	return s.lvl
}

// Equal defines whether the slice of exact values
// equals the provided slice.
func (s Strings) Equal(other []string) bool {
	return reflect.DeepEqual(s.vv, other)
}

// Has defines whether the slice of exact values
// contains the provided strings.
func (s Strings) Has(v string, vv ...string) bool {
	if !slices.Contains(s.vv, v) {
		return false
	}
	for _, v := range vv {
		if !slices.Contains(s.vv, v) {
			return false
		}
	}
	return true
}
