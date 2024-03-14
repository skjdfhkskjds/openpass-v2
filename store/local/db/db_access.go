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
	"errors"

	"github.com/dgraph-io/badger/v4"

	"github.com/skjdfhkskjds/openpass/v2/store"
)

// read performs a read operation on the db.
func (db *LocalDB) Read(key []byte) ([]byte, error) {
	var data []byte
	if err := db.badgerDB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if errors.Is(err, badger.ErrKeyNotFound) {
			return store.ErrKeyNotFound
		} else if err != nil {
			return err
		}

		// Get the password
		return item.Value(func(val []byte) error {
			data = val
			return nil
		})
	}); err != nil {
		return nil, err
	}
	return data, nil
}

// write performs a write operation on the db.
// It returns an error if the key already exists in the db.
func (db *LocalDB) Write(key []byte, value []byte) error {
	return db.badgerDB.Update(func(txn *badger.Txn) error {
		_, err := txn.Get(key)
		// If key is already in the db, return an error
		if err == nil {
			return store.ErrKeyAlreadyExists
		} else if !errors.Is(err, badger.ErrKeyNotFound) {
			return err
		}

		err = txn.Set(key, value)
		if err != nil {
			return err
		}
		return nil
	})
}

// delete performs a delete operation on the db.
// If the key does not exist in the db, it is a no-op.
func (db *LocalDB) Delete(key []byte) error {
	return db.badgerDB.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

