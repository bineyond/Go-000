// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: url.proto

package url

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ShortenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *ShortenRequest) Reset() {
	*x = ShortenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenRequest) ProtoMessage() {}

func (x *ShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenRequest.ProtoReflect.Descriptor instead.
func (*ShortenRequest) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{0}
}

func (x *ShortenRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

type ShortenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shorten string `protobuf:"bytes,1,opt,name=shorten,proto3" json:"shorten,omitempty"`
}

func (x *ShortenResponse) Reset() {
	*x = ShortenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenResponse) ProtoMessage() {}

func (x *ShortenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenResponse.ProtoReflect.Descriptor instead.
func (*ShortenResponse) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenResponse) GetShorten() string {
	if x != nil {
		return x.Shorten
	}
	return ""
}

var File_url_proto protoreflect.FileDescriptor

var file_url_proto_rawDesc = []byte{
	0x0a, 0x09, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0x28, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x32, 0x44, 0x0a, 0x0a, 0x55, 0x52, 0x4c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x12, 0x13, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_url_proto_rawDescOnce sync.Once
	file_url_proto_rawDescData = file_url_proto_rawDesc
)

func file_url_proto_rawDescGZIP() []byte {
	file_url_proto_rawDescOnce.Do(func() {
		file_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_url_proto_rawDescData)
	})
	return file_url_proto_rawDescData
}

var file_url_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_url_proto_goTypes = []interface{}{
	(*ShortenRequest)(nil),  // 0: url.ShortenRequest
	(*ShortenResponse)(nil), // 1: url.ShortenResponse
}
var file_url_proto_depIdxs = []int32{
	0, // 0: url.URLService.Shorten:input_type -> url.ShortenRequest
	1, // 1: url.URLService.Shorten:output_type -> url.ShortenResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_url_proto_init() }
func file_url_proto_init() {
	if File_url_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_url_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenRequest); i {
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
		file_url_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenResponse); i {
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
			RawDescriptor: file_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_url_proto_goTypes,
		DependencyIndexes: file_url_proto_depIdxs,
		MessageInfos:      file_url_proto_msgTypes,
	}.Build()
	File_url_proto = out.File
	file_url_proto_rawDesc = nil
	file_url_proto_goTypes = nil
	file_url_proto_depIdxs = nil
}
