// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: proto/v1/user.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// Authenticate returns a token if the user is authenticated.
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// DeleteUser deletes a user.
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// Authenticate returns a token if the user is authenticated.
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	// CreateUser creates a new user.
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// DeleteUser deletes a user.
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _UserService_Authenticate_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/user.proto",
}

// UserDataServiceClient is the client API for UserDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserDataServiceClient interface {
	// GetPassword returns the user data for the given user data request.
	GetUserData(ctx context.Context, in *GetUserDataRequest, opts ...grpc.CallOption) (*GetUserDataResponse, error)
	// SetPassword sets the user data for the given user data request.
	SetUserData(ctx context.Context, in *SetUserDataRequest, opts ...grpc.CallOption) (*SetUserDataResponse, error)
}

type userDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserDataServiceClient(cc grpc.ClientConnInterface) UserDataServiceClient {
	return &userDataServiceClient{cc}
}

func (c *userDataServiceClient) GetUserData(ctx context.Context, in *GetUserDataRequest, opts ...grpc.CallOption) (*GetUserDataResponse, error) {
	out := new(GetUserDataResponse)
	err := c.cc.Invoke(ctx, "/v1.UserDataService/GetUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataServiceClient) SetUserData(ctx context.Context, in *SetUserDataRequest, opts ...grpc.CallOption) (*SetUserDataResponse, error) {
	out := new(SetUserDataResponse)
	err := c.cc.Invoke(ctx, "/v1.UserDataService/SetUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDataServiceServer is the server API for UserDataService service.
// All implementations must embed UnimplementedUserDataServiceServer
// for forward compatibility
type UserDataServiceServer interface {
	// GetPassword returns the user data for the given user data request.
	GetUserData(context.Context, *GetUserDataRequest) (*GetUserDataResponse, error)
	// SetPassword sets the user data for the given user data request.
	SetUserData(context.Context, *SetUserDataRequest) (*SetUserDataResponse, error)
	mustEmbedUnimplementedUserDataServiceServer()
}

// UnimplementedUserDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserDataServiceServer struct {
}

func (UnimplementedUserDataServiceServer) GetUserData(context.Context, *GetUserDataRequest) (*GetUserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserData not implemented")
}
func (UnimplementedUserDataServiceServer) SetUserData(context.Context, *SetUserDataRequest) (*SetUserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserData not implemented")
}
func (UnimplementedUserDataServiceServer) mustEmbedUnimplementedUserDataServiceServer() {}

// UnsafeUserDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserDataServiceServer will
// result in compilation errors.
type UnsafeUserDataServiceServer interface {
	mustEmbedUnimplementedUserDataServiceServer()
}

func RegisterUserDataServiceServer(s grpc.ServiceRegistrar, srv UserDataServiceServer) {
	s.RegisterService(&UserDataService_ServiceDesc, srv)
}

func _UserDataService_GetUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServiceServer).GetUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserDataService/GetUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServiceServer).GetUserData(ctx, req.(*GetUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataService_SetUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServiceServer).SetUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserDataService/SetUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServiceServer).SetUserData(ctx, req.(*SetUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserDataService_ServiceDesc is the grpc.ServiceDesc for UserDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UserDataService",
	HandlerType: (*UserDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserData",
			Handler:    _UserDataService_GetUserData_Handler,
		},
		{
			MethodName: "SetUserData",
			Handler:    _UserDataService_SetUserData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/user.proto",
}
