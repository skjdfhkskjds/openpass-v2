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

package keychain

import (
	"github.com/skjdfhkskjds/openpass/v2/crypto"
	"github.com/skjdfhkskjds/openpass/v2/types"
	"github.com/skjdfhkskjds/openpass/v2/types/password"
)

// this file is responsible for handling all
// keychain related operations

type Keychain struct {
	*types.User

	keyFunc   crypto.KeyDerivationFunction
	algorithm crypto.Algorithm
}

func New(
	user *types.User,
	kdf crypto.KeyDerivationFunction,
	a crypto.Algorithm,
) *Keychain {
	return &Keychain{
		User:      user,
		keyFunc:   kdf,
		algorithm: a,
	}
}

func (k *Keychain) SetPassword(url, username, plainText string) (*password.Password, error) {
	key, err := k.keyFunc.DeriveKey(k.Password)
	if err != nil {
		return nil, err
	}

	k.algorithm.SetKey(key)
	encrypted, err := k.algorithm.Encrypt(url, username, plainText)
	if err != nil {
		return nil, err
	}

	return encrypted, nil
}
