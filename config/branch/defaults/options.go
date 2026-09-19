package defaults

// Option configures a Config instance.
//
// Option is a functional option accepted by NewConfig
// to customize the created configuration.
type Option func(*Config)

// WithName sets the name of the default branch.
func WithName(name string) Option {
	return func(c *Config) {
		c.name = name
	}
}
