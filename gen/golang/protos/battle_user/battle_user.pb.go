// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/battle_user/battle_user.proto

package battle_user

import (
	model "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRankerListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankerList []*model.Ranker `protobuf:"bytes,1,rep,name=ranker_list,json=rankerList,proto3" json:"ranker_list,omitempty"`
}

func (x *GetRankerListRes) Reset() {
	*x = GetRankerListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_battle_user_battle_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankerListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankerListRes) ProtoMessage() {}

func (x *GetRankerListRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_battle_user_battle_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankerListRes.ProtoReflect.Descriptor instead.
func (*GetRankerListRes) Descriptor() ([]byte, []int) {
	return file_protos_battle_user_battle_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetRankerListRes) GetRankerList() []*model.Ranker {
	if x != nil {
		return x.RankerList
	}
	return nil
}

var File_protos_battle_user_battle_user_proto protoreflect.FileDescriptor

var file_protos_battle_user_battle_user_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x12, 0x2e, 0x0a, 0x0b, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52,
	0x61, 0x6e, 0x6b, 0x65, 0x72, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x32, 0x53, 0x0a, 0x11, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e,
	0x6b, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x53, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x65, 0x6f, 0x6d, 0x2d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0xaa, 0x02, 0x0e, 0x47, 0x72, 0x70,
	0x63, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protos_battle_user_battle_user_proto_rawDescOnce sync.Once
	file_protos_battle_user_battle_user_proto_rawDescData = file_protos_battle_user_battle_user_proto_rawDesc
)

func file_protos_battle_user_battle_user_proto_rawDescGZIP() []byte {
	file_protos_battle_user_battle_user_proto_rawDescOnce.Do(func() {
		file_protos_battle_user_battle_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_battle_user_battle_user_proto_rawDescData)
	})
	return file_protos_battle_user_battle_user_proto_rawDescData
}

var file_protos_battle_user_battle_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_battle_user_battle_user_proto_goTypes = []interface{}{
	(*GetRankerListRes)(nil), // 0: battle_user.GetRankerListRes
	(*model.Ranker)(nil),     // 1: model.Ranker
	(*model.Empty)(nil),      // 2: model.Empty
}
var file_protos_battle_user_battle_user_proto_depIdxs = []int32{
	1, // 0: battle_user.GetRankerListRes.ranker_list:type_name -> model.Ranker
	2, // 1: battle_user.BattleUserService.GetRankerList:input_type -> model.Empty
	0, // 2: battle_user.BattleUserService.GetRankerList:output_type -> battle_user.GetRankerListRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_battle_user_battle_user_proto_init() }
func file_protos_battle_user_battle_user_proto_init() {
	if File_protos_battle_user_battle_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_battle_user_battle_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRankerListRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_battle_user_battle_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_battle_user_battle_user_proto_goTypes,
		DependencyIndexes: file_protos_battle_user_battle_user_proto_depIdxs,
		MessageInfos:      file_protos_battle_user_battle_user_proto_msgTypes,
	}.Build()
	File_protos_battle_user_battle_user_proto = out.File
	file_protos_battle_user_battle_user_proto_rawDesc = nil
	file_protos_battle_user_battle_user_proto_goTypes = nil
	file_protos_battle_user_battle_user_proto_depIdxs = nil
}
