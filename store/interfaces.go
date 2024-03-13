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

package store

import (
	proto "github.com/skjdfhkskjds/openpass/v2/types/proto/v1"
)

// Store is an interface for getting and setting key-value pairs
type Store interface {
	Close() error

	// For user profile data
	GetUserData(req *proto.GetUserDataRequest) (*proto.GetUserDataResponse, error)
	SetUserData(req *proto.SetUserDataRequest) (*proto.SetUserDataResponse, error)

	// For password data
	GetPassword(req *proto.GetPasswordRequest) (*proto.GetPasswordResponse, error)
	SetPassword(req *proto.SetPasswordRequest) (*proto.SetPasswordResponse, error)
	UpdatePassword(req *proto.UpdatePasswordRequest) (*proto.UpdatePasswordResponse, error)
	DeletePassword(req *proto.DeletePasswordRequest) (*proto.DeletePasswordResponse, error)
}
