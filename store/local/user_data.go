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

func (s *Store) GetUserData(
	req *prototypes.GetUserDataRequest,
) (*prototypes.GetUserDataResponse, error) {
	bz, err := s.db.Read(
		localdb.BuildUserDataKeyPath(req.GetUsername()),
	)
	if err != nil {
		return &prototypes.GetUserDataResponse{}, err
	}

	// Unmarshal the returned bytes into a data entry
	data := &prototypes.UserData{}
	if err := proto.Unmarshal(bz, data); err != nil {
		return nil, err
	}

	// Build the response
	return &prototypes.GetUserDataResponse{
		UserData: data,
	}, nil
}

func (s *Store) SetUserData(
	req *prototypes.SetUserDataRequest,
) (*prototypes.SetUserDataResponse, error) {
	if err := s.db.Write(
		localdb.BuildUserDataKeyPath(req.UserData.GetUsername()),
		req.UserData,
	); err != nil {
		return &prototypes.SetUserDataResponse{}, err
	}
	return &prototypes.SetUserDataResponse{
		UserData: req.UserData,
	}, nil
}
