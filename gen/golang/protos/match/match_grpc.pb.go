// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protos/match/match.proto

package match

import (
	context "context"
	model "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MatchServiceClient is the client API for MatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchServiceClient interface {
	ConnectStream(ctx context.Context, opts ...grpc.CallOption) (MatchService_ConnectStreamClient, error)
	LogConnectStream(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*model.Result, error)
}

type matchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchServiceClient(cc grpc.ClientConnInterface) MatchServiceClient {
	return &matchServiceClient{cc}
}

func (c *matchServiceClient) ConnectStream(ctx context.Context, opts ...grpc.CallOption) (MatchService_ConnectStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &MatchService_ServiceDesc.Streams[0], "/match.MatchService/ConnectStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &matchServiceConnectStreamClient{stream}
	return x, nil
}

type MatchService_ConnectStreamClient interface {
	Send(*StreamReq) error
	Recv() (*StreamRes, error)
	grpc.ClientStream
}

type matchServiceConnectStreamClient struct {
	grpc.ClientStream
}

func (x *matchServiceConnectStreamClient) Send(m *StreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *matchServiceConnectStreamClient) Recv() (*StreamRes, error) {
	m := new(StreamRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *matchServiceClient) LogConnectStream(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*model.Result, error) {
	out := new(model.Result)
	err := c.cc.Invoke(ctx, "/match.MatchService/LogConnectStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchServiceServer is the server API for MatchService service.
// All implementations must embed UnimplementedMatchServiceServer
// for forward compatibility
type MatchServiceServer interface {
	ConnectStream(MatchService_ConnectStreamServer) error
	LogConnectStream(context.Context, *model.Empty) (*model.Result, error)
	mustEmbedUnimplementedMatchServiceServer()
}

// UnimplementedMatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMatchServiceServer struct {
}

func (UnimplementedMatchServiceServer) ConnectStream(MatchService_ConnectStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectStream not implemented")
}
func (UnimplementedMatchServiceServer) LogConnectStream(context.Context, *model.Empty) (*model.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogConnectStream not implemented")
}
func (UnimplementedMatchServiceServer) mustEmbedUnimplementedMatchServiceServer() {}

// UnsafeMatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchServiceServer will
// result in compilation errors.
type UnsafeMatchServiceServer interface {
	mustEmbedUnimplementedMatchServiceServer()
}

func RegisterMatchServiceServer(s grpc.ServiceRegistrar, srv MatchServiceServer) {
	s.RegisterService(&MatchService_ServiceDesc, srv)
}

func _MatchService_ConnectStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MatchServiceServer).ConnectStream(&matchServiceConnectStreamServer{stream})
}

type MatchService_ConnectStreamServer interface {
	Send(*StreamRes) error
	Recv() (*StreamReq, error)
	grpc.ServerStream
}

type matchServiceConnectStreamServer struct {
	grpc.ServerStream
}

func (x *matchServiceConnectStreamServer) Send(m *StreamRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *matchServiceConnectStreamServer) Recv() (*StreamReq, error) {
	m := new(StreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MatchService_LogConnectStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).LogConnectStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/match.MatchService/LogConnectStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).LogConnectStream(ctx, req.(*model.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MatchService_ServiceDesc is the grpc.ServiceDesc for MatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "match.MatchService",
	HandlerType: (*MatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogConnectStream",
			Handler:    _MatchService_LogConnectStream_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectStream",
			Handler:       _MatchService_ConnectStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/match/match.proto",
}
