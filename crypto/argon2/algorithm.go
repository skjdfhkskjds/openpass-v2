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

package argon2

import (
	"golang.org/x/crypto/argon2"

	"github.com/skjdfhkskjds/openpass/v2/crypto"
	"github.com/skjdfhkskjds/openpass/v2/types/key"
	keyparams "github.com/skjdfhkskjds/openpass/v2/types/key/params"
)

// Compile time interface assertion.
var _ crypto.KeyDerivationFunction = (*Algorithm)(nil)

type Algorithm struct {
	*keyparams.Params
}

func New(params *keyparams.Params) *Algorithm {
	return &Algorithm{
		Params: params,
	}
}

// TODO: figure out good params system i dont really like how coupled
// it feels
func (a *Algorithm) DeriveKey(password string) (*key.Key, error) {
	idKey := argon2.IDKey(
		[]byte(password),
		a.Salt,
		defaultTime,
		defaultMemorySize,
		defaultThreads,
		a.KeySize,
	)

	var hashedKey [key.Size]byte
	copy(hashedKey[:], idKey)
	return key.New(hashedKey, a.Params), nil
}
