// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protos/fate_card/fate_card.proto

package fate_card

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

// FateCardServiceClient is the client API for FateCardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FateCardServiceClient interface {
	GetFateCards(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*GetFateCardsRes, error)
	EquipFateCard(ctx context.Context, in *EquipFateCardReq, opts ...grpc.CallOption) (*model.Result, error)
	UnquipFateCard(ctx context.Context, in *UnequipFateCardReq, opts ...grpc.CallOption) (*model.Result, error)
}

type fateCardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFateCardServiceClient(cc grpc.ClientConnInterface) FateCardServiceClient {
	return &fateCardServiceClient{cc}
}

func (c *fateCardServiceClient) GetFateCards(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*GetFateCardsRes, error) {
	out := new(GetFateCardsRes)
	err := c.cc.Invoke(ctx, "/fate_card.FateCardService/GetFateCards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fateCardServiceClient) EquipFateCard(ctx context.Context, in *EquipFateCardReq, opts ...grpc.CallOption) (*model.Result, error) {
	out := new(model.Result)
	err := c.cc.Invoke(ctx, "/fate_card.FateCardService/EquipFateCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fateCardServiceClient) UnquipFateCard(ctx context.Context, in *UnequipFateCardReq, opts ...grpc.CallOption) (*model.Result, error) {
	out := new(model.Result)
	err := c.cc.Invoke(ctx, "/fate_card.FateCardService/UnquipFateCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FateCardServiceServer is the server API for FateCardService service.
// All implementations must embed UnimplementedFateCardServiceServer
// for forward compatibility
type FateCardServiceServer interface {
	GetFateCards(context.Context, *model.Empty) (*GetFateCardsRes, error)
	EquipFateCard(context.Context, *EquipFateCardReq) (*model.Result, error)
	UnquipFateCard(context.Context, *UnequipFateCardReq) (*model.Result, error)
	mustEmbedUnimplementedFateCardServiceServer()
}

// UnimplementedFateCardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFateCardServiceServer struct {
}

func (UnimplementedFateCardServiceServer) GetFateCards(context.Context, *model.Empty) (*GetFateCardsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFateCards not implemented")
}
func (UnimplementedFateCardServiceServer) EquipFateCard(context.Context, *EquipFateCardReq) (*model.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EquipFateCard not implemented")
}
func (UnimplementedFateCardServiceServer) UnquipFateCard(context.Context, *UnequipFateCardReq) (*model.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnquipFateCard not implemented")
}
func (UnimplementedFateCardServiceServer) mustEmbedUnimplementedFateCardServiceServer() {}

// UnsafeFateCardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FateCardServiceServer will
// result in compilation errors.
type UnsafeFateCardServiceServer interface {
	mustEmbedUnimplementedFateCardServiceServer()
}

func RegisterFateCardServiceServer(s grpc.ServiceRegistrar, srv FateCardServiceServer) {
	s.RegisterService(&FateCardService_ServiceDesc, srv)
}

func _FateCardService_GetFateCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FateCardServiceServer).GetFateCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fate_card.FateCardService/GetFateCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FateCardServiceServer).GetFateCards(ctx, req.(*model.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FateCardService_EquipFateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EquipFateCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FateCardServiceServer).EquipFateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fate_card.FateCardService/EquipFateCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FateCardServiceServer).EquipFateCard(ctx, req.(*EquipFateCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FateCardService_UnquipFateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnequipFateCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FateCardServiceServer).UnquipFateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fate_card.FateCardService/UnquipFateCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FateCardServiceServer).UnquipFateCard(ctx, req.(*UnequipFateCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FateCardService_ServiceDesc is the grpc.ServiceDesc for FateCardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FateCardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fate_card.FateCardService",
	HandlerType: (*FateCardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFateCards",
			Handler:    _FateCardService_GetFateCards_Handler,
		},
		{
			MethodName: "EquipFateCard",
			Handler:    _FateCardService_EquipFateCard_Handler,
		},
		{
			MethodName: "UnquipFateCard",
			Handler:    _FateCardService_UnquipFateCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/fate_card/fate_card.proto",
}
