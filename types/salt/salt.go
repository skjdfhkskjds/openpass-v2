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

package salt

import "crypto/rand"

// A Salt is a random value used to increase the security of
// a hash.
// It helps prevent precomputed rainbow table attacks.
// Since it is unique for each password, we must store the
// generated salt as a parameter of the password in the database.
type Salt []byte

// New creates a new salt with size Length
func New() (Salt, error) {
	salt := make(Salt, Length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}
