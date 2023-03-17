// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/gacha_log/gacha_log.proto

package gacha_log

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

type GetGachaLogCategoriesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GachaLogCategories []string `protobuf:"bytes,1,rep,name=gacha_log_categories,json=gachaLogCategories,proto3" json:"gacha_log_categories,omitempty"`
}

func (x *GetGachaLogCategoriesRes) Reset() {
	*x = GetGachaLogCategoriesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_gacha_log_gacha_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGachaLogCategoriesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGachaLogCategoriesRes) ProtoMessage() {}

func (x *GetGachaLogCategoriesRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_gacha_log_gacha_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGachaLogCategoriesRes.ProtoReflect.Descriptor instead.
func (*GetGachaLogCategoriesRes) Descriptor() ([]byte, []int) {
	return file_protos_gacha_log_gacha_log_proto_rawDescGZIP(), []int{0}
}

func (x *GetGachaLogCategoriesRes) GetGachaLogCategories() []string {
	if x != nil {
		return x.GachaLogCategories
	}
	return nil
}

type GetGachaLogsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnumId string `protobuf:"bytes,1,opt,name=enum_id,json=enumId,proto3" json:"enum_id,omitempty"`
	Page   int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetGachaLogsReq) Reset() {
	*x = GetGachaLogsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_gacha_log_gacha_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGachaLogsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGachaLogsReq) ProtoMessage() {}

func (x *GetGachaLogsReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_gacha_log_gacha_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGachaLogsReq.ProtoReflect.Descriptor instead.
func (*GetGachaLogsReq) Descriptor() ([]byte, []int) {
	return file_protos_gacha_log_gacha_log_proto_rawDescGZIP(), []int{1}
}

func (x *GetGachaLogsReq) GetEnumId() string {
	if x != nil {
		return x.EnumId
	}
	return ""
}

func (x *GetGachaLogsReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetGachaLogsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GachaLogs []*model.GachaLogData `protobuf:"bytes,1,rep,name=gacha_logs,json=gachaLogs,proto3" json:"gacha_logs,omitempty"`
	TotalPage int64                 `protobuf:"varint,2,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
}

func (x *GetGachaLogsRes) Reset() {
	*x = GetGachaLogsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_gacha_log_gacha_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGachaLogsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGachaLogsRes) ProtoMessage() {}

func (x *GetGachaLogsRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_gacha_log_gacha_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGachaLogsRes.ProtoReflect.Descriptor instead.
func (*GetGachaLogsRes) Descriptor() ([]byte, []int) {
	return file_protos_gacha_log_gacha_log_proto_rawDescGZIP(), []int{2}
}

func (x *GetGachaLogsRes) GetGachaLogs() []*model.GachaLogData {
	if x != nil {
		return x.GachaLogs
	}
	return nil
}

func (x *GetGachaLogsRes) GetTotalPage() int64 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

var File_protos_gacha_log_gacha_log_proto protoreflect.FileDescriptor

var file_protos_gacha_log_gacha_log_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x6c,
	0x6f, 0x67, 0x2f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x6c, 0x6f, 0x67, 0x1a, 0x18, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x47, 0x61,
	0x63, 0x68, 0x61, 0x4c, 0x6f, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x6c, 0x6f, 0x67,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x12, 0x67, 0x61, 0x63, 0x68, 0x61, 0x4c, 0x6f, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x61, 0x63, 0x68,
	0x61, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e, 0x75, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x75, 0x6d, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x64, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x61, 0x63, 0x68,
	0x61, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x63, 0x68, 0x61, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x09, 0x67, 0x61, 0x63, 0x68, 0x61, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x32, 0xa9, 0x01, 0x0a, 0x0f,
	0x47, 0x61, 0x63, 0x68, 0x61, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x47, 0x61, 0x63, 0x68, 0x61, 0x4c, 0x6f, 0x67, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x6c,
	0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x63, 0x68, 0x61, 0x4c, 0x6f, 0x67, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x47, 0x61, 0x63, 0x68, 0x61, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1a, 0x2e,
	0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x63,
	0x68, 0x61, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x63, 0x68, 0x61, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x4f, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x65, 0x6f, 0x6d, 0x2d, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x6c, 0x6f, 0x67, 0xaa, 0x02, 0x0c, 0x47, 0x72, 0x70, 0x63,
	0x47, 0x61, 0x63, 0x68, 0x61, 0x4c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_gacha_log_gacha_log_proto_rawDescOnce sync.Once
	file_protos_gacha_log_gacha_log_proto_rawDescData = file_protos_gacha_log_gacha_log_proto_rawDesc
)

func file_protos_gacha_log_gacha_log_proto_rawDescGZIP() []byte {
	file_protos_gacha_log_gacha_log_proto_rawDescOnce.Do(func() {
		file_protos_gacha_log_gacha_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_gacha_log_gacha_log_proto_rawDescData)
	})
	return file_protos_gacha_log_gacha_log_proto_rawDescData
}

var file_protos_gacha_log_gacha_log_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_gacha_log_gacha_log_proto_goTypes = []interface{}{
	(*GetGachaLogCategoriesRes)(nil), // 0: gacha_log.GetGachaLogCategoriesRes
	(*GetGachaLogsReq)(nil),          // 1: gacha_log.GetGachaLogsReq
	(*GetGachaLogsRes)(nil),          // 2: gacha_log.GetGachaLogsRes
	(*model.GachaLogData)(nil),       // 3: model.GachaLogData
	(*model.Empty)(nil),              // 4: model.Empty
}
var file_protos_gacha_log_gacha_log_proto_depIdxs = []int32{
	3, // 0: gacha_log.GetGachaLogsRes.gacha_logs:type_name -> model.GachaLogData
	4, // 1: gacha_log.GachaLogService.GetGachaLogCategories:input_type -> model.Empty
	1, // 2: gacha_log.GachaLogService.GetGachaLogs:input_type -> gacha_log.GetGachaLogsReq
	0, // 3: gacha_log.GachaLogService.GetGachaLogCategories:output_type -> gacha_log.GetGachaLogCategoriesRes
	2, // 4: gacha_log.GachaLogService.GetGachaLogs:output_type -> gacha_log.GetGachaLogsRes
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_gacha_log_gacha_log_proto_init() }
func file_protos_gacha_log_gacha_log_proto_init() {
	if File_protos_gacha_log_gacha_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_gacha_log_gacha_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGachaLogCategoriesRes); i {
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
		file_protos_gacha_log_gacha_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGachaLogsReq); i {
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
		file_protos_gacha_log_gacha_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGachaLogsRes); i {
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
			RawDescriptor: file_protos_gacha_log_gacha_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_gacha_log_gacha_log_proto_goTypes,
		DependencyIndexes: file_protos_gacha_log_gacha_log_proto_depIdxs,
		MessageInfos:      file_protos_gacha_log_gacha_log_proto_msgTypes,
	}.Build()
	File_protos_gacha_log_gacha_log_proto = out.File
	file_protos_gacha_log_gacha_log_proto_rawDesc = nil
	file_protos_gacha_log_gacha_log_proto_goTypes = nil
	file_protos_gacha_log_gacha_log_proto_depIdxs = nil
}