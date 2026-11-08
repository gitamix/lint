package config

import "fmt"

// Validate checks the whole lint configuration for correctness.
//
// It validates the branch and the commit configurations
// returns the error of the first invalid one,
// wrapped with the prefix describing the failed section.
func (c Config) Validate() error {
	err := c.Branch().Validate()
	if err != nil {
		return fmt.Errorf("invalid branch: %w", err)
	}
	err = c.Commit().Validate()
	if err != nil {
		return fmt.Errorf("invalid commit: %w", err)
	}
	return nil
}
