//go:build integration
// +build integration

package current_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/gitamix/lint/internal/test/fixture/container/shared"
)

// TestMain is the entry point for the integration tests in this package.
//
// It runs the tests and ensures that the shared container fixture
// is terminated after all tests have completed.
func TestMain(m *testing.M) {
	exitCode := m.Run()
	shared.Teardown(
		context.Background(),
		5*time.Second,
	)
	os.Exit(exitCode)
}
