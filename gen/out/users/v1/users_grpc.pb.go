// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package usersv1

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
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	GetUser(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error)
	BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*BlockUserResponse, error)
	CreateUserIdentity(ctx context.Context, in *CreateUserIdentityRequest, opts ...grpc.CallOption) (*CreateUserIdentityResponse, error)
	UpdateUserIdentity(ctx context.Context, in *UpdateUserIdentityRequest, opts ...grpc.CallOption) (*UpdateUserIdentityResponse, error)
	GetUserIdentities(ctx context.Context, in *GetUserIdentitiesRequest, opts ...grpc.CallOption) (*GetUserIdentitiesResponse, error)
	CreateUserRelation(ctx context.Context, in *CreateUserRelationRequest, opts ...grpc.CallOption) (*CreateUserRelationResponse, error)
	DeleteUserRelation(ctx context.Context, in *DeleteUserRelationRequest, opts ...grpc.CallOption) (*DeleteUserRelationResponse, error)
	CreateUserAddress(ctx context.Context, in *CreateUserAddressRequest, opts ...grpc.CallOption) (*CreateUserAddressResponse, error)
	UpdateUserAddress(ctx context.Context, in *UpdateUserAddressRequest, opts ...grpc.CallOption) (*UpdateUserAddressResponse, error)
	GetUserAddresses(ctx context.Context, in *GetUserAddressesRequest, opts ...grpc.CallOption) (*GetUserAddressesResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error) {
	out := new(GetUserByIdResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*BlockUserResponse, error) {
	out := new(BlockUserResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UserService/BlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUserIdentity(ctx context.Context, in *CreateUserIdentityRequest, opts ...grpc.CallOption) (*CreateUserIdentityResponse, error) {
	out := new(CreateUserIdentityResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UserService/CreateUserIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserIdentity(ctx context.Context, in *UpdateUserIdentityRequest, opts ...grpc.CallOption) (*UpdateUserIdentityResponse, error) {
	out := new(UpdateUserIdentityResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UserService/UpdateUserIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserIdentities(ctx context.Context, in *GetUserIdentitiesRequest, opts ...grpc.CallOption) (*GetUserIdentitiesResponse, error) {
	out := new(GetUserIdentitiesResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UserService/GetUserIdentities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUserRelation(ctx context.Context, in *CreateUserRelationRequest, opts ...grpc.CallOption) (*CreateUserRelationResponse, error) {
	out := new(CreateUserRelationResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UserService/CreateUserRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUserRelation(ctx context.Context, in *DeleteUserRelationRequest, opts ...grpc.CallOption) (*DeleteUserRelationResponse, error) {
	out := new(DeleteUserRelationResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UserService/DeleteUserRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUserAddress(ctx context.Context, in *CreateUserAddressRequest, opts ...grpc.CallOption) (*CreateUserAddressResponse, error) {
	out := new(CreateUserAddressResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UserService/CreateUserAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserAddress(ctx context.Context, in *UpdateUserAddressRequest, opts ...grpc.CallOption) (*UpdateUserAddressResponse, error) {
	out := new(UpdateUserAddressResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UserService/UpdateUserAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserAddresses(ctx context.Context, in *GetUserAddressesRequest, opts ...grpc.CallOption) (*GetUserAddressesResponse, error) {
	out := new(GetUserAddressesResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UserService/GetUserAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	GetUser(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error)
	BlockUser(context.Context, *BlockUserRequest) (*BlockUserResponse, error)
	CreateUserIdentity(context.Context, *CreateUserIdentityRequest) (*CreateUserIdentityResponse, error)
	UpdateUserIdentity(context.Context, *UpdateUserIdentityRequest) (*UpdateUserIdentityResponse, error)
	GetUserIdentities(context.Context, *GetUserIdentitiesRequest) (*GetUserIdentitiesResponse, error)
	CreateUserRelation(context.Context, *CreateUserRelationRequest) (*CreateUserRelationResponse, error)
	DeleteUserRelation(context.Context, *DeleteUserRelationRequest) (*DeleteUserRelationResponse, error)
	CreateUserAddress(context.Context, *CreateUserAddressRequest) (*CreateUserAddressResponse, error)
	UpdateUserAddress(context.Context, *UpdateUserAddressRequest) (*UpdateUserAddressResponse, error)
	GetUserAddresses(context.Context, *GetUserAddressesRequest) (*GetUserAddressesResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) BlockUser(context.Context, *BlockUserRequest) (*BlockUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (UnimplementedUserServiceServer) CreateUserIdentity(context.Context, *CreateUserIdentityRequest) (*CreateUserIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserIdentity not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserIdentity(context.Context, *UpdateUserIdentityRequest) (*UpdateUserIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserIdentity not implemented")
}
func (UnimplementedUserServiceServer) GetUserIdentities(context.Context, *GetUserIdentitiesRequest) (*GetUserIdentitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserIdentities not implemented")
}
func (UnimplementedUserServiceServer) CreateUserRelation(context.Context, *CreateUserRelationRequest) (*CreateUserRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserRelation not implemented")
}
func (UnimplementedUserServiceServer) DeleteUserRelation(context.Context, *DeleteUserRelationRequest) (*DeleteUserRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserRelation not implemented")
}
func (UnimplementedUserServiceServer) CreateUserAddress(context.Context, *CreateUserAddressRequest) (*CreateUserAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserAddress not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserAddress(context.Context, *UpdateUserAddressRequest) (*UpdateUserAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAddress not implemented")
}
func (UnimplementedUserServiceServer) GetUserAddresses(context.Context, *GetUserAddressesRequest) (*GetUserAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAddresses not implemented")
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
		FullMethod: "/users.v1.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UserService/BlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BlockUser(ctx, req.(*BlockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUserIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUserIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UserService/CreateUserIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUserIdentity(ctx, req.(*CreateUserIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UserService/UpdateUserIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserIdentity(ctx, req.(*UpdateUserIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserIdentities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserIdentitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserIdentities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UserService/GetUserIdentities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserIdentities(ctx, req.(*GetUserIdentitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUserRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUserRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UserService/CreateUserRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUserRelation(ctx, req.(*CreateUserRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUserRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUserRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UserService/DeleteUserRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUserRelation(ctx, req.(*DeleteUserRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUserAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUserAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UserService/CreateUserAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUserAddress(ctx, req.(*CreateUserAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UserService/UpdateUserAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserAddress(ctx, req.(*UpdateUserAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UserService/GetUserAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserAddresses(ctx, req.(*GetUserAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _UserService_BlockUser_Handler,
		},
		{
			MethodName: "CreateUserIdentity",
			Handler:    _UserService_CreateUserIdentity_Handler,
		},
		{
			MethodName: "UpdateUserIdentity",
			Handler:    _UserService_UpdateUserIdentity_Handler,
		},
		{
			MethodName: "GetUserIdentities",
			Handler:    _UserService_GetUserIdentities_Handler,
		},
		{
			MethodName: "CreateUserRelation",
			Handler:    _UserService_CreateUserRelation_Handler,
		},
		{
			MethodName: "DeleteUserRelation",
			Handler:    _UserService_DeleteUserRelation_Handler,
		},
		{
			MethodName: "CreateUserAddress",
			Handler:    _UserService_CreateUserAddress_Handler,
		},
		{
			MethodName: "UpdateUserAddress",
			Handler:    _UserService_UpdateUserAddress_Handler,
		},
		{
			MethodName: "GetUserAddresses",
			Handler:    _UserService_GetUserAddresses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users/v1/users.proto",
}