// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0--rc2
// source: modules/player/playerPb/playerPb.proto

package Golang_Project

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

// PlayerGrpcServiceClient is the client API for PlayerGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerGrpcServiceClient interface {
	CreadentialSearch(ctx context.Context, in *CreadentialSearchRequest, opts ...grpc.CallOption) (*PlayerProfile, error)
	FindOnePlayerProfileToRefresh(ctx context.Context, in *FindOnePlayerProfileToRefreshRequest, opts ...grpc.CallOption) (*PlayerProfile, error)
	GetPlayerSavingAccout(ctx context.Context, in *GetPlayerSavingAccoutRequest, opts ...grpc.CallOption) (*GetPlayerSavingAccoutResponse, error)
}

type playerGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerGrpcServiceClient(cc grpc.ClientConnInterface) PlayerGrpcServiceClient {
	return &playerGrpcServiceClient{cc}
}

func (c *playerGrpcServiceClient) CreadentialSearch(ctx context.Context, in *CreadentialSearchRequest, opts ...grpc.CallOption) (*PlayerProfile, error) {
	out := new(PlayerProfile)
	err := c.cc.Invoke(ctx, "/PlayerGrpcService/CreadentialSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerGrpcServiceClient) FindOnePlayerProfileToRefresh(ctx context.Context, in *FindOnePlayerProfileToRefreshRequest, opts ...grpc.CallOption) (*PlayerProfile, error) {
	out := new(PlayerProfile)
	err := c.cc.Invoke(ctx, "/PlayerGrpcService/FindOnePlayerProfileToRefresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerGrpcServiceClient) GetPlayerSavingAccout(ctx context.Context, in *GetPlayerSavingAccoutRequest, opts ...grpc.CallOption) (*GetPlayerSavingAccoutResponse, error) {
	out := new(GetPlayerSavingAccoutResponse)
	err := c.cc.Invoke(ctx, "/PlayerGrpcService/GetPlayerSavingAccout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerGrpcServiceServer is the server API for PlayerGrpcService service.
// All implementations must embed UnimplementedPlayerGrpcServiceServer
// for forward compatibility
type PlayerGrpcServiceServer interface {
	CreadentialSearch(context.Context, *CreadentialSearchRequest) (*PlayerProfile, error)
	FindOnePlayerProfileToRefresh(context.Context, *FindOnePlayerProfileToRefreshRequest) (*PlayerProfile, error)
	GetPlayerSavingAccout(context.Context, *GetPlayerSavingAccoutRequest) (*GetPlayerSavingAccoutResponse, error)
	mustEmbedUnimplementedPlayerGrpcServiceServer()
}

// UnimplementedPlayerGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlayerGrpcServiceServer struct {
}

func (UnimplementedPlayerGrpcServiceServer) CreadentialSearch(context.Context, *CreadentialSearchRequest) (*PlayerProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreadentialSearch not implemented")
}
func (UnimplementedPlayerGrpcServiceServer) FindOnePlayerProfileToRefresh(context.Context, *FindOnePlayerProfileToRefreshRequest) (*PlayerProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOnePlayerProfileToRefresh not implemented")
}
func (UnimplementedPlayerGrpcServiceServer) GetPlayerSavingAccout(context.Context, *GetPlayerSavingAccoutRequest) (*GetPlayerSavingAccoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerSavingAccout not implemented")
}
func (UnimplementedPlayerGrpcServiceServer) mustEmbedUnimplementedPlayerGrpcServiceServer() {}

// UnsafePlayerGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerGrpcServiceServer will
// result in compilation errors.
type UnsafePlayerGrpcServiceServer interface {
	mustEmbedUnimplementedPlayerGrpcServiceServer()
}

func RegisterPlayerGrpcServiceServer(s grpc.ServiceRegistrar, srv PlayerGrpcServiceServer) {
	s.RegisterService(&PlayerGrpcService_ServiceDesc, srv)
}

func _PlayerGrpcService_CreadentialSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreadentialSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerGrpcServiceServer).CreadentialSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlayerGrpcService/CreadentialSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerGrpcServiceServer).CreadentialSearch(ctx, req.(*CreadentialSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerGrpcService_FindOnePlayerProfileToRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOnePlayerProfileToRefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerGrpcServiceServer).FindOnePlayerProfileToRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlayerGrpcService/FindOnePlayerProfileToRefresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerGrpcServiceServer).FindOnePlayerProfileToRefresh(ctx, req.(*FindOnePlayerProfileToRefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerGrpcService_GetPlayerSavingAccout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerSavingAccoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerGrpcServiceServer).GetPlayerSavingAccout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlayerGrpcService/GetPlayerSavingAccout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerGrpcServiceServer).GetPlayerSavingAccout(ctx, req.(*GetPlayerSavingAccoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayerGrpcService_ServiceDesc is the grpc.ServiceDesc for PlayerGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayerGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PlayerGrpcService",
	HandlerType: (*PlayerGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreadentialSearch",
			Handler:    _PlayerGrpcService_CreadentialSearch_Handler,
		},
		{
			MethodName: "FindOnePlayerProfileToRefresh",
			Handler:    _PlayerGrpcService_FindOnePlayerProfileToRefresh_Handler,
		},
		{
			MethodName: "GetPlayerSavingAccout",
			Handler:    _PlayerGrpcService_GetPlayerSavingAccout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/player/playerPb/playerPb.proto",
}
