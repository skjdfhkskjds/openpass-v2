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

package passwordparams

import (
	"fmt"

	keyparams "github.com/skjdfhkskjds/openpass/v2/types/key/params"
)

type Params struct {
	Algorithm string
	Nonce     []byte

	KeyParams *keyparams.Params
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
	return &Params{
		Algorithm: "AES-GCM",
	}
}

func (p *Params) String() string {
	return fmt.Sprintf("Algorithm: %s", p.Algorithm)
}

func WithAlgorithm(algorithm string) ParamsOption {
	return func(p *Params) {
		p.Algorithm = algorithm
	}
}

func WithNonce(nonce []byte) ParamsOption {
	return func(p *Params) {
		p.Nonce = nonce
	}
}

func WithKeyParams(kp *keyparams.Params) ParamsOption {
	return func(p *Params) {
		p.KeyParams = kp
	}
}
