package branch

import "fmt"

// Validate checks the branch configuration for correctness.
//
// It validates the default branch, the name and the task configurations
// in that order and returns the error of the first invalid one,
// wrapped with the prefix describing the failed section.
func (b Config) Validate() error {
	var err error
	err = b.Default().Validate()
	if err != nil {
		return fmt.Errorf("invalid default branch: %w", err)
	}
	err = b.Name().Validate()
	if err != nil {
		return fmt.Errorf("invalid name: %w", err)
	}
	err = b.Task().Validate()
	if err != nil {
		return fmt.Errorf("invalid name: %w", err)
	}
	return nil
}
