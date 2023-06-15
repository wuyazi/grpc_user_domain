// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: grpc_user_domain.proto

package user_domain

import (
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

type CreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_user_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReq.ProtoReflect.Descriptor instead.
func (*CreateReq) Descriptor() ([]byte, []int) {
	return file_grpc_user_domain_proto_rawDescGZIP(), []int{0}
}

func (x *CreateReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type UpdateNicknameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *UpdateNicknameReq) Reset() {
	*x = UpdateNicknameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_user_domain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNicknameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNicknameReq) ProtoMessage() {}

func (x *UpdateNicknameReq) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_domain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNicknameReq.ProtoReflect.Descriptor instead.
func (*UpdateNicknameReq) Descriptor() ([]byte, []int) {
	return file_grpc_user_domain_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateNicknameReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateNicknameReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type UpdateAgeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Age    int64 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UpdateAgeReq) Reset() {
	*x = UpdateAgeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_user_domain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAgeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAgeReq) ProtoMessage() {}

func (x *UpdateAgeReq) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_domain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAgeReq.ProtoReflect.Descriptor instead.
func (*UpdateAgeReq) Descriptor() ([]byte, []int) {
	return file_grpc_user_domain_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAgeReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateAgeReq) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type UpdateGenderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Gender string `protobuf:"bytes,2,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *UpdateGenderReq) Reset() {
	*x = UpdateGenderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_user_domain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGenderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGenderReq) ProtoMessage() {}

func (x *UpdateGenderReq) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_domain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGenderReq.ProtoReflect.Descriptor instead.
func (*UpdateGenderReq) Descriptor() ([]byte, []int) {
	return file_grpc_user_domain_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateGenderReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateGenderReq) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

// service
type UserIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserIdResp) Reset() {
	*x = UserIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_user_domain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdResp) ProtoMessage() {}

func (x *UserIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_domain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdResp.ProtoReflect.Descriptor instead.
func (*UserIdResp) Descriptor() ([]byte, []int) {
	return file_grpc_user_domain_proto_rawDescGZIP(), []int{4}
}

func (x *UserIdResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_grpc_user_domain_proto protoreflect.FileDescriptor

var file_grpc_user_domain_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x27, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x47,
	0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x22, 0x41, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x9b, 0x02, 0x0a, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x49, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x3f, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x12, 0x19, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x45, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x75, 0x79, 0x61, 0x7a, 0x69, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grpc_user_domain_proto_rawDescOnce sync.Once
	file_grpc_user_domain_proto_rawDescData = file_grpc_user_domain_proto_rawDesc
)

func file_grpc_user_domain_proto_rawDescGZIP() []byte {
	file_grpc_user_domain_proto_rawDescOnce.Do(func() {
		file_grpc_user_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_user_domain_proto_rawDescData)
	})
	return file_grpc_user_domain_proto_rawDescData
}

var file_grpc_user_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_user_domain_proto_goTypes = []interface{}{
	(*CreateReq)(nil),         // 0: user_domain.createReq
	(*UpdateNicknameReq)(nil), // 1: user_domain.updateNicknameReq
	(*UpdateAgeReq)(nil),      // 2: user_domain.updateAgeReq
	(*UpdateGenderReq)(nil),   // 3: user_domain.updateGenderReq
	(*UserIdResp)(nil),        // 4: user_domain.userIdResp
}
var file_grpc_user_domain_proto_depIdxs = []int32{
	0, // 0: user_domain.userCommend.create:input_type -> user_domain.createReq
	1, // 1: user_domain.userCommend.updateNickname:input_type -> user_domain.updateNicknameReq
	2, // 2: user_domain.userCommend.updateAge:input_type -> user_domain.updateAgeReq
	3, // 3: user_domain.userCommend.updateGender:input_type -> user_domain.updateGenderReq
	4, // 4: user_domain.userCommend.create:output_type -> user_domain.userIdResp
	4, // 5: user_domain.userCommend.updateNickname:output_type -> user_domain.userIdResp
	4, // 6: user_domain.userCommend.updateAge:output_type -> user_domain.userIdResp
	4, // 7: user_domain.userCommend.updateGender:output_type -> user_domain.userIdResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_user_domain_proto_init() }
func file_grpc_user_domain_proto_init() {
	if File_grpc_user_domain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_user_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReq); i {
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
		file_grpc_user_domain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNicknameReq); i {
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
		file_grpc_user_domain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAgeReq); i {
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
		file_grpc_user_domain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGenderReq); i {
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
		file_grpc_user_domain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdResp); i {
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
			RawDescriptor: file_grpc_user_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_user_domain_proto_goTypes,
		DependencyIndexes: file_grpc_user_domain_proto_depIdxs,
		MessageInfos:      file_grpc_user_domain_proto_msgTypes,
	}.Build()
	File_grpc_user_domain_proto = out.File
	file_grpc_user_domain_proto_rawDesc = nil
	file_grpc_user_domain_proto_goTypes = nil
	file_grpc_user_domain_proto_depIdxs = nil
}