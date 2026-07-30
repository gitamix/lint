package git

import (
	"context"

	"github.com/gitamix/types/branch"
)

// funcs holds the functions that will be executed
// when the corresponding methods of the fake repo are called.
type funcs struct {
	// CurrentBranch is the function that will be executed
	// when CurrentBranch() of the fake repo is called.
	CurrentBranch func(
		ctx context.Context,
	) (branch.Branch, error)
}
