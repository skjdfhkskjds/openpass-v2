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

package user

import (
	"bytes"
	"encoding/gob"
)

// NewFromBytes returns a new user.Data struct from bytes
func NewFromBytes(b []byte) (*Data, error) {
	buffer := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buffer)
	var d Data
	if err := decoder.Decode(&d); err != nil {
		return nil, err
	}

	return &d, nil
}

// Bytes returns the bytes representation of a user.Data struct
func (d *Data) Bytes() ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(d); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (d *Data) BytesUnsafe() []byte {
	bz, err := d.Bytes()
	if err != nil {
		panic(err)
	}

	return bz
}
