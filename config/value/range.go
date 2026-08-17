package value

import (
	"fmt"

	"github.com/gitamix/lint/issue"
)

// Range represents an inclusive integer interval [min, max]
// that can be configured with an issue type level to alert on linting.
//
// Range is immutable and validates its invariant
// on every method call:
// min must not be greater than max.
type Range struct {
	// min stores the lower bound of the interval.
	min int

	// max stores the upper bound of the interval.
	max int

	// lvl is an issue type level to alert on linting.
	lvl issue.Type
}

// NewRange creates a Range with the provided issue type level
// and inclusive bounds.
func NewRange(lvl issue.Type, min, max int) Range {
	return Range{
		min: min,
		max: max,
		lvl: lvl,
	}
}

// Level returns issue type level to alert on linting.
//
// Returns Warning if not set.
func (r Range) Level() issue.Type {
	return r.lvl
}

// String returns the canonical "min-max"
// representation of the interval.
//
// Panics if min is greater than max.
func (r Range) String() string {
	r.mustValidate()
	return fmt.Sprintf("%d-%d", r.min, r.max)
}

// Fit reports whether n falls within
// the inclusive interval [min, max].
//
// Panics if min is greater than max.
func (r Range) Fit(n int) bool {
	return !r.Below(n) && !r.Above(n)
}

// Below reports whether n is strictly less
// than the lower bound of the interval.
//
// Panics if min is greater than max.
func (r Range) Below(n int) bool {
	r.mustValidate()
	return n < r.min
}

// Above reports whether n is strictly greater
// than the upper bound of the interval.
//
// Panics if min is greater than max.
func (r Range) Above(n int) bool {
	r.mustValidate()
	return n > r.max
}

// mustValidate panics if the Range invariant is broken,
// i.e. min is greater than max.
func (r Range) mustValidate() {
	if r.min > r.max {
		panic(
			fmt.Sprintf(
				"min value is greater than max: %d > %d",
				r.min,
				r.max,
			),
		)
	}
}
