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

package localstore

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"

	localdb "github.com/skjdfhkskjds/openpass/v2/store/local/db"
	prototypes "github.com/skjdfhkskjds/openpass/v2/types/proto/types/v1"
)

var (
	testUserData = &prototypes.UserData{
		Username: "test",
	}

	invalidUserData = &prototypes.UserData{
		Username: "",
	}

	testUserDataBz, _ = proto.Marshal(testUserData)
)

var (
	validUserKeyPath   = localdb.BuildUserDataKeyPath(testUserData.GetUsername())
	invalidUserKeyPath = localdb.BuildUserDataKeyPath(invalidUserData.GetUsername())
)

func TestGetUserData(t *testing.T) {
	t.Run("Valid User", func(t *testing.T) {
		db.On("Read", mock.MatchedBy(func(key []byte) bool {
			return bytes.Equal(key, validUserKeyPath)
		})).Return(testUserDataBz, nil)

		res, err := store.GetUserData(&prototypes.GetUserDataRequest{
			Username: testUserData.GetUsername(),
		})
		if ok := assert.NoError(t, err); !ok {
			t.FailNow()
		}

		if ok := assert.True(t, proto.Equal(testUserData, res.GetUserData())); !ok {
			t.Fatalf("expected: %v, got: %v", testUserData, res.GetUserData())
		}
	})

	t.Run("Invalid User", func(t *testing.T) {
		db.On("Read", mock.MatchedBy(func(key []byte) bool {
			return bytes.Equal(key, invalidUserKeyPath)
		})).Return(nil, errForcedError)

		_, err := store.GetUserData(&prototypes.GetUserDataRequest{
			Username: invalidUserData.GetUsername(),
		})
		if ok := assert.ErrorIs(t, err, errForcedError); !ok {
			t.Fatalf("expected: %v, got: %v", errForcedError, err)
		}
	})
}

func TestSetUserData(t *testing.T) {
	t.Run("Valid User", func(t *testing.T) {
		db.On("Write", mock.MatchedBy(func(key []byte) bool {
			return bytes.Equal(key, validUserKeyPath)
		}), testUserData).Return(nil)

		res, err := store.SetUserData(&prototypes.SetUserDataRequest{
			UserData: testUserData,
		})
		if ok := assert.NoError(t, err); !ok {
			t.FailNow()
		}

		if ok := assert.True(t, proto.Equal(testUserData, res.GetUserData())); !ok {
			t.Fatalf("expected: %v, got: %v", testUserData, res.GetUserData())
		}
	})

	t.Run("Invalid User", func(t *testing.T) {
		db.On("Write", mock.MatchedBy(func(key []byte) bool {
			return bytes.Equal(key, invalidUserKeyPath)
		}), invalidUserData).Return(errForcedError)

		_, err := store.SetUserData(&prototypes.SetUserDataRequest{
			UserData: invalidUserData,
		})
		if ok := assert.ErrorIs(t, err, errForcedError); !ok {
			t.Fatalf("expected: %v, got: %v", errForcedError, err)
		}
	})
}
