// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/battle_result/battle_result.proto

package battle_result

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

type ConfirmBattleResultReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BattleChannelId string `protobuf:"bytes,1,opt,name=battle_channel_id,json=battleChannelId,proto3" json:"battle_channel_id,omitempty"`
	BattleResult    int32  `protobuf:"varint,2,opt,name=battle_result,json=battleResult,proto3" json:"battle_result,omitempty"`
}

func (x *ConfirmBattleResultReq) Reset() {
	*x = ConfirmBattleResultReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_battle_result_battle_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmBattleResultReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmBattleResultReq) ProtoMessage() {}

func (x *ConfirmBattleResultReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_battle_result_battle_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmBattleResultReq.ProtoReflect.Descriptor instead.
func (*ConfirmBattleResultReq) Descriptor() ([]byte, []int) {
	return file_protos_battle_result_battle_result_proto_rawDescGZIP(), []int{0}
}

func (x *ConfirmBattleResultReq) GetBattleChannelId() string {
	if x != nil {
		return x.BattleChannelId
	}
	return ""
}

func (x *ConfirmBattleResultReq) GetBattleResult() int32 {
	if x != nil {
		return x.BattleResult
	}
	return 0
}

type ConfirmBattleResultRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchPoint int32         `protobuf:"varint,1,opt,name=match_point,json=matchPoint,proto3" json:"match_point,omitempty"`
	AddPoint   int32         `protobuf:"varint,2,opt,name=add_point,json=addPoint,proto3" json:"add_point,omitempty"`
	Reward     *model.Reward `protobuf:"bytes,3,opt,name=reward,proto3" json:"reward,omitempty"`
}

func (x *ConfirmBattleResultRes) Reset() {
	*x = ConfirmBattleResultRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_battle_result_battle_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmBattleResultRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmBattleResultRes) ProtoMessage() {}

func (x *ConfirmBattleResultRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_battle_result_battle_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmBattleResultRes.ProtoReflect.Descriptor instead.
func (*ConfirmBattleResultRes) Descriptor() ([]byte, []int) {
	return file_protos_battle_result_battle_result_proto_rawDescGZIP(), []int{1}
}

func (x *ConfirmBattleResultRes) GetMatchPoint() int32 {
	if x != nil {
		return x.MatchPoint
	}
	return 0
}

func (x *ConfirmBattleResultRes) GetAddPoint() int32 {
	if x != nil {
		return x.AddPoint
	}
	return 0
}

func (x *ConfirmBattleResultRes) GetReward() *model.Reward {
	if x != nil {
		return x.Reward
	}
	return nil
}

var File_protos_battle_result_battle_result_proto protoreflect.FileDescriptor

var file_protos_battle_result_battle_result_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6c,
	0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x16, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x7d, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x64, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x32, 0x6a, 0x0a, 0x13, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x13, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x1c, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x42, 0x57, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x65, 0x6f, 0x6d, 0x2d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0xaa, 0x02, 0x10, 0x47, 0x72, 0x70, 0x63, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_battle_result_battle_result_proto_rawDescOnce sync.Once
	file_protos_battle_result_battle_result_proto_rawDescData = file_protos_battle_result_battle_result_proto_rawDesc
)

func file_protos_battle_result_battle_result_proto_rawDescGZIP() []byte {
	file_protos_battle_result_battle_result_proto_rawDescOnce.Do(func() {
		file_protos_battle_result_battle_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_battle_result_battle_result_proto_rawDescData)
	})
	return file_protos_battle_result_battle_result_proto_rawDescData
}

var file_protos_battle_result_battle_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_battle_result_battle_result_proto_goTypes = []interface{}{
	(*ConfirmBattleResultReq)(nil), // 0: mail.ConfirmBattleResultReq
	(*ConfirmBattleResultRes)(nil), // 1: mail.ConfirmBattleResultRes
	(*model.Reward)(nil),           // 2: model.Reward
}
var file_protos_battle_result_battle_result_proto_depIdxs = []int32{
	2, // 0: mail.ConfirmBattleResultRes.reward:type_name -> model.Reward
	0, // 1: mail.BattleResultService.ConfirmBattleResult:input_type -> mail.ConfirmBattleResultReq
	1, // 2: mail.BattleResultService.ConfirmBattleResult:output_type -> mail.ConfirmBattleResultRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_battle_result_battle_result_proto_init() }
func file_protos_battle_result_battle_result_proto_init() {
	if File_protos_battle_result_battle_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_battle_result_battle_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmBattleResultReq); i {
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
		file_protos_battle_result_battle_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmBattleResultRes); i {
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
			RawDescriptor: file_protos_battle_result_battle_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_battle_result_battle_result_proto_goTypes,
		DependencyIndexes: file_protos_battle_result_battle_result_proto_depIdxs,
		MessageInfos:      file_protos_battle_result_battle_result_proto_msgTypes,
	}.Build()
	File_protos_battle_result_battle_result_proto = out.File
	file_protos_battle_result_battle_result_proto_rawDesc = nil
	file_protos_battle_result_battle_result_proto_goTypes = nil
	file_protos_battle_result_battle_result_proto_depIdxs = nil
}
