// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protos/character_broadcast/character_broadcast.proto

package character_broadcast

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

// CharacterBroadcastServiceClient is the client API for CharacterBroadcastService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CharacterBroadcastServiceClient interface {
	GetOnAirCharacterBroadcasts(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*GetOnAirCharacterBroadcastsRes, error)
	GetCompletedCharacterBroadcasts(ctx context.Context, in *GetCompletedCharacterBroadcastsReq, opts ...grpc.CallOption) (*GetCompletedCharacterBroadcastsRes, error)
	CompleteCharacterBroadcast(ctx context.Context, in *CompleteCharacterBroadcastReq, opts ...grpc.CallOption) (*CompleteCharacterBroadcastRes, error)
}

type characterBroadcastServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCharacterBroadcastServiceClient(cc grpc.ClientConnInterface) CharacterBroadcastServiceClient {
	return &characterBroadcastServiceClient{cc}
}

func (c *characterBroadcastServiceClient) GetOnAirCharacterBroadcasts(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*GetOnAirCharacterBroadcastsRes, error) {
	out := new(GetOnAirCharacterBroadcastsRes)
	err := c.cc.Invoke(ctx, "/character_broadcast.CharacterBroadcastService/GetOnAirCharacterBroadcasts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterBroadcastServiceClient) GetCompletedCharacterBroadcasts(ctx context.Context, in *GetCompletedCharacterBroadcastsReq, opts ...grpc.CallOption) (*GetCompletedCharacterBroadcastsRes, error) {
	out := new(GetCompletedCharacterBroadcastsRes)
	err := c.cc.Invoke(ctx, "/character_broadcast.CharacterBroadcastService/GetCompletedCharacterBroadcasts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterBroadcastServiceClient) CompleteCharacterBroadcast(ctx context.Context, in *CompleteCharacterBroadcastReq, opts ...grpc.CallOption) (*CompleteCharacterBroadcastRes, error) {
	out := new(CompleteCharacterBroadcastRes)
	err := c.cc.Invoke(ctx, "/character_broadcast.CharacterBroadcastService/CompleteCharacterBroadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharacterBroadcastServiceServer is the server API for CharacterBroadcastService service.
// All implementations must embed UnimplementedCharacterBroadcastServiceServer
// for forward compatibility
type CharacterBroadcastServiceServer interface {
	GetOnAirCharacterBroadcasts(context.Context, *model.Empty) (*GetOnAirCharacterBroadcastsRes, error)
	GetCompletedCharacterBroadcasts(context.Context, *GetCompletedCharacterBroadcastsReq) (*GetCompletedCharacterBroadcastsRes, error)
	CompleteCharacterBroadcast(context.Context, *CompleteCharacterBroadcastReq) (*CompleteCharacterBroadcastRes, error)
	mustEmbedUnimplementedCharacterBroadcastServiceServer()
}

// UnimplementedCharacterBroadcastServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCharacterBroadcastServiceServer struct {
}

func (UnimplementedCharacterBroadcastServiceServer) GetOnAirCharacterBroadcasts(context.Context, *model.Empty) (*GetOnAirCharacterBroadcastsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnAirCharacterBroadcasts not implemented")
}
func (UnimplementedCharacterBroadcastServiceServer) GetCompletedCharacterBroadcasts(context.Context, *GetCompletedCharacterBroadcastsReq) (*GetCompletedCharacterBroadcastsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompletedCharacterBroadcasts not implemented")
}
func (UnimplementedCharacterBroadcastServiceServer) CompleteCharacterBroadcast(context.Context, *CompleteCharacterBroadcastReq) (*CompleteCharacterBroadcastRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteCharacterBroadcast not implemented")
}
func (UnimplementedCharacterBroadcastServiceServer) mustEmbedUnimplementedCharacterBroadcastServiceServer() {
}

// UnsafeCharacterBroadcastServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CharacterBroadcastServiceServer will
// result in compilation errors.
type UnsafeCharacterBroadcastServiceServer interface {
	mustEmbedUnimplementedCharacterBroadcastServiceServer()
}

func RegisterCharacterBroadcastServiceServer(s grpc.ServiceRegistrar, srv CharacterBroadcastServiceServer) {
	s.RegisterService(&CharacterBroadcastService_ServiceDesc, srv)
}

func _CharacterBroadcastService_GetOnAirCharacterBroadcasts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterBroadcastServiceServer).GetOnAirCharacterBroadcasts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/character_broadcast.CharacterBroadcastService/GetOnAirCharacterBroadcasts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterBroadcastServiceServer).GetOnAirCharacterBroadcasts(ctx, req.(*model.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterBroadcastService_GetCompletedCharacterBroadcasts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompletedCharacterBroadcastsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterBroadcastServiceServer).GetCompletedCharacterBroadcasts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/character_broadcast.CharacterBroadcastService/GetCompletedCharacterBroadcasts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterBroadcastServiceServer).GetCompletedCharacterBroadcasts(ctx, req.(*GetCompletedCharacterBroadcastsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterBroadcastService_CompleteCharacterBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteCharacterBroadcastReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterBroadcastServiceServer).CompleteCharacterBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/character_broadcast.CharacterBroadcastService/CompleteCharacterBroadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterBroadcastServiceServer).CompleteCharacterBroadcast(ctx, req.(*CompleteCharacterBroadcastReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CharacterBroadcastService_ServiceDesc is the grpc.ServiceDesc for CharacterBroadcastService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CharacterBroadcastService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "character_broadcast.CharacterBroadcastService",
	HandlerType: (*CharacterBroadcastServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOnAirCharacterBroadcasts",
			Handler:    _CharacterBroadcastService_GetOnAirCharacterBroadcasts_Handler,
		},
		{
			MethodName: "GetCompletedCharacterBroadcasts",
			Handler:    _CharacterBroadcastService_GetCompletedCharacterBroadcasts_Handler,
		},
		{
			MethodName: "CompleteCharacterBroadcast",
			Handler:    _CharacterBroadcastService_CompleteCharacterBroadcast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/character_broadcast/character_broadcast.proto",
}
