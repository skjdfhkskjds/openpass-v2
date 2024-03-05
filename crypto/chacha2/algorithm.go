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

package chacha2

import (
	"golang.org/x/crypto/chacha20poly1305"

	"github.com/skjdfhkskjds/openpass/v2/crypto"
	"github.com/skjdfhkskjds/openpass/v2/types/key"
	"github.com/skjdfhkskjds/openpass/v2/types/password"
)

// Compile time interface assertion.
var _ crypto.Algorithm = (*Algorithm)(nil)

// Algorithm is an implementation of the crypto.Algorithm interface
// using the ChaCha20-Poly1305 AEAD cipher.
type Algorithm struct {
	key key.Key

	params *password.Params
}

func New(k key.Key) *Algorithm {
	return &Algorithm{
		key:    k,
		params: DefaultChaCha2Params(),
	}
}

func (a *Algorithm) Encrypt(plainText string) (*password.Password, error) {
	aead, err := chacha20poly1305.NewX(a.key.Hash[:])
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize())
	encryptedHash := aead.Seal(nonce, nonce, []byte(plainText), nil)

	// TODO: fill password parameters
	return password.New(encryptedHash, nonce,
		password.Params{Key: a.key},
	), nil
}

func (a *Algorithm) Decrypt(pswd *password.Password) (string, error) {
	aead, err := chacha20poly1305.NewX(a.key.Hash[:])
	if err != nil {
		return "", err
	}

	decryptedPassword, err := aead.Open(nil, pswd.Nonce, pswd.Hash, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedPassword), nil
}
