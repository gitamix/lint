package git

import (
	"context"

	"github.com/gitamix/types/branch"
)

// Option defines a functional option
// to configure fake Repository instance on its creation.
type Option func(*Repository)

// WithCurrentBranch sets the function that will be executed
// when CurrentBranch() of the fake repo is called.
func WithCurrentBranch(
	fn func(
		ctx context.Context,
	) (branch.Branch, error),
) Option {
	return func(r *Repository) {
		r.funcs.CurrentBranch = fn
	}
}
