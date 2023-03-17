// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/user/user.proto

package user

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

type GetUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *model.UserData `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserRes) Reset() {
	*x = GetUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRes) ProtoMessage() {}

func (x *GetUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRes.ProtoReflect.Descriptor instead.
func (*GetUserRes) Descriptor() ([]byte, []int) {
	return file_protos_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRes) GetUser() *model.UserData {
	if x != nil {
		return x.User
	}
	return nil
}

type SaveTutorialReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnumId string `protobuf:"bytes,1,opt,name=enum_id,json=enumId,proto3" json:"enum_id,omitempty"`
}

func (x *SaveTutorialReq) Reset() {
	*x = SaveTutorialReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveTutorialReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveTutorialReq) ProtoMessage() {}

func (x *SaveTutorialReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveTutorialReq.ProtoReflect.Descriptor instead.
func (*SaveTutorialReq) Descriptor() ([]byte, []int) {
	return file_protos_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *SaveTutorialReq) GetEnumId() string {
	if x != nil {
		return x.EnumId
	}
	return ""
}

type SaveTutorialRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reward *model.Reward `protobuf:"bytes,1,opt,name=reward,proto3" json:"reward,omitempty"`
}

func (x *SaveTutorialRes) Reset() {
	*x = SaveTutorialRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveTutorialRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveTutorialRes) ProtoMessage() {}

func (x *SaveTutorialRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveTutorialRes.ProtoReflect.Descriptor instead.
func (*SaveTutorialRes) Descriptor() ([]byte, []int) {
	return file_protos_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *SaveTutorialRes) GetReward() *model.Reward {
	if x != nil {
		return x.Reward
	}
	return nil
}

var File_protos_user_user_proto protoreflect.FileDescriptor

var file_protos_user_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x18,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x0f, 0x53,
	0x61, 0x76, 0x65, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x6e, 0x75, 0x6d, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x54,
	0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x32, 0x7a, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a,
	0x0c, 0x53, 0x61, 0x76, 0x65, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x15, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x46, 0x5a,
	0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x65, 0x6f, 0x6d,
	0x2d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0xaa, 0x02, 0x08, 0x47, 0x72, 0x70,
	0x63, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_user_user_proto_rawDescOnce sync.Once
	file_protos_user_user_proto_rawDescData = file_protos_user_user_proto_rawDesc
)

func file_protos_user_user_proto_rawDescGZIP() []byte {
	file_protos_user_user_proto_rawDescOnce.Do(func() {
		file_protos_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_user_user_proto_rawDescData)
	})
	return file_protos_user_user_proto_rawDescData
}

var file_protos_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_user_user_proto_goTypes = []interface{}{
	(*GetUserRes)(nil),      // 0: user.GetUserRes
	(*SaveTutorialReq)(nil), // 1: user.SaveTutorialReq
	(*SaveTutorialRes)(nil), // 2: user.SaveTutorialRes
	(*model.UserData)(nil),  // 3: model.UserData
	(*model.Reward)(nil),    // 4: model.Reward
	(*model.Empty)(nil),     // 5: model.Empty
}
var file_protos_user_user_proto_depIdxs = []int32{
	3, // 0: user.GetUserRes.user:type_name -> model.UserData
	4, // 1: user.SaveTutorialRes.reward:type_name -> model.Reward
	5, // 2: user.UserService.GetUser:input_type -> model.Empty
	1, // 3: user.UserService.SaveTutorial:input_type -> user.SaveTutorialReq
	0, // 4: user.UserService.GetUser:output_type -> user.GetUserRes
	2, // 5: user.UserService.SaveTutorial:output_type -> user.SaveTutorialRes
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_user_user_proto_init() }
func file_protos_user_user_proto_init() {
	if File_protos_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRes); i {
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
		file_protos_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveTutorialReq); i {
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
		file_protos_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveTutorialRes); i {
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
			RawDescriptor: file_protos_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_user_user_proto_goTypes,
		DependencyIndexes: file_protos_user_user_proto_depIdxs,
		MessageInfos:      file_protos_user_user_proto_msgTypes,
	}.Build()
	File_protos_user_user_proto = out.File
	file_protos_user_user_proto_rawDesc = nil
	file_protos_user_user_proto_goTypes = nil
	file_protos_user_user_proto_depIdxs = nil
}