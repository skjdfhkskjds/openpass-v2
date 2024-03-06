// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 skjdfhkskjds
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

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
