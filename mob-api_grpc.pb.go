// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: mob-api.proto

package gerns_world_protobuf

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

// MobServiceClient is the client API for MobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MobServiceClient interface {
	CreateMobStream(ctx context.Context, opts ...grpc.CallOption) (MobService_CreateMobStreamClient, error)
}

type mobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMobServiceClient(cc grpc.ClientConnInterface) MobServiceClient {
	return &mobServiceClient{cc}
}

func (c *mobServiceClient) CreateMobStream(ctx context.Context, opts ...grpc.CallOption) (MobService_CreateMobStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &MobService_ServiceDesc.Streams[0], "/gernsworld.MobService/CreateMobStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &mobServiceCreateMobStreamClient{stream}
	return x, nil
}

type MobService_CreateMobStreamClient interface {
	Send(*MobStreamRequest) error
	Recv() (*MobStreamResponse, error)
	grpc.ClientStream
}

type mobServiceCreateMobStreamClient struct {
	grpc.ClientStream
}

func (x *mobServiceCreateMobStreamClient) Send(m *MobStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mobServiceCreateMobStreamClient) Recv() (*MobStreamResponse, error) {
	m := new(MobStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MobServiceServer is the server API for MobService service.
// All implementations must embed UnimplementedMobServiceServer
// for forward compatibility
type MobServiceServer interface {
	CreateMobStream(MobService_CreateMobStreamServer) error
	mustEmbedUnimplementedMobServiceServer()
}

// UnimplementedMobServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMobServiceServer struct {
}

func (UnimplementedMobServiceServer) CreateMobStream(MobService_CreateMobStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateMobStream not implemented")
}
func (UnimplementedMobServiceServer) mustEmbedUnimplementedMobServiceServer() {}

// UnsafeMobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MobServiceServer will
// result in compilation errors.
type UnsafeMobServiceServer interface {
	mustEmbedUnimplementedMobServiceServer()
}

func RegisterMobServiceServer(s grpc.ServiceRegistrar, srv MobServiceServer) {
	s.RegisterService(&MobService_ServiceDesc, srv)
}

func _MobService_CreateMobStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MobServiceServer).CreateMobStream(&mobServiceCreateMobStreamServer{stream})
}

type MobService_CreateMobStreamServer interface {
	Send(*MobStreamResponse) error
	Recv() (*MobStreamRequest, error)
	grpc.ServerStream
}

type mobServiceCreateMobStreamServer struct {
	grpc.ServerStream
}

func (x *mobServiceCreateMobStreamServer) Send(m *MobStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mobServiceCreateMobStreamServer) Recv() (*MobStreamRequest, error) {
	m := new(MobStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MobService_ServiceDesc is the grpc.ServiceDesc for MobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gernsworld.MobService",
	HandlerType: (*MobServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateMobStream",
			Handler:       _MobService_CreateMobStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mob-api.proto",
}
