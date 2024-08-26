// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto/character/character.proto

package character

import (
	common "genshin-test/proto/common"
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

type GetCharacterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCharacterRequest) Reset() {
	*x = GetCharacterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_character_character_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCharacterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCharacterRequest) ProtoMessage() {}

func (x *GetCharacterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_character_character_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCharacterRequest.ProtoReflect.Descriptor instead.
func (*GetCharacterRequest) Descriptor() ([]byte, []int) {
	return file_proto_character_character_proto_rawDescGZIP(), []int{0}
}

func (x *GetCharacterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCharacterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Character []*common.Character `protobuf:"bytes,1,rep,name=character,proto3" json:"character,omitempty"`
}

func (x *GetCharacterResponse) Reset() {
	*x = GetCharacterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_character_character_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCharacterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCharacterResponse) ProtoMessage() {}

func (x *GetCharacterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_character_character_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCharacterResponse.ProtoReflect.Descriptor instead.
func (*GetCharacterResponse) Descriptor() ([]byte, []int) {
	return file_proto_character_character_proto_rawDescGZIP(), []int{1}
}

func (x *GetCharacterResponse) GetCharacter() []*common.Character {
	if x != nil {
		return x.Character
	}
	return nil
}

var File_proto_character_character_proto protoreflect.FileDescriptor

var file_proto_character_character_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x1a, 0x19, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x32, 0x65, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1e,
	0x5a, 0x1c, 0x67, 0x65, 0x6e, 0x73, 0x68, 0x69, 0x6e, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_character_character_proto_rawDescOnce sync.Once
	file_proto_character_character_proto_rawDescData = file_proto_character_character_proto_rawDesc
)

func file_proto_character_character_proto_rawDescGZIP() []byte {
	file_proto_character_character_proto_rawDescOnce.Do(func() {
		file_proto_character_character_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_character_character_proto_rawDescData)
	})
	return file_proto_character_character_proto_rawDescData
}

var file_proto_character_character_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_character_character_proto_goTypes = []any{
	(*GetCharacterRequest)(nil),  // 0: character.GetCharacterRequest
	(*GetCharacterResponse)(nil), // 1: character.GetCharacterResponse
	(*common.Character)(nil),     // 2: common.Character
}
var file_proto_character_character_proto_depIdxs = []int32{
	2, // 0: character.GetCharacterResponse.character:type_name -> common.Character
	0, // 1: character.CharacterService.GetCharacter:input_type -> character.GetCharacterRequest
	1, // 2: character.CharacterService.GetCharacter:output_type -> character.GetCharacterResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_character_character_proto_init() }
func file_proto_character_character_proto_init() {
	if File_proto_character_character_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_character_character_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetCharacterRequest); i {
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
		file_proto_character_character_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetCharacterResponse); i {
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
			RawDescriptor: file_proto_character_character_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_character_character_proto_goTypes,
		DependencyIndexes: file_proto_character_character_proto_depIdxs,
		MessageInfos:      file_proto_character_character_proto_msgTypes,
	}.Build()
	File_proto_character_character_proto = out.File
	file_proto_character_character_proto_rawDesc = nil
	file_proto_character_character_proto_goTypes = nil
	file_proto_character_character_proto_depIdxs = nil
}
