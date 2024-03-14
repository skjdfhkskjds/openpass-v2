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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	localdb "github.com/skjdfhkskjds/openpass/v2/store/local/db"
	"github.com/skjdfhkskjds/openpass/v2/types/user"
)

var (
	testUserData = &user.Data{
		Username: "test",
	}

	invalidUserData = &user.Data{
		Username: "",
	}

	testUserDataBz    = testUserData.BytesUnsafe()
	invalidUserDataBz = invalidUserData.BytesUnsafe()
)

var (
	validUserKeyPath   = localdb.BuildUserDataKeyPath(testUserData.Username)
	invalidUserKeyPath = localdb.BuildUserDataKeyPath(invalidUserData.Username)
)

func TestGetUserData(t *testing.T) {
	t.Run("Valid User", func(t *testing.T) {
		db.On("Read", validUserKeyPath).Return(testUserDataBz, nil)
		data, err := store.GetUserData(testUserData.Username)
		if ok := assert.NoError(t, err); !ok {
			t.FailNow()
		}

		require.Equal(t, testUserData, data)
	})

	t.Run("Invalid User", func(t *testing.T) {
		db.On("Read", invalidUserKeyPath).Return(nil, errForcedError)

		_, err := store.GetUserData(invalidUserData.Username)
		require.ErrorIs(t, err, errForcedError)
	})
}

func TestSetUserData(t *testing.T) {
	t.Run("Valid User", func(t *testing.T) {
		db.On("Write", validUserKeyPath, testUserDataBz).Return(nil)
		require.NoError(t, store.SetUserData(testUserData))
	})

	t.Run("Invalid User", func(t *testing.T) {
		db.On("Write", invalidUserKeyPath, invalidUserDataBz).Return(errForcedError)
		require.ErrorIs(t, store.SetUserData(invalidUserData), errForcedError)
	})
}
