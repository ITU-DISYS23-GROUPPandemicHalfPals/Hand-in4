// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0--rc2
// source: me.proto

package me

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

const (
	MutualExclusion_Election_FullMethodName     = "/me.MutualExclusion/Election"
	MutualExclusion_Coordinator_FullMethodName  = "/me.MutualExclusion/Coordinator"
	MutualExclusion_RequestToken_FullMethodName = "/me.MutualExclusion/RequestToken"
	MutualExclusion_GrantToken_FullMethodName   = "/me.MutualExclusion/GrantToken"
	MutualExclusion_ReleaseToken_FullMethodName = "/me.MutualExclusion/ReleaseToken"
)

// MutualExclusionClient is the client API for MutualExclusion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MutualExclusionClient interface {
	Election(ctx context.Context, in *ElectionMessage, opts ...grpc.CallOption) (*Response, error)
	Coordinator(ctx context.Context, in *CoordinatorMessage, opts ...grpc.CallOption) (*Response, error)
	RequestToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*Response, error)
	GrantToken(ctx context.Context, in *TokenMessage, opts ...grpc.CallOption) (*Response, error)
	ReleaseToken(ctx context.Context, in *TokenMessage, opts ...grpc.CallOption) (*Response, error)
}

type mutualExclusionClient struct {
	cc grpc.ClientConnInterface
}

func NewMutualExclusionClient(cc grpc.ClientConnInterface) MutualExclusionClient {
	return &mutualExclusionClient{cc}
}

func (c *mutualExclusionClient) Election(ctx context.Context, in *ElectionMessage, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, MutualExclusion_Election_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutualExclusionClient) Coordinator(ctx context.Context, in *CoordinatorMessage, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, MutualExclusion_Coordinator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutualExclusionClient) RequestToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, MutualExclusion_RequestToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutualExclusionClient) GrantToken(ctx context.Context, in *TokenMessage, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, MutualExclusion_GrantToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutualExclusionClient) ReleaseToken(ctx context.Context, in *TokenMessage, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, MutualExclusion_ReleaseToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MutualExclusionServer is the server API for MutualExclusion service.
// All implementations must embed UnimplementedMutualExclusionServer
// for forward compatibility
type MutualExclusionServer interface {
	Election(context.Context, *ElectionMessage) (*Response, error)
	Coordinator(context.Context, *CoordinatorMessage) (*Response, error)
	RequestToken(context.Context, *TokenRequest) (*Response, error)
	GrantToken(context.Context, *TokenMessage) (*Response, error)
	ReleaseToken(context.Context, *TokenMessage) (*Response, error)
	mustEmbedUnimplementedMutualExclusionServer()
}

// UnimplementedMutualExclusionServer must be embedded to have forward compatible implementations.
type UnimplementedMutualExclusionServer struct {
}

func (UnimplementedMutualExclusionServer) Election(context.Context, *ElectionMessage) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Election not implemented")
}
func (UnimplementedMutualExclusionServer) Coordinator(context.Context, *CoordinatorMessage) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Coordinator not implemented")
}
func (UnimplementedMutualExclusionServer) RequestToken(context.Context, *TokenRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestToken not implemented")
}
func (UnimplementedMutualExclusionServer) GrantToken(context.Context, *TokenMessage) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantToken not implemented")
}
func (UnimplementedMutualExclusionServer) ReleaseToken(context.Context, *TokenMessage) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseToken not implemented")
}
func (UnimplementedMutualExclusionServer) mustEmbedUnimplementedMutualExclusionServer() {}

// UnsafeMutualExclusionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MutualExclusionServer will
// result in compilation errors.
type UnsafeMutualExclusionServer interface {
	mustEmbedUnimplementedMutualExclusionServer()
}

func RegisterMutualExclusionServer(s grpc.ServiceRegistrar, srv MutualExclusionServer) {
	s.RegisterService(&MutualExclusion_ServiceDesc, srv)
}

func _MutualExclusion_Election_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElectionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutualExclusionServer).Election(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MutualExclusion_Election_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutualExclusionServer).Election(ctx, req.(*ElectionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _MutualExclusion_Coordinator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoordinatorMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutualExclusionServer).Coordinator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MutualExclusion_Coordinator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutualExclusionServer).Coordinator(ctx, req.(*CoordinatorMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _MutualExclusion_RequestToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutualExclusionServer).RequestToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MutualExclusion_RequestToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutualExclusionServer).RequestToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MutualExclusion_GrantToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutualExclusionServer).GrantToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MutualExclusion_GrantToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutualExclusionServer).GrantToken(ctx, req.(*TokenMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _MutualExclusion_ReleaseToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutualExclusionServer).ReleaseToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MutualExclusion_ReleaseToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutualExclusionServer).ReleaseToken(ctx, req.(*TokenMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// MutualExclusion_ServiceDesc is the grpc.ServiceDesc for MutualExclusion service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MutualExclusion_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "me.MutualExclusion",
	HandlerType: (*MutualExclusionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Election",
			Handler:    _MutualExclusion_Election_Handler,
		},
		{
			MethodName: "Coordinator",
			Handler:    _MutualExclusion_Coordinator_Handler,
		},
		{
			MethodName: "RequestToken",
			Handler:    _MutualExclusion_RequestToken_Handler,
		},
		{
			MethodName: "GrantToken",
			Handler:    _MutualExclusion_GrantToken_Handler,
		},
		{
			MethodName: "ReleaseToken",
			Handler:    _MutualExclusion_ReleaseToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "me.proto",
}
