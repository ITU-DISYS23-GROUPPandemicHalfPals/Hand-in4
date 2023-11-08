// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0--rc2
// source: mutualExclusion.proto

package mutualExclusion

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

type ElectionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ElectionMessage) Reset() {
	*x = ElectionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mutualExclusion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionMessage) ProtoMessage() {}

func (x *ElectionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mutualExclusion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionMessage.ProtoReflect.Descriptor instead.
func (*ElectionMessage) Descriptor() ([]byte, []int) {
	return file_mutualExclusion_proto_rawDescGZIP(), []int{0}
}

func (x *ElectionMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CoordinatorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CoordinatorMessage) Reset() {
	*x = CoordinatorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mutualExclusion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoordinatorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoordinatorMessage) ProtoMessage() {}

func (x *CoordinatorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mutualExclusion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoordinatorMessage.ProtoReflect.Descriptor instead.
func (*CoordinatorMessage) Descriptor() ([]byte, []int) {
	return file_mutualExclusion_proto_rawDescGZIP(), []int{1}
}

func (x *CoordinatorMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mutualExclusion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_mutualExclusion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_mutualExclusion_proto_rawDescGZIP(), []int{2}
}

var File_mutualExclusion_proto protoreflect.FileDescriptor

var file_mutualExclusion_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x45,
	0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x0f, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xad, 0x01,
	0x0a, 0x0f, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x49, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e,
	0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x19, 0x2e, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0b,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x23, 0x2e, 0x6d, 0x75,
	0x74, 0x75, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x19, 0x2e, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x21, 0x5a,
	0x1f, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mutualExclusion_proto_rawDescOnce sync.Once
	file_mutualExclusion_proto_rawDescData = file_mutualExclusion_proto_rawDesc
)

func file_mutualExclusion_proto_rawDescGZIP() []byte {
	file_mutualExclusion_proto_rawDescOnce.Do(func() {
		file_mutualExclusion_proto_rawDescData = protoimpl.X.CompressGZIP(file_mutualExclusion_proto_rawDescData)
	})
	return file_mutualExclusion_proto_rawDescData
}

var file_mutualExclusion_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mutualExclusion_proto_goTypes = []interface{}{
	(*ElectionMessage)(nil),    // 0: mutualExclusion.ElectionMessage
	(*CoordinatorMessage)(nil), // 1: mutualExclusion.CoordinatorMessage
	(*Response)(nil),           // 2: mutualExclusion.Response
}
var file_mutualExclusion_proto_depIdxs = []int32{
	0, // 0: mutualExclusion.MutualExclusion.Election:input_type -> mutualExclusion.ElectionMessage
	1, // 1: mutualExclusion.MutualExclusion.Coordinator:input_type -> mutualExclusion.CoordinatorMessage
	2, // 2: mutualExclusion.MutualExclusion.Election:output_type -> mutualExclusion.Response
	2, // 3: mutualExclusion.MutualExclusion.Coordinator:output_type -> mutualExclusion.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mutualExclusion_proto_init() }
func file_mutualExclusion_proto_init() {
	if File_mutualExclusion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mutualExclusion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionMessage); i {
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
		file_mutualExclusion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoordinatorMessage); i {
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
		file_mutualExclusion_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_mutualExclusion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mutualExclusion_proto_goTypes,
		DependencyIndexes: file_mutualExclusion_proto_depIdxs,
		MessageInfos:      file_mutualExclusion_proto_msgTypes,
	}.Build()
	File_mutualExclusion_proto = out.File
	file_mutualExclusion_proto_rawDesc = nil
	file_mutualExclusion_proto_goTypes = nil
	file_mutualExclusion_proto_depIdxs = nil
}