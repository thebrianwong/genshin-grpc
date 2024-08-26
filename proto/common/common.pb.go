// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto/common/common.proto

package common

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

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Gender  string `protobuf:"bytes,2,opt,name=gender,proto3" json:"gender,omitempty"`
	Height  string `protobuf:"bytes,3,opt,name=height,proto3" json:"height,omitempty"`
	Element string `protobuf:"bytes,4,opt,name=element,proto3" json:"element,omitempty"`
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_proto_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *Character) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Character) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Character) GetHeight() string {
	if x != nil {
		return x.Height
	}
	return ""
}

func (x *Character) GetElement() string {
	if x != nil {
		return x.Element
	}
	return ""
}

type Constellation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Level     uint32     `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Character *Character `protobuf:"bytes,3,opt,name=character,proto3" json:"character,omitempty"`
}

func (x *Constellation) Reset() {
	*x = Constellation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Constellation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Constellation) ProtoMessage() {}

func (x *Constellation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Constellation.ProtoReflect.Descriptor instead.
func (*Constellation) Descriptor() ([]byte, []int) {
	return file_proto_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *Constellation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Constellation) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Constellation) GetCharacter() *Character {
	if x != nil {
		return x.Character
	}
	return nil
}

var File_proto_common_common_proto protoreflect.FileDescriptor

var file_proto_common_common_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x22, 0x69, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x6a,
	0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2f, 0x0a, 0x09, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x65,
	0x6e, 0x73, 0x68, 0x69, 0x6e, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_common_common_proto_rawDescOnce sync.Once
	file_proto_common_common_proto_rawDescData = file_proto_common_common_proto_rawDesc
)

func file_proto_common_common_proto_rawDescGZIP() []byte {
	file_proto_common_common_proto_rawDescOnce.Do(func() {
		file_proto_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_common_common_proto_rawDescData)
	})
	return file_proto_common_common_proto_rawDescData
}

var file_proto_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_common_common_proto_goTypes = []any{
	(*Character)(nil),     // 0: common.Character
	(*Constellation)(nil), // 1: common.Constellation
}
var file_proto_common_common_proto_depIdxs = []int32{
	0, // 0: common.Constellation.character:type_name -> common.Character
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_common_common_proto_init() }
func file_proto_common_common_proto_init() {
	if File_proto_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_common_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Character); i {
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
		file_proto_common_common_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Constellation); i {
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
			RawDescriptor: file_proto_common_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_common_common_proto_goTypes,
		DependencyIndexes: file_proto_common_common_proto_depIdxs,
		MessageInfos:      file_proto_common_common_proto_msgTypes,
	}.Build()
	File_proto_common_common_proto = out.File
	file_proto_common_common_proto_rawDesc = nil
	file_proto_common_common_proto_goTypes = nil
	file_proto_common_common_proto_depIdxs = nil
}
