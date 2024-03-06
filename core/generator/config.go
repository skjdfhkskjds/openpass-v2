package generator

// Config is the configuration for the password generator
type Config struct {
	// Length of the password to generate
	length int
	// Whether to include special characters in the password
	includeSpecial bool
	// Whether to include numbers in the password
	includeNumbers bool
	// Whether to include uppercase letters in the password
	includeUppercase bool
}

type ConfigOpts func(*Config)

// Length sets the length of the password to generate
func NewConfig(opts ...ConfigOpts) *Config {
	c := DefaultConfig()
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// DefaultConfig returns the default configuration for the password generator
func DefaultConfig() *Config {
	return NewConfig(
		WithLength(defaultLength),
		WithIncludeSpecial(defaultIncludeSpecial),
		WithIncludeNumbers(defaultIncludeNumbers),
		WithIncludeUppercase(defaultIncludeUppercase),
	)
}

// WithLength sets the length of the password to generate
func WithLength(length int) ConfigOpts {
	return func(c *Config) {
		c.length = length
	}
}

// WithIncludeSpecial sets whether to include special characters in the password
func WithIncludeSpecial(includeSpecial bool) ConfigOpts {
	return func(c *Config) {
		c.includeSpecial = includeSpecial
	}
}

// WithIncludeNumbers sets whether to include numbers in the password
func WithIncludeNumbers(includeNumbers bool) ConfigOpts {
	return func(c *Config) {
		c.includeNumbers = includeNumbers
	}
}

// WithIncludeUppercase sets whether to include uppercase letters in the password
func WithIncludeUppercase(includeUppercase bool) ConfigOpts {
	return func(c *Config) {
		c.includeUppercase = includeUppercase
	}
}
