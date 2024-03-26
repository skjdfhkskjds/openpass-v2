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

import (
	"crypto/rand"
	"math/big"
)

type Generator struct {
	*Config
}

// New returns a new password generator with the given configuration
func New(config *Config) *Generator {
	return &Generator{config}
}

// NewDefault returns a new password generator with the default configuration
func NewDefault() *Generator {
	return New(DefaultConfig())
}

// Generate returns a randomly generated password based on the configuration
func (g *Generator) Generate() (string, error) {
	var charset string
	var password []byte

	// get selected options from config
	var options = g.getPassOptions()
	for _, option := range options {
		charset += option
	}

	// Ensure password contains at least one char from each option
	for _, option := range options {
		pos, err := rand.Int(rand.Reader, big.NewInt(int64(len(option))))
		if err != nil {
			return "", err
		}
		password = append(password, option[pos.Int64()])
	}
	// fill up rest of password
	for len(password) < g.length {
		pos, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password = append(password, charset[pos.Int64()])
	}
	password = password[:g.length] // trim in case desired length < # chosen options

	err := g.shufflePass(password)
	if err != nil {
		return "", err
	}
	return string(password), nil
}

// getPassOptions returns the selected charsets based on the configuration
func (g *Generator) getPassOptions() []string {
	var options []string
	options = append(options, lowercaseLetters)
	if g.includeUppercase {
		options = append(options, uppercaseLetters)
	}
	if g.includeSpecial {
		options = append(options, specialChars)
	}
	if g.includeNumbers {
		options = append(options, numbers)
	}
	return options
}

// shufflePass mutates <password> into an unbiased permutation of itself using the Fisher-Yates algorithm
func (g *Generator) shufflePass(password []byte) error {
	for i := range password {
		jBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(password)-i)))
		if err != nil {
			return err
		}
		j := int(jBig.Int64()) + i
		password[i], password[j] = password[j], password[i]
	}
	return nil
}
