// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConnClient is the client API for Conn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnClient interface {
}

type connClient struct {
	cc grpc.ClientConnInterface
}

func NewConnClient(cc grpc.ClientConnInterface) ConnClient {
	return &connClient{cc}
}

// ConnServer is the server API for Conn service.
// All implementations must embed UnimplementedConnServer
// for forward compatibility
type ConnServer interface {
	mustEmbedUnimplementedConnServer()
}

// UnimplementedConnServer must be embedded to have forward compatible implementations.
type UnimplementedConnServer struct {
}

func (UnimplementedConnServer) mustEmbedUnimplementedConnServer() {}

// UnsafeConnServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnServer will
// result in compilation errors.
type UnsafeConnServer interface {
	mustEmbedUnimplementedConnServer()
}

func RegisterConnServer(s grpc.ServiceRegistrar, srv ConnServer) {
	s.RegisterService(&Conn_ServiceDesc, srv)
}

// Conn_ServiceDesc is the grpc.ServiceDesc for Conn service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Conn_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Conn",
	HandlerType: (*ConnServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "v1/conn.proto",
}
