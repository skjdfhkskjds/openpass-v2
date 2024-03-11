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
	"google.golang.org/protobuf/reflect/protoreflect"

	localdb "github.com/skjdfhkskjds/openpass/v2/store/local/db"
	prototypes "github.com/skjdfhkskjds/openpass/v2/types/proto/types/v1"
)

var (
	testPassword = &prototypes.PasswordEntry{
		Url:      "url",
		Username: "username",
	}

	invalidPassword = &prototypes.PasswordEntry{
		Url:      "",
		Username: "",
	}

	testPasswordBz, _ = proto.Marshal(testPassword)
)

var (
	validPath = localdb.BuildPasswordKeyPath(
		testPassword.GetUrl(),
		testPassword.GetUsername(),
	)
	invalidPath = localdb.BuildPasswordKeyPath(
		invalidPassword.GetUrl(),
		invalidPassword.GetUsername(),
	)
)

func TestGetPassword(t *testing.T) {
	t.Run("Valid call", func(t *testing.T) {
		db.On("Read", mock.MatchedBy(func(bz []byte) bool {
			return bytes.Equal(bz, validPath)
		})).Return(testPasswordBz, nil)

		res, err := store.GetPassword(&prototypes.GetPasswordRequest{
			Url:      testPassword.GetUrl(),
			Username: testPassword.GetUsername(),
		})
		if ok := assert.NoError(t, err); !ok {
			t.Fatal(err)
		}

		if ok := assert.True(t, proto.Equal(testPassword, res.GetPassword())); !ok {
			t.Fatalf("expected: %v, got: %v", testPassword, res.GetPassword())
		}
	})

	t.Run("Invalid call", func(t *testing.T) {
		db.On("Read", mock.MatchedBy(func(bz []byte) bool {
			return bytes.Equal(bz, invalidPath)
		})).Return(nil, errForcedError)

		_, err := store.GetPassword(&prototypes.GetPasswordRequest{
			Url:      "",
			Username: "",
		})
		if ok := assert.ErrorIs(t, err, errForcedError); !ok {
			t.Fatalf("expected: %v, got: %v", errForcedError, err)
		}
	})
}

func TestSetPassword(t *testing.T) {
	t.Run("Valid call", func(t *testing.T) {
		db.On("Write", mock.MatchedBy(func(bz []byte) bool {
			return bytes.Equal(bz, validPath)
		}), mock.MatchedBy(func(msg protoreflect.ProtoMessage) bool {
			return proto.Equal(msg, testPassword)
		})).Return(nil)

		res, err := store.SetPassword(&prototypes.SetPasswordRequest{Password: testPassword})
		if ok := assert.NoError(t, err); !ok {
			t.Fatal(err)
		}

		if ok := assert.True(t, proto.Equal(testPassword, res.GetPassword())); !ok {
			t.Fatalf("expected: %v, got: %v", testPassword, res.GetPassword())
		}
	})

	// Invalid Call
	t.Run("Invalid call", func(t *testing.T) {
		db.On("Write", mock.MatchedBy(func(bz []byte) bool {
			return bytes.Equal(bz, invalidPath)
		}), mock.MatchedBy(func(msg protoreflect.ProtoMessage) bool {
			return proto.Equal(msg, invalidPassword)
		})).Return(errForcedError)

		_, err := store.SetPassword(&prototypes.SetPasswordRequest{
			Password: invalidPassword,
		})
		if ok := assert.ErrorIs(t, err, errForcedError); !ok {
			t.Fatalf("expected: %v, got: %v", errForcedError, err)
		}
	})
}

func TestUpdatePassword(t *testing.T) {
	t.Run("Valid call", func(t *testing.T) {
		db.On("Update", mock.MatchedBy(func(bz []byte) bool {
			return bytes.Equal(bz, validPath)
		}), mock.MatchedBy(func(msg protoreflect.ProtoMessage) bool {
			return proto.Equal(msg, testPassword)
		})).Return(nil)

		res, err := store.UpdatePassword(&prototypes.UpdatePasswordRequest{Password: testPassword})
		if ok := assert.NoError(t, err); !ok {
			t.Fatal(err)
		}

		if ok := assert.True(t, proto.Equal(testPassword, res.GetPassword())); !ok {
			t.Fatalf("expected: %v, got: %v", testPassword, res.GetPassword())
		}
	})

	t.Run("Invalid call", func(t *testing.T) {
		db.On("Update", mock.MatchedBy(func(bz []byte) bool {
			return bytes.Equal(bz, invalidPath)
		}), mock.MatchedBy(func(msg protoreflect.ProtoMessage) bool {
			return proto.Equal(msg, invalidPassword)
		})).Return(errForcedError)

		_, err := store.UpdatePassword(&prototypes.UpdatePasswordRequest{
			Password: invalidPassword,
		})
		if ok := assert.ErrorIs(t, err, errForcedError); !ok {
			t.Fatalf("expected: %v, got: %v", errForcedError, err)
		}
	})
}

func TestDeletePassword(t *testing.T) {
	t.Run("Valid call", func(t *testing.T) {
		db.On("Delete", mock.MatchedBy(func(bz []byte) bool {
			return bytes.Equal(bz, validPath)
		})).Return(nil)

		_, err := store.DeletePassword(&prototypes.DeletePasswordRequest{
			Url:      testPassword.GetUrl(),
			Username: testPassword.GetUsername(),
		})
		if ok := assert.NoError(t, err); !ok {
			t.Fatal(err)
		}
	})

	t.Run("Invalid call", func(t *testing.T) {
		db.On("Delete", mock.MatchedBy(func(bz []byte) bool {
			return bytes.Equal(bz, invalidPath)
		})).Return(errForcedError)

		_, err := store.DeletePassword(&prototypes.DeletePasswordRequest{
			Url:      invalidPassword.GetUrl(),
			Username: invalidPassword.GetUsername(),
		})
		if ok := assert.ErrorIs(t, err, errForcedError); !ok {
			t.Fatalf("expected: %v, got: %v", errForcedError, err)
		}
	})
}
