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
	"google.golang.org/protobuf/proto"

	localdb "github.com/skjdfhkskjds/openpass/v2/store/local/db"
	prototypes "github.com/skjdfhkskjds/openpass/v2/types/proto/v1"
)

func (s *Store) GetPassword(
	req *prototypes.GetPasswordRequest,
) (*prototypes.GetPasswordResponse, error) {
	bz, err := s.db.Read(
		localdb.BuildPasswordKeyPath(req.GetUrl(), req.GetUsername()),
	)
	if err != nil {
		return &prototypes.GetPasswordResponse{}, err
	}

	// Unmarshal the returned bytes into a password entry
	password := &prototypes.PasswordEntry{}
	if err := proto.Unmarshal(bz, password); err != nil {
		return nil, err
	}

	// Build the response
	return &prototypes.GetPasswordResponse{
		Password: password,
	}, nil
}

func (s *Store) SetPassword(
	req *prototypes.SetPasswordRequest,
) (*prototypes.SetPasswordResponse, error) {
	if err := s.db.Write(
		localdb.BuildPasswordKeyPath(
			req.Password.GetUrl(),
			req.Password.GetUsername(),
		),
		req.Password,
	); err != nil {
		return &prototypes.SetPasswordResponse{}, err
	}
	return &prototypes.SetPasswordResponse{
		Password: req.Password,
	}, nil
}

func (s *Store) UpdatePassword(
	req *prototypes.UpdatePasswordRequest,
) (*prototypes.UpdatePasswordResponse, error) {
	if err := s.db.Update(
		localdb.BuildPasswordKeyPath(
			req.Password.GetUrl(),
			req.Password.GetUsername(),
		),
		req.Password,
	); err != nil {
		return &prototypes.UpdatePasswordResponse{}, err
	}
	return &prototypes.UpdatePasswordResponse{
		Password: req.Password,
	}, nil
}

func (s *Store) DeletePassword(
	req *prototypes.DeletePasswordRequest,
) (*prototypes.DeletePasswordResponse, error) {
	return &prototypes.DeletePasswordResponse{},
		s.db.Delete(
			localdb.BuildPasswordKeyPath(
				req.GetUrl(),
				req.GetUsername(),
			),
		)
}
