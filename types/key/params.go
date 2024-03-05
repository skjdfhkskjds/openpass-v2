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

package key

import (
	"fmt"

	"github.com/skjdfhkskjds/openpass/v2/types/salt"
)

// Params is a struct that holds the parameters for the
// key derivation function
type Params struct {
	Algorithm string

	Salt    salt.Salt
	KeySize uint32
}

type ParamsOption func(*Params)

func NewParams(opts ...ParamsOption) *Params {
	params := DefaultParams()
	for _, opt := range opts {
		opt(params)
	}
	return params
}

// TODO: fill this with sensible default values
// maybe read from an app.toml or yaml or json file
func DefaultParams() *Params {
	return &Params{}
}

// TODO: i hate fmt
func (p *Params) String() string {
	return fmt.Sprintf("Algorithm: %s, Salt: %s", p.Algorithm, p.Salt)
}

func WithAlgorithm(algorithm string) ParamsOption {
	return func(p *Params) {
		p.Algorithm = algorithm
	}
}

func WithSalt(s salt.Salt) ParamsOption {
	return func(p *Params) {
		p.Salt = s
	}
}

func WithKeySize(keySize uint32) ParamsOption {
	return func(p *Params) {
		p.KeySize = keySize
	}
}
