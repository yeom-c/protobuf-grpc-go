// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protos/asset/asset.proto

package asset

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

// AssetServiceClient is the client API for AssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetServiceClient interface {
	GetAssets(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*GetAssetsRes, error)
}

type assetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetServiceClient(cc grpc.ClientConnInterface) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) GetAssets(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*GetAssetsRes, error) {
	out := new(GetAssetsRes)
	err := c.cc.Invoke(ctx, "/asset.AssetService/GetAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServiceServer is the server API for AssetService service.
// All implementations must embed UnimplementedAssetServiceServer
// for forward compatibility
type AssetServiceServer interface {
	GetAssets(context.Context, *model.Empty) (*GetAssetsRes, error)
	mustEmbedUnimplementedAssetServiceServer()
}

// UnimplementedAssetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssetServiceServer struct {
}

func (UnimplementedAssetServiceServer) GetAssets(context.Context, *model.Empty) (*GetAssetsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssets not implemented")
}
func (UnimplementedAssetServiceServer) mustEmbedUnimplementedAssetServiceServer() {}

// UnsafeAssetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetServiceServer will
// result in compilation errors.
type UnsafeAssetServiceServer interface {
	mustEmbedUnimplementedAssetServiceServer()
}

func RegisterAssetServiceServer(s grpc.ServiceRegistrar, srv AssetServiceServer) {
	s.RegisterService(&AssetService_ServiceDesc, srv)
}

func _AssetService_GetAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).GetAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asset.AssetService/GetAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).GetAssets(ctx, req.(*model.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetService_ServiceDesc is the grpc.ServiceDesc for AssetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "asset.AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAssets",
			Handler:    _AssetService_GetAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/asset/asset.proto",
}