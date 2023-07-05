// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: sample-v2/sample-v2.proto

package sample_v2

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

type SampleV2Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SampleV2Request) Reset() {
	*x = SampleV2Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_v2_sample_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleV2Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleV2Request) ProtoMessage() {}

func (x *SampleV2Request) ProtoReflect() protoreflect.Message {
	mi := &file_sample_v2_sample_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleV2Request.ProtoReflect.Descriptor instead.
func (*SampleV2Request) Descriptor() ([]byte, []int) {
	return file_sample_v2_sample_v2_proto_rawDescGZIP(), []int{0}
}

func (x *SampleV2Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SampleV2Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SampleV2Response) Reset() {
	*x = SampleV2Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_v2_sample_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleV2Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleV2Response) ProtoMessage() {}

func (x *SampleV2Response) ProtoReflect() protoreflect.Message {
	mi := &file_sample_v2_sample_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleV2Response.ProtoReflect.Descriptor instead.
func (*SampleV2Response) Descriptor() ([]byte, []int) {
	return file_sample_v2_sample_v2_proto_rawDescGZIP(), []int{1}
}

func (x *SampleV2Response) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SampleV2Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_sample_v2_sample_v2_proto protoreflect.FileDescriptor

var file_sample_v2_sample_v2_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x76, 0x32, 0x2f, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2d, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x32, 0x2e,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x56, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a,
	0x10, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9f, 0x01, 0x0a, 0x08,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x32, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x32, 0x12, 0x1a, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x32, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x67, 0x61,
	0x69, 0x6e, 0x56, 0x32, 0x12, 0x1a, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x56, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sample_v2_sample_v2_proto_rawDescOnce sync.Once
	file_sample_v2_sample_v2_proto_rawDescData = file_sample_v2_sample_v2_proto_rawDesc
)

func file_sample_v2_sample_v2_proto_rawDescGZIP() []byte {
	file_sample_v2_sample_v2_proto_rawDescOnce.Do(func() {
		file_sample_v2_sample_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_sample_v2_sample_v2_proto_rawDescData)
	})
	return file_sample_v2_sample_v2_proto_rawDescData
}

var file_sample_v2_sample_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sample_v2_sample_v2_proto_goTypes = []interface{}{
	(*SampleV2Request)(nil),  // 0: v2.sample.SampleV2Request
	(*SampleV2Response)(nil), // 1: v2.sample.SampleV2Response
}
var file_sample_v2_sample_v2_proto_depIdxs = []int32{
	0, // 0: v2.sample.SampleV2.GetSampleV2:input_type -> v2.sample.SampleV2Request
	0, // 1: v2.sample.SampleV2.GetSampleAgainV2:input_type -> v2.sample.SampleV2Request
	1, // 2: v2.sample.SampleV2.GetSampleV2:output_type -> v2.sample.SampleV2Response
	1, // 3: v2.sample.SampleV2.GetSampleAgainV2:output_type -> v2.sample.SampleV2Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sample_v2_sample_v2_proto_init() }
func file_sample_v2_sample_v2_proto_init() {
	if File_sample_v2_sample_v2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sample_v2_sample_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleV2Request); i {
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
		file_sample_v2_sample_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleV2Response); i {
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
			RawDescriptor: file_sample_v2_sample_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sample_v2_sample_v2_proto_goTypes,
		DependencyIndexes: file_sample_v2_sample_v2_proto_depIdxs,
		MessageInfos:      file_sample_v2_sample_v2_proto_msgTypes,
	}.Build()
	File_sample_v2_sample_v2_proto = out.File
	file_sample_v2_sample_v2_proto_rawDesc = nil
	file_sample_v2_sample_v2_proto_goTypes = nil
	file_sample_v2_sample_v2_proto_depIdxs = nil
}
