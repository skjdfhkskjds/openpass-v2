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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.3
// source: proto/types/v1/user.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Password hash is a hash of the user's password.
	// This is used to authenticate the user.
	PasswordHash string `protobuf:"bytes,2,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	// Algorithm used to generate the password hash.
	Algorithm string `protobuf:"bytes,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	// Params stores relevant user metadata as well as
	// key values required to decode the password hash
	// for the given algorithm.
	Params map[string]string `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_proto_types_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserData) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

func (x *UserData) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *UserData) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

// GetUserDataRequest is a request to get user data
type GetUserDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetUserDataRequest) Reset() {
	*x = GetUserDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDataRequest) ProtoMessage() {}

func (x *GetUserDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDataRequest.ProtoReflect.Descriptor instead.
func (*GetUserDataRequest) Descriptor() ([]byte, []int) {
	return file_proto_types_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserDataRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// GetUserDataResponse is a response to a GetUserDataRequest
type GetUserDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserData *UserData `protobuf:"bytes,1,opt,name=user_data,json=userData,proto3,oneof" json:"user_data,omitempty"`
}

func (x *GetUserDataResponse) Reset() {
	*x = GetUserDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDataResponse) ProtoMessage() {}

func (x *GetUserDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDataResponse.ProtoReflect.Descriptor instead.
func (*GetUserDataResponse) Descriptor() ([]byte, []int) {
	return file_proto_types_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserDataResponse) GetUserData() *UserData {
	if x != nil {
		return x.UserData
	}
	return nil
}

// SetUserDataRequest is a request to set user data
type SetUserDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserData *UserData `protobuf:"bytes,1,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
}

func (x *SetUserDataRequest) Reset() {
	*x = SetUserDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserDataRequest) ProtoMessage() {}

func (x *SetUserDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserDataRequest.ProtoReflect.Descriptor instead.
func (*SetUserDataRequest) Descriptor() ([]byte, []int) {
	return file_proto_types_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *SetUserDataRequest) GetUserData() *UserData {
	if x != nil {
		return x.UserData
	}
	return nil
}

// SetUserDataResponse is a response to a SetUserDataRequest
type SetUserDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserData *UserData `protobuf:"bytes,1,opt,name=user_data,json=userData,proto3,oneof" json:"user_data,omitempty"`
}

func (x *SetUserDataResponse) Reset() {
	*x = SetUserDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserDataResponse) ProtoMessage() {}

func (x *SetUserDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserDataResponse.ProtoReflect.Descriptor instead.
func (*SetUserDataResponse) Descriptor() ([]byte, []int) {
	return file_proto_types_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *SetUserDataResponse) GetUserData() *UserData {
	if x != nil {
		return x.UserData
	}
	return nil
}

var File_proto_types_v1_user_proto protoreflect.FileDescriptor

var file_proto_types_v1_user_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xdc, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x30, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x59, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x45, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x59, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6b, 0x6a, 0x64, 0x66, 0x68, 0x6b, 0x73, 0x6b, 0x6a, 0x64, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_v1_user_proto_rawDescOnce sync.Once
	file_proto_types_v1_user_proto_rawDescData = file_proto_types_v1_user_proto_rawDesc
)

func file_proto_types_v1_user_proto_rawDescGZIP() []byte {
	file_proto_types_v1_user_proto_rawDescOnce.Do(func() {
		file_proto_types_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_v1_user_proto_rawDescData)
	})
	return file_proto_types_v1_user_proto_rawDescData
}

var file_proto_types_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_types_v1_user_proto_goTypes = []interface{}{
	(*UserData)(nil),            // 0: types.v1.UserData
	(*GetUserDataRequest)(nil),  // 1: types.v1.GetUserDataRequest
	(*GetUserDataResponse)(nil), // 2: types.v1.GetUserDataResponse
	(*SetUserDataRequest)(nil),  // 3: types.v1.SetUserDataRequest
	(*SetUserDataResponse)(nil), // 4: types.v1.SetUserDataResponse
	nil,                         // 5: types.v1.UserData.ParamsEntry
}
var file_proto_types_v1_user_proto_depIdxs = []int32{
	5, // 0: types.v1.UserData.params:type_name -> types.v1.UserData.ParamsEntry
	0, // 1: types.v1.GetUserDataResponse.user_data:type_name -> types.v1.UserData
	0, // 2: types.v1.SetUserDataRequest.user_data:type_name -> types.v1.UserData
	0, // 3: types.v1.SetUserDataResponse.user_data:type_name -> types.v1.UserData
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_types_v1_user_proto_init() }
func file_proto_types_v1_user_proto_init() {
	if File_proto_types_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_types_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_types_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_types_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserDataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_types_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserDataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_types_v1_user_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_proto_types_v1_user_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_types_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_v1_user_proto_goTypes,
		DependencyIndexes: file_proto_types_v1_user_proto_depIdxs,
		MessageInfos:      file_proto_types_v1_user_proto_msgTypes,
	}.Build()
	File_proto_types_v1_user_proto = out.File
	file_proto_types_v1_user_proto_rawDesc = nil
	file_proto_types_v1_user_proto_goTypes = nil
	file_proto_types_v1_user_proto_depIdxs = nil
}
