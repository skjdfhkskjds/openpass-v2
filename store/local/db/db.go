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

package localdb

import (
	badger "github.com/dgraph-io/badger/v4"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//go:generate mockery --name DB
type DB interface {
	Read(key []byte) ([]byte, error)
	Write(key []byte, message protoreflect.ProtoMessage) error
	Delete(key []byte) error
	Update(key []byte, message protoreflect.ProtoMessage) error
	Close() error
}

var _ DB = (*LocalDB)(nil)

type LocalDB struct {
	badgerDB *badger.DB
}

func New(path string) (*LocalDB, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return &LocalDB{db}, nil
}

func (db *LocalDB) Close() error {
	return db.badgerDB.Close()
}
