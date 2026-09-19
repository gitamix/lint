package defaults

// Config represents configuration of the default branch.
//
// Config keeps the name of the branch
// that the linter treats as the main branch of the repository.
type Config struct {
	// name stores the name of the default branch.
	//
	// The value names the main branch of the repository,
	// for example "master", and is used by the linter
	// whenever the well-known branch is needed.
	name string
}

// NewConfig creates a new Config instance
// with provided functional options to customize the configuration.
//
// Without options the created Config has an empty default branch name.
func NewConfig(opts ...Option) Config {
	c := Config{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// Name returns the name of the default branch.
func (c Config) Name() string {
	return c.name
}
