// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protos/character/character.proto

package character

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

// CharacterServiceClient is the client API for CharacterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CharacterServiceClient interface {
	GetCharacters(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*GetCharactersRes, error)
	LevelUpSignatureWeapon(ctx context.Context, in *LevelUpSignatureWeaponReq, opts ...grpc.CallOption) (*model.Result, error)
	ExtinctCharacter(ctx context.Context, in *ExtinctCharacterReq, opts ...grpc.CallOption) (*ExtinctCharacterRes, error)
}

type characterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCharacterServiceClient(cc grpc.ClientConnInterface) CharacterServiceClient {
	return &characterServiceClient{cc}
}

func (c *characterServiceClient) GetCharacters(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*GetCharactersRes, error) {
	out := new(GetCharactersRes)
	err := c.cc.Invoke(ctx, "/character.CharacterService/GetCharacters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterServiceClient) LevelUpSignatureWeapon(ctx context.Context, in *LevelUpSignatureWeaponReq, opts ...grpc.CallOption) (*model.Result, error) {
	out := new(model.Result)
	err := c.cc.Invoke(ctx, "/character.CharacterService/LevelUpSignatureWeapon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterServiceClient) ExtinctCharacter(ctx context.Context, in *ExtinctCharacterReq, opts ...grpc.CallOption) (*ExtinctCharacterRes, error) {
	out := new(ExtinctCharacterRes)
	err := c.cc.Invoke(ctx, "/character.CharacterService/ExtinctCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharacterServiceServer is the server API for CharacterService service.
// All implementations must embed UnimplementedCharacterServiceServer
// for forward compatibility
type CharacterServiceServer interface {
	GetCharacters(context.Context, *model.Empty) (*GetCharactersRes, error)
	LevelUpSignatureWeapon(context.Context, *LevelUpSignatureWeaponReq) (*model.Result, error)
	ExtinctCharacter(context.Context, *ExtinctCharacterReq) (*ExtinctCharacterRes, error)
	mustEmbedUnimplementedCharacterServiceServer()
}

// UnimplementedCharacterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCharacterServiceServer struct {
}

func (UnimplementedCharacterServiceServer) GetCharacters(context.Context, *model.Empty) (*GetCharactersRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCharacters not implemented")
}
func (UnimplementedCharacterServiceServer) LevelUpSignatureWeapon(context.Context, *LevelUpSignatureWeaponReq) (*model.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LevelUpSignatureWeapon not implemented")
}
func (UnimplementedCharacterServiceServer) ExtinctCharacter(context.Context, *ExtinctCharacterReq) (*ExtinctCharacterRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtinctCharacter not implemented")
}
func (UnimplementedCharacterServiceServer) mustEmbedUnimplementedCharacterServiceServer() {}

// UnsafeCharacterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CharacterServiceServer will
// result in compilation errors.
type UnsafeCharacterServiceServer interface {
	mustEmbedUnimplementedCharacterServiceServer()
}

func RegisterCharacterServiceServer(s grpc.ServiceRegistrar, srv CharacterServiceServer) {
	s.RegisterService(&CharacterService_ServiceDesc, srv)
}

func _CharacterService_GetCharacters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).GetCharacters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/character.CharacterService/GetCharacters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).GetCharacters(ctx, req.(*model.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterService_LevelUpSignatureWeapon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LevelUpSignatureWeaponReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).LevelUpSignatureWeapon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/character.CharacterService/LevelUpSignatureWeapon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).LevelUpSignatureWeapon(ctx, req.(*LevelUpSignatureWeaponReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterService_ExtinctCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtinctCharacterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).ExtinctCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/character.CharacterService/ExtinctCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).ExtinctCharacter(ctx, req.(*ExtinctCharacterReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CharacterService_ServiceDesc is the grpc.ServiceDesc for CharacterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CharacterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "character.CharacterService",
	HandlerType: (*CharacterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCharacters",
			Handler:    _CharacterService_GetCharacters_Handler,
		},
		{
			MethodName: "LevelUpSignatureWeapon",
			Handler:    _CharacterService_LevelUpSignatureWeapon_Handler,
		},
		{
			MethodName: "ExtinctCharacter",
			Handler:    _CharacterService_ExtinctCharacter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/character/character.proto",
}
