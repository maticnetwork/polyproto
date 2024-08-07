// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: common/common.proto

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

type H128 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hi uint64 `protobuf:"varint,1,opt,name=Hi,proto3" json:"Hi,omitempty"`
	Lo uint64 `protobuf:"varint,2,opt,name=Lo,proto3" json:"Lo,omitempty"`
}

func (x *H128) Reset() {
	*x = H128{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *H128) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*H128) ProtoMessage() {}

func (x *H128) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use H128.ProtoReflect.Descriptor instead.
func (*H128) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *H128) GetHi() uint64 {
	if x != nil {
		return x.Hi
	}
	return 0
}

func (x *H128) GetLo() uint64 {
	if x != nil {
		return x.Lo
	}
	return 0
}

type H160 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hi *H128  `protobuf:"bytes,1,opt,name=Hi,proto3" json:"Hi,omitempty"`
	Lo uint32 `protobuf:"varint,2,opt,name=Lo,proto3" json:"Lo,omitempty"`
}

func (x *H160) Reset() {
	*x = H160{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *H160) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*H160) ProtoMessage() {}

func (x *H160) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use H160.ProtoReflect.Descriptor instead.
func (*H160) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *H160) GetHi() *H128 {
	if x != nil {
		return x.Hi
	}
	return nil
}

func (x *H160) GetLo() uint32 {
	if x != nil {
		return x.Lo
	}
	return 0
}

type H256 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hi *H128 `protobuf:"bytes,1,opt,name=Hi,proto3" json:"Hi,omitempty"`
	Lo *H128 `protobuf:"bytes,2,opt,name=Lo,proto3" json:"Lo,omitempty"`
}

func (x *H256) Reset() {
	*x = H256{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *H256) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*H256) ProtoMessage() {}

func (x *H256) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use H256.ProtoReflect.Descriptor instead.
func (*H256) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *H256) GetHi() *H128 {
	if x != nil {
		return x.Hi
	}
	return nil
}

func (x *H256) GetLo() *H128 {
	if x != nil {
		return x.Lo
	}
	return nil
}

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x26, 0x0a,
	0x04, 0x48, 0x31, 0x32, 0x38, 0x12, 0x0e, 0x0a, 0x02, 0x48, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x48, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x4c, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x4c, 0x6f, 0x22, 0x34, 0x0a, 0x04, 0x48, 0x31, 0x36, 0x30, 0x12, 0x1c, 0x0a,
	0x02, 0x48, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x48, 0x31, 0x32, 0x38, 0x52, 0x02, 0x48, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x4c,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x4c, 0x6f, 0x22, 0x42, 0x0a, 0x04, 0x48,
	0x32, 0x35, 0x36, 0x12, 0x1c, 0x0a, 0x02, 0x48, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x31, 0x32, 0x38, 0x52, 0x02, 0x48,
	0x69, 0x12, 0x1c, 0x0a, 0x02, 0x4c, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x31, 0x32, 0x38, 0x52, 0x02, 0x4c, 0x6f, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61,
	0x74, 0x69, 0x63, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData = file_common_common_proto_rawDesc
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_common_proto_rawDescData)
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_common_proto_goTypes = []interface{}{
	(*H128)(nil), // 0: common.H128
	(*H160)(nil), // 1: common.H160
	(*H256)(nil), // 2: common.H256
}
var file_common_common_proto_depIdxs = []int32{
	0, // 0: common.H160.Hi:type_name -> common.H128
	0, // 1: common.H256.Hi:type_name -> common.H128
	0, // 2: common.H256.Lo:type_name -> common.H128
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*H128); i {
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
		file_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*H160); i {
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
		file_common_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*H256); i {
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
			RawDescriptor: file_common_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_rawDesc = nil
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}
