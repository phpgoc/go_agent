// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: user_list.proto

package agent_proto

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

// GetUserListClient is the client API for GetUserList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetUserListClient interface {
	GetUserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error)
}

type getUserListClient struct {
	cc grpc.ClientConnInterface
}

func NewGetUserListClient(cc grpc.ClientConnInterface) GetUserListClient {
	return &getUserListClient{cc}
}

func (c *getUserListClient) GetUserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/agent_proto.GetUserList/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetUserListServer is the server API for GetUserList service.
// All implementations must embed UnimplementedGetUserListServer
// for forward compatibility
type GetUserListServer interface {
	GetUserList(context.Context, *UserListRequest) (*UserListResponse, error)
	mustEmbedUnimplementedGetUserListServer()
}

// UnimplementedGetUserListServer must be embedded to have forward compatible implementations.
type UnimplementedGetUserListServer struct {
}

func (UnimplementedGetUserListServer) GetUserList(context.Context, *UserListRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedGetUserListServer) mustEmbedUnimplementedGetUserListServer() {}

// UnsafeGetUserListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetUserListServer will
// result in compilation errors.
type UnsafeGetUserListServer interface {
	mustEmbedUnimplementedGetUserListServer()
}

func RegisterGetUserListServer(s grpc.ServiceRegistrar, srv GetUserListServer) {
	s.RegisterService(&GetUserList_ServiceDesc, srv)
}

func _GetUserList_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetUserListServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent_proto.GetUserList/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetUserListServer).GetUserList(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetUserList_ServiceDesc is the grpc.ServiceDesc for GetUserList service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetUserList_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent_proto.GetUserList",
	HandlerType: (*GetUserListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _GetUserList_GetUserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_list.proto",
}
