// Package shared provides a global shared container fixture
// used across all integration test packages.
//
// The fixture is initialized once per test process via sync.Once,
// ensuring a single test container is shared across packages.
package shared
