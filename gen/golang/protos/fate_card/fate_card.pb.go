// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/fate_card/fate_card.proto

package fate_card

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

type GetFateCardsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FateCards []*model.FateCardData `protobuf:"bytes,1,rep,name=fate_cards,json=fateCards,proto3" json:"fate_cards,omitempty"`
}

func (x *GetFateCardsRes) Reset() {
	*x = GetFateCardsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_fate_card_fate_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFateCardsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFateCardsRes) ProtoMessage() {}

func (x *GetFateCardsRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fate_card_fate_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFateCardsRes.ProtoReflect.Descriptor instead.
func (*GetFateCardsRes) Descriptor() ([]byte, []int) {
	return file_protos_fate_card_fate_card_proto_rawDescGZIP(), []int{0}
}

func (x *GetFateCardsRes) GetFateCards() []*model.FateCardData {
	if x != nil {
		return x.FateCards
	}
	return nil
}

type EquipFateCardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FateCardId  int32 `protobuf:"varint,1,opt,name=fate_card_id,json=fateCardId,proto3" json:"fate_card_id,omitempty"`
	CharacterId int32 `protobuf:"varint,2,opt,name=character_id,json=characterId,proto3" json:"character_id,omitempty"`
}

func (x *EquipFateCardReq) Reset() {
	*x = EquipFateCardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_fate_card_fate_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipFateCardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipFateCardReq) ProtoMessage() {}

func (x *EquipFateCardReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fate_card_fate_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipFateCardReq.ProtoReflect.Descriptor instead.
func (*EquipFateCardReq) Descriptor() ([]byte, []int) {
	return file_protos_fate_card_fate_card_proto_rawDescGZIP(), []int{1}
}

func (x *EquipFateCardReq) GetFateCardId() int32 {
	if x != nil {
		return x.FateCardId
	}
	return 0
}

func (x *EquipFateCardReq) GetCharacterId() int32 {
	if x != nil {
		return x.CharacterId
	}
	return 0
}

type UnequipFateCardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FateCardId int32 `protobuf:"varint,1,opt,name=fate_card_id,json=fateCardId,proto3" json:"fate_card_id,omitempty"`
}

func (x *UnequipFateCardReq) Reset() {
	*x = UnequipFateCardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_fate_card_fate_card_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnequipFateCardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnequipFateCardReq) ProtoMessage() {}

func (x *UnequipFateCardReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fate_card_fate_card_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnequipFateCardReq.ProtoReflect.Descriptor instead.
func (*UnequipFateCardReq) Descriptor() ([]byte, []int) {
	return file_protos_fate_card_fate_card_proto_rawDescGZIP(), []int{2}
}

func (x *UnequipFateCardReq) GetFateCardId() int32 {
	if x != nil {
		return x.FateCardId
	}
	return 0
}

var File_protos_fate_card_fate_card_proto protoreflect.FileDescriptor

var file_protos_fate_card_fate_card_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x2f, 0x66, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x66, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x1a, 0x18, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x66, 0x61,
	0x74, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x09, 0x66, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x73, 0x22, 0x57,
	0x0a, 0x10, 0x45, 0x71, 0x75, 0x69, 0x70, 0x46, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x12, 0x55, 0x6e, 0x65, 0x71, 0x75,
	0x69, 0x70, 0x46, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a,
	0x0c, 0x66, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x32,
	0xce, 0x01, 0x0a, 0x0f, 0x46, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1a, 0x2e, 0x66, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x0d, 0x45, 0x71, 0x75, 0x69, 0x70, 0x46, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x1b, 0x2e, 0x66, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x46, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x0e, 0x55, 0x6e, 0x71, 0x75, 0x69, 0x70, 0x46, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x1d, 0x2e, 0x66, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x55, 0x6e, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x46, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x42, 0x4f, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x65, 0x6f, 0x6d, 0x2d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0xaa, 0x02, 0x0c, 0x47, 0x72, 0x70, 0x63, 0x46, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_fate_card_fate_card_proto_rawDescOnce sync.Once
	file_protos_fate_card_fate_card_proto_rawDescData = file_protos_fate_card_fate_card_proto_rawDesc
)

func file_protos_fate_card_fate_card_proto_rawDescGZIP() []byte {
	file_protos_fate_card_fate_card_proto_rawDescOnce.Do(func() {
		file_protos_fate_card_fate_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_fate_card_fate_card_proto_rawDescData)
	})
	return file_protos_fate_card_fate_card_proto_rawDescData
}

var file_protos_fate_card_fate_card_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_fate_card_fate_card_proto_goTypes = []interface{}{
	(*GetFateCardsRes)(nil),    // 0: fate_card.GetFateCardsRes
	(*EquipFateCardReq)(nil),   // 1: fate_card.EquipFateCardReq
	(*UnequipFateCardReq)(nil), // 2: fate_card.UnequipFateCardReq
	(*model.FateCardData)(nil), // 3: model.FateCardData
	(*model.Empty)(nil),        // 4: model.Empty
	(*model.Result)(nil),       // 5: model.Result
}
var file_protos_fate_card_fate_card_proto_depIdxs = []int32{
	3, // 0: fate_card.GetFateCardsRes.fate_cards:type_name -> model.FateCardData
	4, // 1: fate_card.FateCardService.GetFateCards:input_type -> model.Empty
	1, // 2: fate_card.FateCardService.EquipFateCard:input_type -> fate_card.EquipFateCardReq
	2, // 3: fate_card.FateCardService.UnquipFateCard:input_type -> fate_card.UnequipFateCardReq
	0, // 4: fate_card.FateCardService.GetFateCards:output_type -> fate_card.GetFateCardsRes
	5, // 5: fate_card.FateCardService.EquipFateCard:output_type -> model.Result
	5, // 6: fate_card.FateCardService.UnquipFateCard:output_type -> model.Result
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_fate_card_fate_card_proto_init() }
func file_protos_fate_card_fate_card_proto_init() {
	if File_protos_fate_card_fate_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_fate_card_fate_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFateCardsRes); i {
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
		file_protos_fate_card_fate_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipFateCardReq); i {
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
		file_protos_fate_card_fate_card_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnequipFateCardReq); i {
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
			RawDescriptor: file_protos_fate_card_fate_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_fate_card_fate_card_proto_goTypes,
		DependencyIndexes: file_protos_fate_card_fate_card_proto_depIdxs,
		MessageInfos:      file_protos_fate_card_fate_card_proto_msgTypes,
	}.Build()
	File_protos_fate_card_fate_card_proto = out.File
	file_protos_fate_card_fate_card_proto_rawDesc = nil
	file_protos_fate_card_fate_card_proto_goTypes = nil
	file_protos_fate_card_fate_card_proto_depIdxs = nil
}
