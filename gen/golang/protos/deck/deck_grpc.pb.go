// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protos/deck/deck.proto

package deck

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

// DeckServiceClient is the client API for DeckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeckServiceClient interface {
	GetDecks(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*GetDecksRes, error)
	SaveDeck(ctx context.Context, in *SaveDeckReq, opts ...grpc.CallOption) (*SaveDeckRes, error)
}

type deckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeckServiceClient(cc grpc.ClientConnInterface) DeckServiceClient {
	return &deckServiceClient{cc}
}

func (c *deckServiceClient) GetDecks(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*GetDecksRes, error) {
	out := new(GetDecksRes)
	err := c.cc.Invoke(ctx, "/deck.DeckService/GetDecks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deckServiceClient) SaveDeck(ctx context.Context, in *SaveDeckReq, opts ...grpc.CallOption) (*SaveDeckRes, error) {
	out := new(SaveDeckRes)
	err := c.cc.Invoke(ctx, "/deck.DeckService/SaveDeck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeckServiceServer is the server API for DeckService service.
// All implementations must embed UnimplementedDeckServiceServer
// for forward compatibility
type DeckServiceServer interface {
	GetDecks(context.Context, *model.Empty) (*GetDecksRes, error)
	SaveDeck(context.Context, *SaveDeckReq) (*SaveDeckRes, error)
	mustEmbedUnimplementedDeckServiceServer()
}

// UnimplementedDeckServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeckServiceServer struct {
}

func (UnimplementedDeckServiceServer) GetDecks(context.Context, *model.Empty) (*GetDecksRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDecks not implemented")
}
func (UnimplementedDeckServiceServer) SaveDeck(context.Context, *SaveDeckReq) (*SaveDeckRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveDeck not implemented")
}
func (UnimplementedDeckServiceServer) mustEmbedUnimplementedDeckServiceServer() {}

// UnsafeDeckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeckServiceServer will
// result in compilation errors.
type UnsafeDeckServiceServer interface {
	mustEmbedUnimplementedDeckServiceServer()
}

func RegisterDeckServiceServer(s grpc.ServiceRegistrar, srv DeckServiceServer) {
	s.RegisterService(&DeckService_ServiceDesc, srv)
}

func _DeckService_GetDecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeckServiceServer).GetDecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deck.DeckService/GetDecks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeckServiceServer).GetDecks(ctx, req.(*model.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeckService_SaveDeck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveDeckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeckServiceServer).SaveDeck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deck.DeckService/SaveDeck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeckServiceServer).SaveDeck(ctx, req.(*SaveDeckReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DeckService_ServiceDesc is the grpc.ServiceDesc for DeckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "deck.DeckService",
	HandlerType: (*DeckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDecks",
			Handler:    _DeckService_GetDecks_Handler,
		},
		{
			MethodName: "SaveDeck",
			Handler:    _DeckService_SaveDeck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/deck/deck.proto",
}
