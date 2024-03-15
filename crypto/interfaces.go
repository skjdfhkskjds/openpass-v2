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

package crypto

import (
	"github.com/skjdfhkskjds/openpass/v2/types/key"
	"github.com/skjdfhkskjds/openpass/v2/types/password"
)

// Algorithm is an interface for encryption and decryption
// of a password.
type Algorithm interface {
	SetKey(key *key.Key)

	Encrypt(URL, username, plainText string) (*password.Password, error)
	Decrypt(password *password.Password) (string, error)
}

// KeyDerivationFunction is an interface for deriving a key
// from a master password.
// This key is used as a standard length key for encryption
// and decryption by the Algorithm interface.
type KeyDerivationFunction interface {
	DeriveKey(password string) (*key.Key, error)
}
