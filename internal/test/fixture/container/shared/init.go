package shared

import (
	"context"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"runtime"
	"time"

	tc "github.com/testcontainers/testcontainers-go"

	"github.com/gitamix/lint/internal/test/container"
	"github.com/gitamix/lint/internal/test/container/env"
	ctrfx "github.com/gitamix/lint/internal/test/fixture/container"
)

// init initializes the shared container fixture.
//
// Creates and sets a container fixture backed
// by the shared test container
// and loads the fixture environment variables from it.
func init() {
	fixtureOnce.Do(func() {
		ctx, cancel := context.WithTimeout(
			context.Background(),
			2*time.Minute,
		)
		defer cancel()
		sharedFX, sharedErr = setupSharedContainerFixture(ctx)
	})
	if sharedErr != nil {
		log.Fatalf("shared container fixture: %v", sharedErr)
	}
}

// setupSharedContainerFixture creates a container fixture backed
// by the shared test container
// and loads the fixture environment variables from it.
func setupSharedContainerFixture(ctx context.Context) (*ctrfx.Container, error) {
	ctr, err := container.TestContainer(
		ctx,
		2*time.Minute,
		tc.FromDockerfile{
			Context:        mustRepoRoot(),
			Dockerfile:     container.PathToDockerfile,
			Repo:           container.ImageRepo,
			Tag:            container.ImageTag,
			KeepImage:      true,
			BuildLogWriter: io.Discard,
		},
	)
	if err != nil {
		return nil, err
	}
	vars, err := env.Load(
		ctx,
		ctr,
		container.PathToEnv,
	)
	if err != nil {
		_ = ctr.Terminate(ctx)
		return nil, fmt.Errorf("failed to load env: %w", err)
	}
	return ctrfx.NewContainer(ctr, vars), nil
}

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
