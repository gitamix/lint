package shared

import (
	"path/filepath"
	"runtime"
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

// mustRepoRoot returns the root directory of the repository containing this test code
// and panics if it cannot determine the repository root.
//
// This function is useful for locating files relative to the repository root in tests.
func mustRepoRoot() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("resolve test file path")
	}
	return filepath.Clean(
		filepath.Join(
			filepath.Dir(file),
			"../../../../..",
		),
	)
}
