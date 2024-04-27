// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: api.proto

package server_buffer

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

// ServerBufferClient is the client API for ServerBuffer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerBufferClient interface {
	GetResponse(ctx context.Context, opts ...grpc.CallOption) (ServerBuffer_GetResponseClient, error)
}

type serverBufferClient struct {
	cc grpc.ClientConnInterface
}

func NewServerBufferClient(cc grpc.ClientConnInterface) ServerBufferClient {
	return &serverBufferClient{cc}
}

func (c *serverBufferClient) GetResponse(ctx context.Context, opts ...grpc.CallOption) (ServerBuffer_GetResponseClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServerBuffer_ServiceDesc.Streams[0], "/server_buffer.ServerBuffer/GetResponse", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverBufferGetResponseClient{stream}
	return x, nil
}

type ServerBuffer_GetResponseClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type serverBufferGetResponseClient struct {
	grpc.ClientStream
}

func (x *serverBufferGetResponseClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serverBufferGetResponseClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerBufferServer is the server API for ServerBuffer service.
// All implementations must embed UnimplementedServerBufferServer
// for forward compatibility
type ServerBufferServer interface {
	GetResponse(ServerBuffer_GetResponseServer) error
	mustEmbedUnimplementedServerBufferServer()
}

// UnimplementedServerBufferServer must be embedded to have forward compatible implementations.
type UnimplementedServerBufferServer struct {
}

func (UnimplementedServerBufferServer) GetResponse(ServerBuffer_GetResponseServer) error {
	return status.Errorf(codes.Unimplemented, "method GetResponse not implemented")
}
func (UnimplementedServerBufferServer) mustEmbedUnimplementedServerBufferServer() {}

// UnsafeServerBufferServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerBufferServer will
// result in compilation errors.
type UnsafeServerBufferServer interface {
	mustEmbedUnimplementedServerBufferServer()
}

func RegisterServerBufferServer(s grpc.ServiceRegistrar, srv ServerBufferServer) {
	s.RegisterService(&ServerBuffer_ServiceDesc, srv)
}

func _ServerBuffer_GetResponse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServerBufferServer).GetResponse(&serverBufferGetResponseServer{stream})
}

type ServerBuffer_GetResponseServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type serverBufferGetResponseServer struct {
	grpc.ServerStream
}

func (x *serverBufferGetResponseServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serverBufferGetResponseServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerBuffer_ServiceDesc is the grpc.ServiceDesc for ServerBuffer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerBuffer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server_buffer.ServerBuffer",
	HandlerType: (*ServerBufferServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetResponse",
			Handler:       _ServerBuffer_GetResponse_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}