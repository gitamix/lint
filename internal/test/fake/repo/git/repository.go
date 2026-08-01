package git

import (
	"context"

	"github.com/gitamix/types/branch"
)

// Repository represents fake implemenation
// of git repository for test purposes.
type Repository struct {
	// funcs holds the functions that define
	//  behavior of the fake repo methods.
	funcs funcs
}

// NewRepository creates a new instance of fake Repository
// with the provided options to configure its behavior.
func NewRepository(opts ...Option) *Repository {
	r := &Repository{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// CurrentBranch simulates the behavior of the CurrentBranch method.
//
// Panics if the behavior for the Get method is not specified.
func (r *Repository) CurrentBranch(
	ctx context.Context,
) (branch.Branch, error) {
	if r.funcs.CurrentBranch == nil {
		panic("not specified behavior for CurrentBranch method")
	}
	return r.funcs.CurrentBranch(ctx)
}
