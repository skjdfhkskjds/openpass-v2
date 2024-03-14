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

package localdb_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/dgraph-io/badger/v4"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/skjdfhkskjds/openpass/v2/store"
	localdb "github.com/skjdfhkskjds/openpass/v2/store/local/db"
	prototypes "github.com/skjdfhkskjds/openpass/v2/types/proto/v1"
)

type test struct {
	name  string
	key   []byte
	value []byte
	error error

	preRunFunc  preRunFunc
	postRunFunc postRunFunc
}

type preRunFunc func(*badger.DB) error
type postRunFunc func(*badger.DB) ([]byte, error)

var (
	testMessage = &prototypes.UserData{
		Username: "testValue",
	}
	testMessage2 = &prototypes.UserData{
		Username: "testValue2",
	}
)

var (
	testKey       = []byte("testKey")
	testValue, _  = proto.Marshal(testMessage)
	testValue2, _ = proto.Marshal(testMessage2)

	doesNotExistKey = []byte("doesNotExist")
)

var (
	readTests = []test{
		{
			name:  "readSuccess",
			key:   testKey,
			value: testValue,
			preRunFunc: func(db *badger.DB) error {
				err := db.Update(func(txn *badger.Txn) error {
					return txn.Set(testKey, testValue)
				})
				return err
			},
		},
		{
			name:  "readFailKeyDoesNotExist",
			key:   doesNotExistKey,
			error: store.ErrKeyNotFound,
		},
		{
			name:  "readFailGetError",
			error: badger.ErrEmptyKey,
		},
	}

	writeTests = []test{
		{
			name:  "writeSuccess",
			key:   testKey,
			value: testValue,
			postRunFunc: func(db *badger.DB) ([]byte, error) {
				var val []byte
				err := db.View(func(txn *badger.Txn) error {
					item, err := txn.Get(testKey)
					if err != nil {
						return err
					}
					return item.Value(func(v []byte) error {
						val = v
						return nil
					})
				})
				return val, err
			},
		},
		{
			name:  "writeFailKeyAlreadyExists",
			key:   testKey,
			value: testValue,
			error: store.ErrKeyAlreadyExists,
			preRunFunc: func(db *badger.DB) error {
				err := db.Update(func(txn *badger.Txn) error {
					return txn.Set(testKey, testValue)
				})
				return err
			},
		},
		{
			name:  "writeFailSetError",
			error: badger.ErrEmptyKey,
		},
	}

	deleteTests = []test{
		{
			name:  "deleteSuccess",
			key:   testKey,
			value: testValue,
			preRunFunc: func(db *badger.DB) error {
				err := db.Update(func(txn *badger.Txn) error {
					return txn.Set(testKey, testValue)
				})
				return err
			},
			postRunFunc: func(db *badger.DB) ([]byte, error) {
				err := db.View(func(txn *badger.Txn) error {
					_, err := txn.Get(testKey)
					if errors.Is(err, badger.ErrKeyNotFound) {
						return nil
					}
					return err
				})
				return nil, err
			},
		},
		{
			name:  "deleteKeyDoesNotExistNoOp",
			key:   doesNotExistKey,
			error: nil,
		},
	}

	updateTests = []test{
		{
			name:  "updateSuccess",
			key:   testKey,
			value: testValue2,
			preRunFunc: func(db *badger.DB) error {
				return db.Update(func(txn *badger.Txn) error {
					return txn.Set(testKey, testValue)
				})
			},
			postRunFunc: func(db *badger.DB) ([]byte, error) {
				var val []byte
				err := db.View(func(txn *badger.Txn) error {
					item, err := txn.Get(testKey)
					if err != nil {
						return err
					}
					return item.Value(func(v []byte) error {
						val = v
						return nil
					})
				})
				return val, err
			},
		},
		{
			name:  "updateFailKeyDoesNotExist",
			key:   doesNotExistKey,
			error: store.ErrKeyNotFound,
		},
	}
)

func createTestDirAndDB(dbSetupFunc preRunFunc) (*localdb.LocalDB, string, error) {
	dir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		return nil, "", err
	}

	// manually open a new copy of the db to run setup functions
	if dbSetupFunc != nil {
		var db *badger.DB
		db, err = badger.Open(badger.DefaultOptions(dir))
		if err != nil {
			return nil, "", err
		}
		err = dbSetupFunc(db)
		if err != nil {
			return nil, "", err
		}
		db.Close()
	}

	localDB, err := localdb.New(dir)
	if err != nil {
		return nil, "", err
	}

	return localDB, dir, nil
}

// verifies if the db contents are as expected after a test executes
func verifyDBContents(dir string, expected []byte, postFunc postRunFunc) error {
	if postFunc == nil {
		return nil
	}
	db, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		return err
	}
	defer db.Close()

	key, err := postFunc(db)
	if err != nil {
		return err
	}
	if !bytes.Equal(expected, key) {
		return fmt.Errorf("expected value %s, got %s", expected, key)
	}
	return nil
}

func TestRead(t *testing.T) {
	for _, tt := range readTests {
		db, dir, err := createTestDirAndDB(tt.preRunFunc)
		if ok := assert.NoError(t, err); !ok {
			t.Fatal(err)
		}
		t.Run(tt.name, func(t *testing.T) {
			data, err := db.Read(tt.key)
			if ok := assert.ErrorIs(t, err, tt.error); !ok {
				t.Fatalf("test %s: expected error %s, got %s", tt.name, tt.error, err)
			}
			if ok := assert.Equal(t, tt.value, data); !ok {
				t.Fatalf(
					"test %s: expected value %s, got %s",
					tt.name, string(tt.value), string(data),
				)
			}
			os.RemoveAll(dir)
		})
	}
}

func TestWrite(t *testing.T) {
	for _, tt := range writeTests {
		t.Run(tt.name, func(t *testing.T) {
			db, dir, err := createTestDirAndDB(tt.preRunFunc)
			if ok := assert.NoError(t, err); !ok {
				t.Fatal(err)
			}
			var message protoreflect.ProtoMessage
			if tt.value != nil {
				message = testMessage
			}
			err = db.Write(tt.key, message)
			if ok := assert.ErrorIs(t, err, tt.error); !ok {
				t.Fatalf("test %s: expected error %s, got %s", tt.name, tt.error, err)
			}

			db.Close()
			// Check if the key-value pair was written
			if ok := assert.NoError(t, verifyDBContents(dir, tt.value, tt.postRunFunc)); !ok {
				t.Fatalf("test %s: expected no error, got %s", tt.name, err)
			}
			os.RemoveAll(dir)
		})
	}
}

func TestDelete(t *testing.T) {
	for _, tt := range deleteTests {
		t.Run(tt.name, func(t *testing.T) {
			db, dir, err := createTestDirAndDB(tt.preRunFunc)
			if ok := assert.NoError(t, err); !ok {
				t.Fatal(err)
			}
			err = db.Delete(tt.key)
			if err != nil && !errors.Is(err, tt.error) {
				t.Fatalf("test %s: expected err %s, got %s", tt.name, tt.error, err)
			}

			db.Close()
			// Check if the key-value pair was deleted
			if ok := assert.NoError(t, verifyDBContents(dir, nil, tt.postRunFunc)); !ok {
				t.Fatalf("test %s: expected no error, got %s", tt.name, err)
			}
			os.RemoveAll(dir)
		})
	}
}

func TestUpdate(t *testing.T) {
	for _, tt := range updateTests {
		t.Run(tt.name, func(t *testing.T) {
			db, dir, err := createTestDirAndDB(tt.preRunFunc)
			if ok := assert.NoError(t, err); !ok {
				t.Fatal(err)
			}

			var message *prototypes.UserData
			var expected []byte
			if tt.value != nil {
				message = testMessage2
				expected = testValue2
			}
			err = db.Update(tt.key, message)
			if tt.error != nil && err != nil && errors.Is(err, tt.error) {
				return
			} else if err != nil {
				t.Fatalf("test %s: expected err %s, got %s", tt.name, tt.error, err)
			}

			db.Close()
			// Check if the key-value pair was deleted
			if ok := assert.NoError(t, verifyDBContents(dir, expected, tt.postRunFunc)); !ok {
				t.Fatalf("test %s: expected no error, got %s", tt.name, err)
			}
			os.RemoveAll(dir)
		})
	}
}
