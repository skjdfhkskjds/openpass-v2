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

	"github.com/stretchr/testify/require"

	localdb "github.com/skjdfhkskjds/openpass/v2/store/local/db"
	"github.com/skjdfhkskjds/openpass/v2/types/password"
)

var (
	testPassword = &password.Password{
		URL:      "URL",
		Username: "username",
	}

	testPassword2 = &password.Password{
		URL:      "URL2",
		Username: "username2",
	}

	invalidPassword = &password.Password{
		URL:      "",
		Username: "",
	}

	testPasswordBz    = testPassword.BytesUnsafe()
	testPassword2Bz   = testPassword2.BytesUnsafe()
	invalidPasswordBz = invalidPassword.BytesUnsafe()
)

var (
	validPath = localdb.BuildPasswordKeyPath(
		testPassword.URL,
		testPassword.Username,
	)
	validPath2 = localdb.BuildPasswordKeyPath(
		testPassword2.URL,
		testPassword2.Username,
	)
	invalidPath = localdb.BuildPasswordKeyPath(
		invalidPassword.URL,
		invalidPassword.Username,
	)
)

func TestGetPassword(t *testing.T) {
	t.Run("Valid call", func(t *testing.T) {
		db.On("Read", validPath).Return(testPasswordBz, nil)

		pswd, err := store.GetPassword(testPassword.URL, testPassword.Username)
		require.NoError(t, err)
		require.Equal(t, testPassword, pswd)
	})

	t.Run("Invalid call", func(t *testing.T) {
		db.On("Read", invalidPath).Return(nil, errForcedError)

		_, err := store.GetPassword(invalidPassword.URL, invalidPassword.Username)
		require.ErrorIs(t, err, errForcedError)
	})
}

func TestSetPassword(t *testing.T) {
	t.Run("Valid call", func(t *testing.T) {
		db.On("Write", validPath, testPasswordBz).Return(nil)
		require.NoError(t, store.SetPassword(testPassword))
	})

	t.Run("Invalid call", func(t *testing.T) {
		db.On("Write", invalidPath, invalidPasswordBz).Return(errForcedError)
		require.ErrorIs(t, store.SetPassword(invalidPassword), errForcedError)
	})
}

func TestUpdatePassword(t *testing.T) {
	t.Run("Valid call", func(t *testing.T) {
		db.On("Delete", validPath).Return(nil)
		db.On("Write", validPath2, testPassword2Bz).Return(nil)

		require.NoError(t,
			store.UpdatePassword(testPassword.URL, testPassword.Username, testPassword2),
		)
	})

	t.Run("Invalid call", func(t *testing.T) {
		db.On("Delete", validPath).Return(nil)
		db.On("Delete", invalidPath).Return(errForcedError)

		db.On("Write", validPath2, testPassword2Bz).Return(nil)
		db.On("Write", invalidPath, invalidPasswordBz).Return(errForcedError)

		// Test with invalid old password
		err := store.UpdatePassword(invalidPassword.URL, invalidPassword.Username, testPassword2)
		require.ErrorIs(t, err, errForcedError)

		// Test with invalid updated password
		err = store.UpdatePassword(testPassword.URL, testPassword.Username, invalidPassword)
		require.ErrorIs(t, err, errForcedError)
	})
}

func TestDeletePassword(t *testing.T) {
	t.Run("Valid call", func(t *testing.T) {
		db.On("Delete", validPath).Return(nil)
		require.NoError(t, store.DeletePassword(testPassword.URL, testPassword.Username))
	})

	t.Run("Invalid call", func(t *testing.T) {
		db.On("Delete", invalidPath).Return(errForcedError)
		require.ErrorIs(t,
			store.DeletePassword(invalidPassword.URL, invalidPassword.Username),
			errForcedError,
		)
	})
}
