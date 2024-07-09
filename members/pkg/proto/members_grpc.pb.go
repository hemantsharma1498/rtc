// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: members.proto

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

const (
	Connection_GetCommServerAddr_FullMethodName = "/helloworld.Connection/GetCommServerAddr"
)

// ConnectionClient is the client API for Connection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionClient interface {
	GetCommServerAddr(ctx context.Context, in *GetCommServerAddrReq, opts ...grpc.CallOption) (*GetCommServerAddrReply, error)
}

type connectionClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionClient(cc grpc.ClientConnInterface) ConnectionClient {
	return &connectionClient{cc}
}

func (c *connectionClient) GetCommServerAddr(ctx context.Context, in *GetCommServerAddrReq, opts ...grpc.CallOption) (*GetCommServerAddrReply, error) {
	out := new(GetCommServerAddrReply)
	err := c.cc.Invoke(ctx, Connection_GetCommServerAddr_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectionServer is the server API for Connection service.
// All implementations must embed UnimplementedConnectionServer
// for forward compatibility
type ConnectionServer interface {
	GetCommServerAddr(context.Context, *GetCommServerAddrReq) (*GetCommServerAddrReply, error)
	mustEmbedUnimplementedConnectionServer()
}

// UnimplementedConnectionServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionServer struct {
}

func (UnimplementedConnectionServer) GetCommServerAddr(context.Context, *GetCommServerAddrReq) (*GetCommServerAddrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommServerAddr not implemented")
}
func (UnimplementedConnectionServer) mustEmbedUnimplementedConnectionServer() {}

// UnsafeConnectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionServer will
// result in compilation errors.
type UnsafeConnectionServer interface {
	mustEmbedUnimplementedConnectionServer()
}

func RegisterConnectionServer(s grpc.ServiceRegistrar, srv ConnectionServer) {
	s.RegisterService(&Connection_ServiceDesc, srv)
}

func _Connection_GetCommServerAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommServerAddrReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServer).GetCommServerAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connection_GetCommServerAddr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServer).GetCommServerAddr(ctx, req.(*GetCommServerAddrReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Connection_ServiceDesc is the grpc.ServiceDesc for Connection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Connection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Connection",
	HandlerType: (*ConnectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCommServerAddr",
			Handler:    _Connection_GetCommServerAddr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "members.proto",
}
