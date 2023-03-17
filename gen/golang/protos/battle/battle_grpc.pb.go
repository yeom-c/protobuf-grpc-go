// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protos/battle/battle.proto

package battle

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

// BattleServiceClient is the client API for BattleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BattleServiceClient interface {
	ConnectStream(ctx context.Context, opts ...grpc.CallOption) (BattleService_ConnectStreamClient, error)
}

type battleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBattleServiceClient(cc grpc.ClientConnInterface) BattleServiceClient {
	return &battleServiceClient{cc}
}

func (c *battleServiceClient) ConnectStream(ctx context.Context, opts ...grpc.CallOption) (BattleService_ConnectStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &BattleService_ServiceDesc.Streams[0], "/battle.BattleService/ConnectStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &battleServiceConnectStreamClient{stream}
	return x, nil
}

type BattleService_ConnectStreamClient interface {
	Send(*StreamReq) error
	Recv() (*StreamRes, error)
	grpc.ClientStream
}

type battleServiceConnectStreamClient struct {
	grpc.ClientStream
}

func (x *battleServiceConnectStreamClient) Send(m *StreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *battleServiceConnectStreamClient) Recv() (*StreamRes, error) {
	m := new(StreamRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BattleServiceServer is the server API for BattleService service.
// All implementations must embed UnimplementedBattleServiceServer
// for forward compatibility
type BattleServiceServer interface {
	ConnectStream(BattleService_ConnectStreamServer) error
	mustEmbedUnimplementedBattleServiceServer()
}

// UnimplementedBattleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBattleServiceServer struct {
}

func (UnimplementedBattleServiceServer) ConnectStream(BattleService_ConnectStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectStream not implemented")
}
func (UnimplementedBattleServiceServer) mustEmbedUnimplementedBattleServiceServer() {}

// UnsafeBattleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BattleServiceServer will
// result in compilation errors.
type UnsafeBattleServiceServer interface {
	mustEmbedUnimplementedBattleServiceServer()
}

func RegisterBattleServiceServer(s grpc.ServiceRegistrar, srv BattleServiceServer) {
	s.RegisterService(&BattleService_ServiceDesc, srv)
}

func _BattleService_ConnectStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BattleServiceServer).ConnectStream(&battleServiceConnectStreamServer{stream})
}

type BattleService_ConnectStreamServer interface {
	Send(*StreamRes) error
	Recv() (*StreamReq, error)
	grpc.ServerStream
}

type battleServiceConnectStreamServer struct {
	grpc.ServerStream
}

func (x *battleServiceConnectStreamServer) Send(m *StreamRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *battleServiceConnectStreamServer) Recv() (*StreamReq, error) {
	m := new(StreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BattleService_ServiceDesc is the grpc.ServiceDesc for BattleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BattleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "battle.BattleService",
	HandlerType: (*BattleServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectStream",
			Handler:       _BattleService_ConnectStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/battle/battle.proto",
}