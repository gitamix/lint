package shared

import (
	"sync"

	ctrfx "github.com/gitamix/lint/internal/test/fixture/container"
)

var (
	// fixtureOnce guarantees the shared integration fixture
	// is initialized only once per test process.
	fixtureOnce sync.Once

	// sharedFX stores the package-level container fixture instance
	// reused by integration tests across all packages.
	sharedFX *ctrfx.Container

	// sharedErr stores the initialization error returned
	// during the first shared fixture setup attempt.
	sharedErr error
)

// ContainerFixture returns the global shared container fixture
// used by integration tests across all packages.
func ContainerFixture() *ctrfx.Container {
	return sharedFX
}
