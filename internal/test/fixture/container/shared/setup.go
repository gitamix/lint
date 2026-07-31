package shared

import (
	"context"
	"fmt"
	"io"
	"testing"
	"time"

	tc "github.com/testcontainers/testcontainers-go"

	"github.com/gitamix/lint/internal/test/container"
	"github.com/gitamix/lint/internal/test/container/env"
	ctrfx "github.com/gitamix/lint/internal/test/fixture/container"
)

// ContainerFixture returns the global shared container fixture
// used by integration tests across all packages.
//
// It initializes the fixture only once and fails the test
// if the shared setup could not be completed successfully.
func ContainerFixture(t *testing.T) *ctrfx.Container {
	t.Helper()
	fixtureOnce.Do(func() {
		ctx, cancel := context.WithTimeout(
			context.Background(),
			2*time.Minute,
		)
		defer cancel()
		sharedFX, sharedErr = setupSharedContainerFixture(ctx)
	})
	if sharedErr != nil {
		t.Fatalf("failed to setup shared container fixture: %v", sharedErr)
	}
	return sharedFX
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
