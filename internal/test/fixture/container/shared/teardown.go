package shared

import (
	"context"
	"time"
)

// Teardown stops and removes the shared container fixture
// if it has been initialized for the current test process.
func Teardown(ctx context.Context, timeout time.Duration) error {
	if sharedFX == nil || sharedFX.Container() == nil {
		return nil
	}
	termCtx, termCancel := context.WithTimeout(ctx, timeout)
	defer termCancel()
	return sharedFX.Terminate(termCtx, timeout)
}
